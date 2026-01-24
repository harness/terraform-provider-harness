package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
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
	var parsed nextgen.ModelError
	if err := json.Unmarshal(err.Body(), &parsed); err != nil {
		return nil, err
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

// HandleChaosApiError handles errors from Chaos Engineering SDK
// This provides detailed error messages extracted from the API response
func HandleChaosApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	return handleChaosApiError(err, d, httpResp, false)
}

// HandleChaosReadApiError handles read errors from Chaos Engineering SDK
// For 404 errors during read operations, it clears the resource from state
func HandleChaosReadApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	return handleChaosApiError(err, d, httpResp, true)
}

// HandleChaosReadApiErrorWithGracefulDestroy handles read errors from Chaos Engineering SDK
// For 404 errors and certain 500 errors (resource not found/inconsistent state) during read,
// it clears the resource from state. This is useful during destroy operations when
// dependent resources may have already been deleted.
//
// gracefulErrorPatterns: list of error message patterns that should be treated as "not found"
// Example: []string{"no matching infra", "at least one hub is required"}
func HandleChaosReadApiErrorWithGracefulDestroy(err error, d *schema.ResourceData, httpResp *http.Response, gracefulErrorPatterns []string) diag.Diagnostics {
	chaosErr, ok := err.(chaos.GenericSwaggerError)
	if !ok {
		return diag.Errorf(err.Error())
	}
	
	// Handle 404 - resource not found
	if httpResp != nil && httpResp.StatusCode == 404 {
		log.Printf("[DEBUG] Resource not found (404), removing from state")
		d.SetId("")
		return nil
	}
	
	// Handle 500 errors that indicate resource is gone or in inconsistent state
	if httpResp != nil && httpResp.StatusCode == 500 && len(gracefulErrorPatterns) > 0 {
		errMsg := ""
		
		// Try to extract error message from model
		if chaosErr.Model() != nil {
			if apiErr, ok := chaosErr.Model().(chaos.ApiRestError); ok {
				if apiErr.Description != "" {
					errMsg = apiErr.Description
				} else if apiErr.Message != "" {
					errMsg = apiErr.Message
				}
			}
		}
		
		// Fallback to error string
		if errMsg == "" {
			errMsg = fmt.Sprintf("%v", err)
		}
		
		// Check if error matches any graceful patterns
		for _, pattern := range gracefulErrorPatterns {
			if strings.Contains(strings.ToLower(errMsg), strings.ToLower(pattern)) {
				log.Printf("[WARN] Resource not found or in inconsistent state during read (matched pattern: %s): %v", pattern, err)
				d.SetId("")
				return nil
			}
		}
	}
	
	// Not a graceful error, handle normally
	return handleChaosApiError(err, d, httpResp, true)
}

// handleChaosApiError is the internal implementation for chaos error handling
// It extracts detailed error messages from the error body or model
func handleChaosApiError(err error, d *schema.ResourceData, httpResp *http.Response, read bool) diag.Diagnostics {
	chaosErr, ok := err.(chaos.GenericSwaggerError)
	if !ok {
		// Not a chaos error, fallback to generic error message
		return diag.Errorf(err.Error())
	}
	if httpResp == nil {
		return diag.Errorf(chaosErr.Error())
	}

	// Debug: Log the raw error body
	log.Printf("[DEBUG] Chaos API Error - Status: %d, Body: %s", httpResp.StatusCode, string(chaosErr.Body()))
	if chaosErr.Model() != nil {
		log.Printf("[DEBUG] Chaos API Error - Model type: %T, Model: %+v", chaosErr.Model(), chaosErr.Model())
	} else {
		log.Printf("[DEBUG] Chaos API Error - Model is nil")
	}

	// Extract detailed error message from model or body
	var errorMessage string

	// Try to get error from model first (for 500 errors, SDK decodes ApiRestError)
	if chaosErr.Model() != nil {
		if apiErr, ok := chaosErr.Model().(chaos.ApiRestError); ok {
			// Prioritize Description as it often contains more detailed error info
			if apiErr.Description != "" {
				errorMessage = apiErr.Description
			} else if apiErr.Message != "" {
				errorMessage = apiErr.Message
			}
		}
	}

	// Fallback to parsing body if model didn't have a message
	if errorMessage == "" && len(chaosErr.Body()) > 0 {
		var jsonMap map[string]interface{}
		if jsonErr := json.Unmarshal(chaosErr.Body(), &jsonMap); jsonErr == nil {
			// Try different common error message fields
			if msg, exists := jsonMap["message"]; exists {
				errorMessage = fmt.Sprintf("%v", msg)
			} else if desc, exists := jsonMap["description"]; exists {
				errorMessage = fmt.Sprintf("%v", desc)
			} else if errMsg, exists := jsonMap["error"]; exists {
				errorMessage = fmt.Sprintf("%v", errMsg)
			} else if details, exists := jsonMap["details"]; exists {
				errorMessage = fmt.Sprintf("%v", details)
			}
		}
	}

	// Fallback to HTTP status if no detailed message found
	if errorMessage == "" {
		// If we have a body but couldn't parse it, include it
		if len(chaosErr.Body()) > 0 && len(chaosErr.Body()) < 500 {
			errorMessage = fmt.Sprintf("%s (response: %s)", httpResp.Status, string(chaosErr.Body()))
		} else if len(chaosErr.Body()) > 0 {
			errorMessage = fmt.Sprintf("%s (response body too large: %d bytes)", httpResp.Status, len(chaosErr.Body()))
		} else {
			errorMessage = httpResp.Status + " (no error details provided by server)"
		}
		log.Printf("[DEBUG] Chaos API Error - No structured error message found, using: %s", errorMessage)
	}

	// Handle specific status codes with helpful hints
	switch httpResp.StatusCode {
	case 400:
		return diag.Errorf("Bad Request: %s", errorMessage)

	case 401:
		return diag.Errorf("%s\n\nHint:\n"+
			"1) Please check if token has expired or is wrong.\n"+
			"2) Harness Provider is misconfigured. Please provide the correct platform_api_key.",
			errorMessage)

	case 403:
		return diag.Errorf("%s\n\nHint:\n"+
			"1) Please check if the token has required permission for this operation.\n"+
			"2) Please check if the token has expired or is wrong.",
			errorMessage)

	case 404:
		if read {
			// For read operations, clear the resource from state
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		resourceID := "unknown"
		if d != nil && d.Id() != "" {
			resourceID = d.Id()
		}
		return diag.Errorf("Resource with ID %s not found: %s", resourceID, errorMessage)

	case 500:
		return diag.Errorf("Internal Server Error: %s", errorMessage)

	default:
		// For other status codes, return the error message with status
		return diag.Errorf("%s: %s", httpResp.Status, errorMessage)
	}
}

// HandleChaosGraphQLError handles errors from Chaos GraphQL APIs
// This provides better error messages by extracting details from GraphQL errors
func HandleChaosGraphQLError(err error, d *schema.ResourceData, operation string) diag.Diagnostics {
	return handleChaosGraphQLError(err, d, operation, false)
}

// HandleChaosGraphQLReadError handles read errors from Chaos GraphQL APIs
// For not found errors during read operations, it clears the resource from state
func HandleChaosGraphQLReadError(err error, d *schema.ResourceData, operation string) diag.Diagnostics {
	return handleChaosGraphQLError(err, d, operation, true)
}

// handleChaosGraphQLError is the internal implementation for Chaos GraphQL error handling
// It extracts meaningful error messages from Chaos GraphQL errors and provides helpful hints
func handleChaosGraphQLError(err error, d *schema.ResourceData, operation string, read bool) diag.Diagnostics {
	if err == nil {
		return nil
	}

	// Get the full error message for logging and analysis
	fullErrorMessage := err.Error()
	errorMessage := fullErrorMessage

	// Log the complete raw error for debugging
	log.Printf("[DEBUG] GraphQL Error (Raw) - Operation: %s, Full Error: %s", operation, fullErrorMessage)

	// Try to extract GraphQL error details
	// Pattern: "graphql error: <message> (path: [<path>])"
	if strings.Contains(errorMessage, "graphql error:") {
		// Extract the actual error message
		parts := strings.Split(errorMessage, "graphql error:")
		if len(parts) > 1 {
			message := strings.TrimSpace(parts[len(parts)-1])

			// Remove the path suffix if present
			if idx := strings.Index(message, " (path:"); idx > 0 {
				message = message[:idx]
			}

			errorMessage = message
		}
	}

	// Remove redundant wrapping (e.g., "failed to update: failed to update:")
	operationPrefix := "failed to " + operation + ": "
	for strings.HasPrefix(errorMessage, operationPrefix) {
		errorMessage = strings.TrimPrefix(errorMessage, operationPrefix)
	}

	// Log the cleaned error message
	log.Printf("[DEBUG] GraphQL Error (Cleaned) - Operation: %s, Message: %s", operation, errorMessage)

	// Check if there's any additional context we might have missed
	// Only warn if the cleaned message is suspiciously short (less than 10 chars)
	// or if it's empty, as this might indicate we're losing actual error content
	if len(errorMessage) < 10 && len(fullErrorMessage) > 20 {
		log.Printf("[WARN] GraphQL Error - Cleaned error message is very short. Original: %q, Cleaned: %q", fullErrorMessage, errorMessage)
	} else if len(errorMessage) == 0 {
		log.Printf("[ERROR] GraphQL Error - Error message was completely removed during cleaning. Original: %q", fullErrorMessage)
		// Fall back to the full error if we accidentally removed everything
		errorMessage = fullErrorMessage
	}

	// Provide helpful hints based on error message
	errorLower := strings.ToLower(errorMessage)

	switch {
	case strings.Contains(errorLower, "duplicate key"):
		return diag.Errorf("Resource already exists: %s. Use 'terraform import' to manage existing resources.", errorMessage)

	case strings.Contains(errorLower, "not found"):
		if read {
			// For read operations, clear the resource from state
			log.Printf("[WARN] Resource not found during read, removing from state: %s", errorMessage)
			d.SetId("")
			return nil
		}
		resourceID := "unknown"
		if d != nil && d.Id() != "" {
			resourceID = d.Id()
		}
		return diag.Errorf("Resource with ID %s not found: %s", resourceID, errorMessage)

	case strings.Contains(errorLower, "permission denied") || strings.Contains(errorLower, "unauthorized"):
		return diag.Errorf("Permission denied: %s. Verify your API key has the required permissions and the correct scope (account/org/project).", errorMessage)

	case strings.Contains(errorLower, "internal system error") || strings.Contains(errorLower, "internal server error"):
		return diag.Errorf("Internal server error: %s. Please check your configuration and try again. If the issue persists, contact Harness support.", errorMessage)

	case strings.Contains(errorLower, "validation") || strings.Contains(errorLower, "invalid"):
		return diag.Errorf("Validation error: %s. Please check your resource configuration.", errorMessage)

	case strings.Contains(errorLower, "timeout"):
		return diag.Errorf("Request timeout: %s. The operation took too long to complete. Please try again.", errorMessage)

	default:
		// For other errors, return with operation context
		return diag.Errorf("Failed to %s: %s", operation, errorMessage)
	}
}
