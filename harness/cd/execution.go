package cd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/executions"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

type ExecutionClient struct {
	ApiClient *ApiClient
}

func (c *ExecutionClient) AbortOrCancelWorkflowOrPipelineById(id, action, appId, envId string) error {
	c.ApiClient.Log.Debugf("%s workflow or pipeline by id: %s", action, id)

	var requestBody bytes.Buffer

	body := []*struct {
		ExeuctionInterruptType string `json:"executionInterruptType"`
	}{
		{
			ExeuctionInterruptType: action,
		},
	}

	// JSON encode our body payload
	if err := json.NewEncoder(&requestBody).Encode(body); err != nil {
		return err
	}

	req, err := c.ApiClient.NewAuthorizedRequest(path.Join(utils.DefaultCDApiUrl, "/executions", id), http.MethodPut, &requestBody)

	if err != nil {
		return err
	}

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(helpers.QueryParametersExecutions.ApplicationId.String(), appId)
	q.Add(helpers.QueryParametersExecutions.EnvironmentId.String(), envId)
	q.Add(helpers.QueryParametersExecutions.RoutingId.String(), c.ApiClient.Configuration.AccountId)
	q.Add(helpers.QueryParametersExecutions.SortField.String(), "createdAt")
	q.Add(helpers.QueryParametersExecutions.SortDirection.String(), "DESC")
	q.Add(helpers.QueryParametersExecutions.Limit.String(), "10")
	req.URL.RawQuery = q.Encode()

	_, err = c.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *ExecutionClient) ExportExecutions(input *graphql.ExportExecutionsInput) (*graphql.ExportExecutionsPayload, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: ExportExecutionsInput!) {
			exportExecutions(input: $input) {
				%s
			}
		}`, exportExecutionsFields),
		Variables: map[string]interface{}{
			"input": &input,
		},
	}

	res := struct {
		ExportExecutions graphql.ExportExecutionsPayload
	}{}
	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)
	if err != nil {
		return nil, err
	}

	return &res.ExportExecutions, err
}

func (c *ExecutionClient) GetWorkflowExecutionById(id string) (*graphql.WorkflowExecution, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($input: String!) {
			execution(executionId: $input) {
				... on WorkflowExecution {
					%s
				}
			}
		}`, workflowExecutionFields),
		Variables: map[string]interface{}{
			"input": id,
		},
	}

	res := struct {
		Execution graphql.WorkflowExecution
	}{}
	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		if strings.Contains(err.Error(), "execution does not exist") {
			return nil, nil
		}
		return nil, err
	} else if res.Execution.Id == "" {
		return nil, nil
	}

	return &res.Execution, nil
}

func (c *ExecutionClient) GetPipelineExecutionById(id string) (*graphql.PipelineExecution, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($input: String!) {
			execution(executionId: $input) {
				... on PipelineExecution {
					%s
				}
			}
		}`, pipelineExecutionFields),
		Variables: map[string]interface{}{
			"input": id,
		},
	}

	res := struct {
		Execution graphql.PipelineExecution
	}{}
	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		if strings.Contains(err.Error(), "execution does not exist") {
			return nil, nil
		}
		return nil, err
	} else if res.Execution.Id == "" {
		return nil, nil
	}

	return &res.Execution, nil
}

func (c *ExecutionClient) StartExecution(input *graphql.StartExecutionInput) (*graphql.StartExecutionPayload, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: StartExecutionInput!) {
			startExecution(input: $input) {
				clientMutationId
				warningMessage
				workflowExecution: execution {
					... on WorkflowExecution {
						%s
					}
				}
				pipelineExecution:execution {
					... on PipelineExecution {
						%s
					}
				}
			}
		}`, workflowExecutionFields, pipelineExecutionFields),
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	res := struct {
		StartExecution graphql.StartExecutionPayload
	}{}
	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	if res.StartExecution.WorkflowExecution.Id == "" {
		res.StartExecution.WorkflowExecution = nil
	}

	if res.StartExecution.PipelineExecution.Id == "" {
		res.StartExecution.PipelineExecution = nil
	}

	return &res.StartExecution, nil
}

func (c *ExecutionClient) ExecuteRequest(request *retryablehttp.Request) (*executions.ExecutionItem, error) {

	res, err := c.ApiClient.Configuration.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	if ok, err := checkStatusCode(res.StatusCode); !ok {
		return nil, err
	}

	defer res.Body.Close()

	// Make sure we can parse the body properly
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}

	responseObj := &executions.Response{}

	// Unmarshal into our response object
	if err := json.NewDecoder(&buf).Decode(&responseObj); err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	if responseObj.IsEmpty() {
		return nil, errors.New("received an empty response")
	}

	if len(responseObj.ResponseMessages) > 0 {
		return nil, responseObj.ResponseMessages[0].ToError()
	}

	return responseObj.Resource, nil
}

var workflowExecutionFields = `
id
application {
	id
	name
}
artifacts {
	artifactSource {
		createdAt
		id
		name
	}
	buildNo
	collectedAt
	id
}
rollbackArtifacts {
	artifactSource {
		createdAt
		id
		name
	}
	buildNo
	collectedAt
	id
}
createdAt
endedAt
failureDetails
id
notes
startedAt
status
tags {
	name
	value
}
`

var pipelineExecutionFields = `
id
application {
	id
	name
}
pipelineStageExecutions {
	pipelineStepName
	pipelineStageName
	pipelineStageElementId
	... on ApprovalStageExecution {
		approvalStepType
		status
	}
	... on WorkflowStageExecution {
		runtimeInputVariables {
			allowMultipleValues
			allowedValues
			defaultValue
			fixed
			name
			required
			type
		}
		workflowExecutionId
		status
	}
}
createdAt
endedAt
failureDetails
id
notes
startedAt
status
tags {
	name
	value
}
`

const exportExecutionsFields = `
clientMutationId
requestId
status
totalExecutions
triggeredAt
statusLink
downloadLink
expiresAt
errorMessage
`
