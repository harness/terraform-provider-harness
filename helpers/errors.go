package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	openapi_client_nextgen "github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HandleApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	erro, ok := err.(nextgen.GenericSwaggerError)
	if ok {
		if httpResp != nil && httpResp.StatusCode == 401 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if token has expired or is wrong.\n" +
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}
		if httpResp != nil && httpResp.StatusCode == 403 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint: " +
				"Please check if the token has required permission for this operation.\n")
		}
		return diag.Errorf(erro.Error())
	}

	err_openapi_client, ok := err.(openapi_client_nextgen.GenericSwaggerError)
	if ok {
		if httpResp != nil && httpResp.StatusCode == 401 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if token has expired or is wrong.\n" +
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}

		if httpResp != nil && httpResp.StatusCode == 403 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint: " +
				"Please check if the token has required permission for this operation.\n")
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

func HandleReadApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	erro, ok := err.(nextgen.GenericSwaggerError)
	if ok {
		if httpResp != nil && httpResp.StatusCode == 401 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if token has expired or is wrong.\n" +
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}
		if httpResp != nil && httpResp.StatusCode == 403 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint: " +
				"Please check if the token has required permission for this operation.\n")
		}
		if erro.Model() != nil {
			if erro.Code() == nextgen.ErrorCodes.ResourceNotFound {
				d.SetId("")
				d.MarkNewResource()
				return nil
			}
		}
		return diag.Errorf(erro.Error())
	}

	err_openapi_client, ok := err.(openapi_client_nextgen.GenericSwaggerError)
	if ok {
		if httpResp != nil && httpResp.StatusCode == 401 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint:\n" +
				"1) Please check if token has expired or is wrong.\n" +
				"2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.")
		}
		if httpResp != nil && httpResp.StatusCode == 403 {
			return diag.Errorf(httpResp.Status + "\n" + "Hint: " +
				"Please check if the token has required permission for this operation.\n")
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
