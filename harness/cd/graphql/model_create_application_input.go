package graphql

type CreateApplicationInput struct {
	AreWebhookSecretsMandated bool   `json:"areWebhookSecretsMandated,omitempty"`
	ClientMutationId          string `json:"clientMutationId,omitempty"`
	Description               string `json:"description,omitempty"`
	IsManualTriggerAuthorized bool   `json:"isManualTriggerAuthorized,omitempty"`
	Name                      string `json:"name,omitempty"`
}
