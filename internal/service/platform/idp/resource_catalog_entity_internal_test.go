package idp

import (
	"errors"
	"net/http"
	"testing"

	"github.com/antihax/optional"
	idp_sdk "github.com/harness/harness-go-sdk/harness/idp"
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

func TestIsTransientPostCreateReadError(t *testing.T) {
	err := errors.New("read failed")

	require.True(t, isTransientPostCreateReadError(err, &http.Response{StatusCode: http.StatusNotFound}))
	require.False(t, isTransientPostCreateReadError(err, &http.Response{StatusCode: http.StatusUnauthorized}))
	require.False(t, isTransientPostCreateReadError(err, &http.Response{StatusCode: http.StatusForbidden}))
	require.False(t, isTransientPostCreateReadError(nil, &http.Response{StatusCode: http.StatusNotFound}))
}

func TestIDPAPIErrorMessage(t *testing.T) {
	err := idpErrorWithBody{body: []byte(`{"code":"INVALID_REQUEST","message":"Invalid request: Entity identifier = tf-idp is invalid"}`)}

	require.Equal(t, "INVALID_REQUEST: Invalid request: Entity identifier = tf-idp is invalid", idpAPIErrorMessage(err))
}

func TestIsIDPNotFoundError(t *testing.T) {
	require.True(t, isIDPNotFoundError(errors.New("not found"), &http.Response{StatusCode: http.StatusNotFound}))
	require.True(t, isIDPNotFoundError(idpErrorWithBody{body: []byte(`{"code":"ENTITY_NOT_FOUND","message":"missing"}`)}, nil))
	require.False(t, isIDPNotFoundError(idpErrorWithBody{body: []byte(`{"code":"INVALID_REQUEST","message":"bad request"}`)}, nil))
}
