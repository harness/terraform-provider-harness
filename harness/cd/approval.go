package cd

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
)

// Helper type for accessing all application related crud methods
type ApprovalClient struct {
	ApiClient *ApiClient
}

// CRUD
func (ac *ApprovalClient) GetApprovalDetails(applicationId string, executionId string) (*graphql.ApprovalDetailsPayload, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($applicationId: String!, $executionId: String!) {
			approvalDetails(applicationId: $applicationId, executionId: $executionId) {
				approvalDetails {
				%s
				}
			}
		}`, approvalDetailsFields),
		Variables: map[string]interface{}{
			"applicationId": applicationId,
			"executionId":   executionId,
		},
	}

	res := struct {
		ApprovalDetails graphql.ApprovalDetailsPayload
	}{}
	ac.ApiClient.ExecuteGraphQLQuery(query, &res)

	return &res.ApprovalDetails, nil
}

func (ac *ApprovalClient) ApproveOrRejectApprovals(input *graphql.ApproveOrRejectApprovalsInput) (*graphql.ApproveOrRejectApprovalsInputPayload, error) {

	query := &GraphQLQuery{
		Query: `mutation approveOrRejectApprovals ($approvalInput: ApproveOrRejectApprovalsInput!) {
			approveOrRejectApprovals(input: $approvalInput)
			 {
			 success
			 clientMutationId
			 }
		   }`,
		Variables: map[string]interface{}{
			"approvalInput": &input,
		},
	}

	res := &struct {
		ApproveOrRejectApprovalsInputPayload *graphql.ApproveOrRejectApprovalsInputPayload
	}{}
	err := ac.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return res.ApproveOrRejectApprovalsInputPayload, nil
}

var approvalDetailsFields = `
approvalId
approvalType
stepName
stageName
startedAt
triggeredBy {
	name
	email
}
willExpireAt
... on UserGroupApprovalDetails {
	approvers
	approvalId
	approvalType
	stepName
	stageName
	startedAt
	executionId
	triggeredBy {
	name
	email
	}
	willExpireAt
	variables {
	name
	value
	}
}
... on ShellScriptDetails {
	approvalId
	approvalType
	retryInterval
	stageName
	stepName
	startedAt
	triggeredBy {
	email
	name
	}
	willExpireAt
}
... on SNOWApprovalDetails {
	approvalCondition
	approvalId
	approvalType
	currentStatus
	rejectionCondition
	stageName
	startedAt
	stepName
	ticketType
	ticketUrl
	triggeredBy {
	email
	name
	}
	willExpireAt
}
... on JiraApprovalDetails {
	approvalCondition
	approvalId
	approvalType
	currentStatus
	issueKey
	issueUrl
	rejectionCondition
	stepName
	stageName
	startedAt
	triggeredBy {
	email
	name
	}
	willExpireAt
}
`
