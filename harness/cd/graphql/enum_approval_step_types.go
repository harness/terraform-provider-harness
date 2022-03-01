package graphql

type ApprovalStepType string

var ApprovalStepTypes = struct {
	Jira        ApprovalStepType
	ServiceNow  ApprovalStepType
	ShellScript ApprovalStepType
	UserGroup   ApprovalStepType
}{
	Jira:        "JIRA",
	ServiceNow:  "SERVICENOW",
	ShellScript: "SHELL_SCRIPT",
	UserGroup:   "USER_GROUP",
}

var ApprovalStepTypeList = []string{
	ApprovalStepTypes.Jira.String(),
	ApprovalStepTypes.ServiceNow.String(),
	ApprovalStepTypes.ShellScript.String(),
	ApprovalStepTypes.UserGroup.String(),
}

func (d ApprovalStepType) String() string {
	return string(d)
}
