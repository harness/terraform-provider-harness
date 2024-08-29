package helpers

import (
	"encoding/json"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	openapi_client_nextgen "github.com/harness/harness-openapi-go-client/nextgen"
	"google.golang.org/grpc/codes"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/dbops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HandleApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	return handleApiError(err, d, httpResp, false)
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
			return diag.Errorf("resource with ID %s not found: %v", d.Id(), errMessage)
		}
		if read {
			if erro.Model() != nil && (erro.Code() == nextgen.ErrorCodes.ResourceNotFound || erro.Code() == nextgen.ErrorCodes.EntityNotFound) {
				d.SetId("")
				d.MarkNewResource()
				return nil
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
