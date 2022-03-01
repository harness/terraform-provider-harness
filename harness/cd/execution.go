package cd

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
)

type ExecutionClient struct {
	ApiClient *ApiClient
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
