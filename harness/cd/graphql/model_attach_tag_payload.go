package graphql

type AttachTagPayload struct {
	ClientMutationId string   `json:"clientMutationId,omitempty"`
	TagLink          *TagLink `json:"tagLink,omitempty"`
}
