package idp

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antihax/optional"
	idp_sdk "github.com/harness/harness-go-sdk/harness/idp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

type idpErrorWithBody struct {
	body []byte
}

func (e idpErrorWithBody) Error() string {
	return "400 Bad Request"
}

func (e idpErrorWithBody) Body() []byte {
	return e.body
}

func TestGetCatalogEntityInfoFromResourceData(t *testing.T) {
	resource := ResourceCatalogEntity()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"identifier": "my_service_dev_landing_zone",
		"kind":       "resource",
		"org_id":     "default",
		"project_id": "idp",
	})

	info, err := getCatalogEntityInfoFromResourceData(data)

	require.NoError(t, err)
	require.Equal(t, "my_service_dev_landing_zone", info.Identifier)
	require.Equal(t, "resource", info.Kind)
	require.Equal(t, "account.default.idp", info.Scope)
	require.True(t, info.OrgId.IsSet())
	require.Equal(t, "default", info.OrgId.Value())
	require.True(t, info.ProjectId.IsSet())
	require.Equal(t, "idp", info.ProjectId.Value())
}

func TestGetCatalogEntityInfoFromResourceDataIncludesBranchName(t *testing.T) {
	resource := ResourceCatalogEntity()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"identifier": "my_service_dev_landing_zone",
		"kind":       "resource",
		"org_id":     "default",
		"project_id": "idp",
		"git_details": []interface{}{
			map[string]interface{}{
				"branch_name": "feat/idp-workflow",
				"base_branch": "main",
				"store_type":  "REMOTE",
			},
		},
	})

	info, err := getCatalogEntityInfoFromResourceData(data)

	require.NoError(t, err)
	require.True(t, info.BranchName.IsSet())
	require.Equal(t, "feat/idp-workflow", info.BranchName.Value())
}

func TestBaseBranchIsOnlySentDuringCreate(t *testing.T) {
	resource := ResourceCatalogEntity()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"git_details": []interface{}{
			map[string]interface{}{
				"branch_name": "feat/idp-workflow",
				"base_branch": "develop",
				"store_type":  "REMOTE",
			},
		},
	})

	createDetails := buildGitCreateDetails(data)
	updateDetails := buildGitUpdateDetails(data)

	require.Equal(t, "develop", createDetails.BaseBranch)
	require.Equal(t, "feat/idp-workflow", createDetails.BranchName)
	require.Empty(t, updateDetails.BaseBranch)
	require.Equal(t, "feat/idp-workflow", updateDetails.BranchName)
}

func TestCatalogEntityGetOptsUsesBranchName(t *testing.T) {
	c := &idp_sdk.APIClient{AccountId: "account-123"}
	opts := catalogEntityGetOpts(c, catalogEntityInfo{
		OrgId:      optional.NewString("default"),
		ProjectId:  optional.NewString("idp"),
		BranchName: optional.NewString("feat/idp-workflow"),
	})

	require.True(t, opts.BranchName.IsSet())
	require.Equal(t, "feat/idp-workflow", opts.BranchName.Value())
	require.False(t, opts.LoadFromFallbackBranch.IsSet())
	require.Equal(t, "account-123", opts.HarnessAccount.Value())
}

func TestCatalogEntityGetOptsOmitsBranchWhenMissing(t *testing.T) {
	c := &idp_sdk.APIClient{AccountId: "account-123"}
	opts := catalogEntityGetOpts(c, catalogEntityInfo{})

	require.False(t, opts.BranchName.IsSet())
	require.False(t, opts.LoadFromFallbackBranch.IsSet())
}

func TestGetCatalogEntitySendsBranchName(t *testing.T) {
	requestQuery := make(chan string, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestQuery <- r.URL.Query().Get("branch_name")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	t.Cleanup(server.Close)

	config := idp_sdk.NewConfiguration()
	config.BasePath = server.URL
	config.AccountId = "account-123"
	client := idp_sdk.NewAPIClient(config)

	_, _, err := getCatalogEntity(context.Background(), client, catalogEntityInfo{
		Scope:      "account.default.idp",
		Kind:       "resource",
		Identifier: "my_service",
		OrgId:      optional.NewString("default"),
		ProjectId:  optional.NewString("idp"),
		BranchName: optional.NewString("feat/idp-workflow"),
	})

	require.NoError(t, err)
	require.Equal(t, "feat/idp-workflow", <-requestQuery)
}

func TestGetCatalogEntityInfoFromResponseCopiesGitBranch(t *testing.T) {
	fallback := catalogEntityInfo{
		Scope:      "account",
		Kind:       "resource",
		Identifier: "my_service",
		OrgId:      optional.EmptyString(),
		ProjectId:  optional.EmptyString(),
	}

	info := getCatalogEntityInfoFromResponse(idp_sdk.EntityResponse{
		Identifier:        "my_service",
		Kind:              "resource",
		OrgIdentifier:     "default",
		ProjectIdentifier: "idp",
		GitDetails: &idp_sdk.GitDetails{
			BranchName: "feat/idp-workflow",
		},
	}, fallback)

	require.True(t, info.BranchName.IsSet())
	require.Equal(t, "feat/idp-workflow", info.BranchName.Value())
}

func TestGetCatalogEntityInfoFromImportResourceDataAllowsComputedKind(t *testing.T) {
	resource := ResourceCatalogEntity()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"identifier":      "my_service_dev_landing_zone",
		"import_from_git": true,
		"org_id":          "default",
		"project_id":      "idp",
	})

	info, err := getCatalogEntityInfoFromImportResourceData(data)

	require.NoError(t, err)
	require.Equal(t, "my_service_dev_landing_zone", info.Identifier)
	require.Empty(t, info.Kind)
	require.Equal(t, "account.default.idp", info.Scope)
	require.True(t, info.OrgId.IsSet())
	require.Equal(t, "default", info.OrgId.Value())
	require.True(t, info.ProjectId.IsSet())
	require.Equal(t, "idp", info.ProjectId.Value())
}

func TestGetCatalogEntityInfoFromResponseUsesCanonicalResponseFields(t *testing.T) {
	fallback := catalogEntityInfo{
		Scope:      "account",
		Kind:       "resource",
		Identifier: "my-service_dev_landing_zone",
		OrgId:      optional.EmptyString(),
		ProjectId:  optional.EmptyString(),
	}

	info := getCatalogEntityInfoFromResponse(idp_sdk.EntityResponse{
		Identifier:        "my_service_dev_landing_zone",
		Kind:              "resource",
		OrgIdentifier:     "default",
		ProjectIdentifier: "idp",
	}, fallback)

	require.Equal(t, "my_service_dev_landing_zone", info.Identifier)
	require.Equal(t, "resource", info.Kind)
	require.Equal(t, "account.default.idp", info.Scope)
	require.True(t, info.OrgId.IsSet())
	require.Equal(t, "default", info.OrgId.Value())
	require.True(t, info.ProjectId.IsSet())
	require.Equal(t, "idp", info.ProjectId.Value())
}

func TestGetCatalogEntityInfoFromResponseFallsBackWhenResponseIsPartial(t *testing.T) {
	fallback := catalogEntityInfo{
		Scope:      "account.default",
		Kind:       "resource",
		Identifier: "my_service_dev_landing_zone",
		OrgId:      optional.NewString("default"),
		ProjectId:  optional.EmptyString(),
	}

	info := getCatalogEntityInfoFromResponse(idp_sdk.EntityResponse{}, fallback)

	require.Equal(t, fallback.Identifier, info.Identifier)
	require.Equal(t, fallback.Kind, info.Kind)
	require.Equal(t, fallback.Scope, info.Scope)
	require.True(t, info.OrgId.IsSet())
	require.Equal(t, "default", info.OrgId.Value())
	require.False(t, info.ProjectId.IsSet())
}

func TestGetCatalogEntityInfoFromResponsePreservesFallbackProjectScope(t *testing.T) {
	fallback := catalogEntityInfo{
		Scope:      "account.default.idp",
		Kind:       "resource",
		Identifier: "my_service_dev_landing_zone",
		OrgId:      optional.NewString("default"),
		ProjectId:  optional.NewString("idp"),
	}

	info := getCatalogEntityInfoFromResponse(idp_sdk.EntityResponse{
		Identifier: "my_service_dev_landing_zone",
		Kind:       "resource",
	}, fallback)

	require.Equal(t, "account.default.idp", info.Scope)
	require.True(t, info.OrgId.IsSet())
	require.Equal(t, "default", info.OrgId.Value())
	require.True(t, info.ProjectId.IsSet())
	require.Equal(t, "idp", info.ProjectId.Value())
}

func TestGetCatalogEntityInfoFromResponseDoesNotMixResponseOrgWithFallbackProject(t *testing.T) {
	fallback := catalogEntityInfo{
		Scope:      "account.old_org.old_project",
		Kind:       "resource",
		Identifier: "my_service_dev_landing_zone",
		OrgId:      optional.NewString("old_org"),
		ProjectId:  optional.NewString("old_project"),
	}

	info := getCatalogEntityInfoFromResponse(idp_sdk.EntityResponse{
		Identifier:    "my_service_dev_landing_zone",
		Kind:          "resource",
		OrgIdentifier: "new_org",
	}, fallback)

	require.Equal(t, "account.old_org.old_project", info.Scope)
	require.True(t, info.OrgId.IsSet())
	require.Equal(t, "old_org", info.OrgId.Value())
	require.True(t, info.ProjectId.IsSet())
	require.Equal(t, "old_project", info.ProjectId.Value())
}

func TestGetCatalogEntityInfoFromResponseUsesOrgOnlyScopeForOrgFallback(t *testing.T) {
	fallback := catalogEntityInfo{
		Scope:      "account.old_org",
		Kind:       "resource",
		Identifier: "my_service_dev_landing_zone",
		OrgId:      optional.NewString("old_org"),
		ProjectId:  optional.EmptyString(),
	}

	info := getCatalogEntityInfoFromResponse(idp_sdk.EntityResponse{
		Identifier:    "my_service_dev_landing_zone",
		Kind:          "resource",
		OrgIdentifier: "new_org",
	}, fallback)

	require.Equal(t, "account.new_org", info.Scope)
	require.True(t, info.OrgId.IsSet())
	require.Equal(t, "new_org", info.OrgId.Value())
	require.False(t, info.ProjectId.IsSet())
}

func TestGetCatalogEntityInfoFromResponseIgnoresProjectOnlyScope(t *testing.T) {
	fallback := catalogEntityInfo{
		Scope:      "account.default.idp",
		Kind:       "resource",
		Identifier: "my_service_dev_landing_zone",
		OrgId:      optional.NewString("default"),
		ProjectId:  optional.NewString("idp"),
	}

	info := getCatalogEntityInfoFromResponse(idp_sdk.EntityResponse{
		Identifier:        "my_service_dev_landing_zone",
		Kind:              "resource",
		ProjectIdentifier: "new_project",
	}, fallback)

	require.Equal(t, "account.default.idp", info.Scope)
	require.True(t, info.OrgId.IsSet())
	require.Equal(t, "default", info.OrgId.Value())
	require.True(t, info.ProjectId.IsSet())
	require.Equal(t, "idp", info.ProjectId.Value())
}

func TestIsTransientPostWriteReadError(t *testing.T) {
	err := errors.New("read failed")

	require.True(t, isTransientPostWriteReadError(err, &http.Response{StatusCode: http.StatusNotFound}))
	require.False(t, isTransientPostWriteReadError(err, &http.Response{StatusCode: http.StatusUnauthorized}))
	require.False(t, isTransientPostWriteReadError(err, &http.Response{StatusCode: http.StatusForbidden}))
	require.False(t, isTransientPostWriteReadError(nil, &http.Response{StatusCode: http.StatusNotFound}))
}

func TestReadGitDetailsWithUnsetConnectorRef(t *testing.T) {
	gitDetails := readGitDetails(idp_sdk.EntityResponse{
		GitDetails: &idp_sdk.GitDetails{
			BranchName: "main",
			FilePath:   "catalog/entity.yaml",
			RepoName:   "catalog",
		},
	}, optional.EmptyString(), optional.EmptyString(), optional.EmptyString(), optional.EmptyString())

	require.NotContains(t, gitDetails, "connector_ref")
	require.Equal(t, true, gitDetails["is_harness_code_repo"])
}

func TestIDPAPIErrorMessage(t *testing.T) {
	err := idpErrorWithBody{body: []byte(`{"code":"INVALID_REQUEST","message":"Invalid request: Entity identifier = tf-idp is invalid"}`)}

	require.Equal(t, "INVALID_REQUEST: Invalid request: Entity identifier = tf-idp is invalid", idpAPIErrorMessage(err))
}

func TestIsIDPNotFoundError(t *testing.T) {
	require.True(t, isIDPNotFoundError(errors.New("not found"), &http.Response{StatusCode: http.StatusNotFound}))
	require.True(t, isIDPNotFoundError(idpErrorWithBody{body: []byte(`{"code":"ENTITY_NOT_FOUND","message":"missing"}`)}, nil))
	require.True(t, isIDPNotFoundError(idpErrorWithBody{body: []byte(`{"message":"Check details not found for checkId [x]"}`)}, &http.Response{StatusCode: http.StatusInternalServerError}))
	require.False(t, isIDPNotFoundError(idpErrorWithBody{body: []byte(`{"code":"INVALID_REQUEST","message":"bad request"}`)}, nil))
}

func TestHandleIDPWriteApiErrorAlreadyExists(t *testing.T) {
	resource := ResourceScorecardCheck()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"identifier": "tf_sanity_readme_exists",
		"name":       "README exists",
	})
	err := idpErrorWithBody{body: []byte(`{"message":"Check [tf_sanity_readme_exists] already created for accountId [account-123]"}`)}

	diags := handleIDPWriteApiError("harness_platform_idp_scorecard_check", err, data, &http.Response{StatusCode: http.StatusInternalServerError})
	require.True(t, diags.HasError())
	require.Contains(t, diags[0].Summary, "already created")
	require.Contains(t, diags[0].Summary, "terraform import harness_platform_idp_scorecard_check.<name> tf_sanity_readme_exists")
}

func TestHandleIDPWriteApiErrorScorecardAlreadyExists(t *testing.T) {
	resource := ResourceScorecard()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"identifier": "tf_sanity_gold_standard",
		"name":       "Gold Standard",
	})
	err := idpErrorWithBody{body: []byte(
		`{"message":"A scorecard with identifier 'tf_sanity_gold_standard' already exists. Please use a different identifier."}`,
	)}

	diags := handleIDPWriteApiError(
		"harness_platform_idp_scorecard",
		err,
		data,
		&http.Response{StatusCode: http.StatusInternalServerError},
	)
	require.True(t, diags.HasError())
	require.Contains(t, diags[0].Summary, "already exists")
	require.Contains(t, diags[0].Summary,
		"terraform import harness_platform_idp_scorecard.<name> tf_sanity_gold_standard")
}

func TestHandleIDPApiErrorReferencedCheck(t *testing.T) {
	resource := ResourceScorecardCheck()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"identifier": "tf_sanity_readme_exists",
		"name":       "README exists",
	})
	err := idpErrorWithBody{body: []byte(`{"message":"Could not delete the check [tf_sanity_readme_exists] as it is referenced by other scorecards"}`)}

	diags := handleIDPApiError(err, data, &http.Response{StatusCode: http.StatusInternalServerError})
	require.True(t, diags.HasError())
	require.Contains(t, diags[0].Summary, "referenced by other scorecards")
	require.Contains(t, diags[0].Summary, "Remove the check from those scorecards first")
}

func TestCatalogEntityImportInfoFromID(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		expectedScope  string
		expectedKind   string
		expectedID     string
		expectedOrg    string
		expectedProj   string
		expectedErrMsg string
	}{
		{
			name:          "account",
			id:            "component/my_component",
			expectedScope: "account",
			expectedKind:  "component",
			expectedID:    "my_component",
		},
		{
			name:          "org",
			id:            "my_org/component/my_component",
			expectedScope: "account.my_org",
			expectedKind:  "component",
			expectedID:    "my_component",
			expectedOrg:   "my_org",
		},
		{
			name:          "project",
			id:            "my_org/my_project/component/my_component",
			expectedScope: "account.my_org.my_project",
			expectedKind:  "component",
			expectedID:    "my_component",
			expectedOrg:   "my_org",
			expectedProj:  "my_project",
		},
		{
			name:          "legacy dot project",
			id:            "my_org.my_project/component/my_component",
			expectedScope: "account.my_org.my_project",
			expectedKind:  "component",
			expectedID:    "my_component",
			expectedOrg:   "my_org",
			expectedProj:  "my_project",
		},
		{
			name:           "invalid",
			id:             "my_org/my_project/component/my_component/extra",
			expectedErrMsg: "invalid import ID format",
		},
		{
			name:           "invalid empty dot project",
			id:             "my_org./component/my_component",
			expectedErrMsg: "invalid import scope",
		},
		{
			name:           "invalid overlong dot project",
			id:             "my_org.my_project.extra/component/my_component",
			expectedErrMsg: "invalid import scope",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := catalogEntityInfoFromImportID(tt.id)
			if tt.expectedErrMsg != "" {
				require.ErrorContains(t, err, tt.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedScope, info.Scope)
			require.Equal(t, tt.expectedKind, info.Kind)
			require.Equal(t, tt.expectedID, info.Identifier)
			if tt.expectedOrg == "" {
				require.False(t, info.OrgId.IsSet())
			} else {
				require.True(t, info.OrgId.IsSet())
				require.Equal(t, tt.expectedOrg, info.OrgId.Value())
			}
			if tt.expectedProj == "" {
				require.False(t, info.ProjectId.IsSet())
			} else {
				require.True(t, info.ProjectId.IsSet())
				require.Equal(t, tt.expectedProj, info.ProjectId.Value())
			}
		})
	}
}
