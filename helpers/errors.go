package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/harness/harness-go-sdk/harness/policymgmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	openapi_client_nextgen "github.com/harness/harness-openapi-go-client/nextgen"
	"google.golang.org/grpc/codes"

	"github.com/harness/harness-go-sdk/harness/dbops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HandleApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	return handleApiError(err, d, httpResp, false)
}

// HandleGitApiError handles API errors specifically related to Git operations
// This provides more specific error messages for common Git-related errors
func HandleGitApiError(err error, d *schema.ResourceData, httpResp *http.Response, connectorRef string, repoName string) diag.Diagnostics {
	if httpResp != nil && httpResp.StatusCode == 400 && err != nil {
		errMsg := err.Error()

		if strings.Contains(errMsg, "No connector found") && strings.Contains(errMsg, connectorRef) {
			return diag.Errorf("Invalid connector reference: %s. Please check if the connector exists and you have the correct permissions.", connectorRef)
		}

		if strings.Contains(errMsg, "Please check the requested file path") && strings.Contains(errMsg, repoName) {
			return diag.Errorf("Invalid repository name: %s. Please check if the repository exists and is accessible.", repoName)
		}
	}

	return HandleApiError(err, d, httpResp)
}

// HandleGitApiErrorWithResourceData extracts Git details from the resource data and handles Git-related API errors
func HandleGitApiErrorWithResourceData(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	connectorRef, hasConnector := d.GetOk("git_details.0.connector_ref")
	repoName, hasRepo := d.GetOk("git_details.0.repo_name")
	if hasConnector && hasRepo {
		return HandleGitApiError(err, d, httpResp, connectorRef.(string), repoName.(string))
	}
	return HandleApiError(err, d, httpResp)
}

func HandleDBOpsApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	erro, ok := err.(dbops.GenericSwaggerError)
	if ok && httpResp != nil {
		if httpResp.StatusCode == 401 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if token has expired or is wrong.\n" +
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}
		if httpResp.StatusCode == 403 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if the token has required permission for this operation.\n" +
				"2) Please check if the token has expired or is wrong.")
		}
		if httpResp.StatusCode == 404 {
			return diag.Errorf("resource with ID %s not found: %v", d.Id(), erro.Error())
		}
	}

	return diag.Errorf(err.Error())
}

func HandlePolicyApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	se, ok := err.(policymgmt.GenericSwaggerError)
	if ok && httpResp != nil {
		if httpResp.StatusCode == 400 {
			// Extract error message from SDK error object instead of re-reading response body
			errorBody := se.Body()
			var jsonMap map[string]interface{}
			if err := json.Unmarshal(errorBody, &jsonMap); err == nil {
				if message, exists := jsonMap["message"]; exists {
					return diag.Errorf("Bad Request: %s", message)
				}
			}
			return diag.Errorf("Bad Request: %s", string(errorBody))
		}
		if httpResp.StatusCode == 401 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if token has expired or is wrong.\n" +
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}
		if httpResp.StatusCode == 403 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if the token has required permission for this operation.\n" +
				"2) Please check if the token has expired or is wrong.")
		}
		if httpResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
	}

	return handleApiError(err, d, httpResp, false)
}

func handleApiError(err error, d *schema.ResourceData, httpResp *http.Response, read bool) diag.Diagnostics {
	erro, ok := err.(nextgen.GenericSwaggerError)
	if ok {
		errMessage := erro.Error()
		gitopsErr, gitopsErrOk := erro.Model().(nextgen.GatewayruntimeError)
		if gitopsErrOk {
			errMessage = gitopsErr.Message
		}

		if httpResp != nil && httpResp.StatusCode == 401 {
			return diag.Errorf("%s", httpResp.Status+"\n"+"Hint:\n"+
				"1) Please check if token has expired or is wrong.\n"+
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}
		if gitopsErrOk && httpResp != nil && httpResp.StatusCode == 403 {
			if len(errMessage) > 0 {
				return diag.Errorf("%s", httpResp.Status+"\n"+"Hint:\n"+
					"1) Please check if the token has required permission for this operation.\n"+
					"2) Please check if the token has expired or is wrong.\n"+
					"3) "+errMessage)
			}
		}
		if httpResp != nil && httpResp.StatusCode == 403 {
			return diag.Errorf("%s", httpResp.Status+"\n"+"Hint:\n"+
				"1) Please check if the token has required permission for this operation.\n"+
				"2) Please check if the token has expired or is wrong.")
		}
		if httpResp != nil && httpResp.StatusCode == 404 {
			// GitOps handling for NotFound
			if gitopsErrOk && read {
				if codes.Code(gitopsErr.Code) == codes.NotFound {
					d.SetId("")
					d.MarkNewResource()
					return nil
				}
			}
			respErrorBody, err := ParseErrorBody(erro)
			if err == nil {
				if read && (respErrorBody.Code == string(nextgen.ErrorCodes.EntityNotFound) || respErrorBody.Code == string(nextgen.ErrorCodes.ResourceNotFound)) {
					d.SetId("")
					d.MarkNewResource()
					return nil
				}
			}
			if read && !gitopsErrOk && (erro.Code() == nextgen.ErrorCodes.EntityNotFound) {
				d.SetId("")
				d.MarkNewResource()
				return nil
			}
			return diag.Errorf("resource with ID %s not found: %v", d.Id(), errMessage)
		}
		if read && !gitopsErrOk {
			respErrorBody, err := ParseErrorBody(erro)
			if err == nil {
				if respErrorBody.Code == string(nextgen.ErrorCodes.EntityNotFound) || respErrorBody.Code == string(nextgen.ErrorCodes.ResourceNotFound) {
					d.SetId("")
					d.MarkNewResource()
					return nil
				}
			}
		}

		// Handle 4xx/5xx errors - extract and display message with HTTP status for better error visibility
		if httpResp != nil && httpResp.StatusCode >= 400 {
			respErrorBody, parseErr := ParseErrorBody(erro)
			if parseErr == nil && respErrorBody != nil && respErrorBody.Message != "" {
				return diag.Errorf("HTTP %d: %s", httpResp.StatusCode, respErrorBody.Message)
			}
		}

		return diag.Errorf(errMessage)
	}

	err_openapi_client, ok := err.(openapi_client_nextgen.GenericSwaggerError)
	if ok {
		if httpResp != nil && httpResp.StatusCode == 401 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if token has expired or is wrong.\n" +
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}
		if httpResp != nil && httpResp.StatusCode == 403 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if the token has required permission for this operation.\n" +
				"2) Please check if the token has expired or is wrong.")
		}
		if httpResp != nil && httpResp.StatusCode == 404 {
			return diag.Errorf("resource with ID %s not found: %v", d.Id(), erro.Error())
		}
		var jsonMap map[string]interface{}
		err := json.Unmarshal(err_openapi_client.Body(), &jsonMap)
		if err == nil {
			return diag.Errorf(jsonMap["message"].(string))
		}
		return diag.Errorf(err_openapi_client.Error())
	}

	return diag.Errorf(err.Error())
}

func ParseErrorBody(err nextgen.GenericSwaggerError) (*nextgen.ModelError, error) {
	// First unmarshal into a generic map so we can normalize types that the
	// SDK expects as strings but the platform sometimes returns as numbers
	// (for example: "failure": { "code": 400 }). We convert those numeric
	// values to strings and then unmarshal into nextgen.ModelError.
	var raw map[string]interface{}
	if uerr := json.Unmarshal(err.Body(), &raw); uerr != nil {
		return nil, uerr
	}

	// Normalize nested failure.code if present
	if f, ok := raw["failure"].(map[string]interface{}); ok {
		if codeVal, ok2 := f["code"]; ok2 && codeVal != nil {
			switch v := codeVal.(type) {
			case float64:
				// JSON numbers are decoded as float64 by default
				f["code"] = fmt.Sprintf("%.0f", v)
			default:
				// leave as-is (handles string already)
				f["code"] = fmt.Sprintf("%v", v)
			}
		}
	}

	// Also normalize top-level code if present
	if codeVal, ok := raw["code"]; ok && codeVal != nil {
		switch v := codeVal.(type) {
		case float64:
			raw["code"] = fmt.Sprintf("%.0f", v)
		default:
			raw["code"] = fmt.Sprintf("%v", v)
		}
	}

	modified, merr := json.Marshal(raw)
	if merr != nil {
		return nil, merr
	}

	var parsed nextgen.ModelError
	if uerr := json.Unmarshal(modified, &parsed); uerr != nil {
		return nil, uerr
	}
	return &parsed, nil
}

func HandleReadApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	return handleApiError(err, d, httpResp, true)
}

func HandleDBOpsReadApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	_, ok := err.(dbops.GenericSwaggerError)
	if ok && httpResp != nil {
		if httpResp.StatusCode == 401 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if token has expired or is wrong.\n" +
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}
		if httpResp.StatusCode == 403 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if the token has required permission for this operation.\n" +
				"2) Please check if the token has expired or is wrong.")
		}
		if httpResp.StatusCode == 404 {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}
	return diag.Errorf(err.Error())
}
