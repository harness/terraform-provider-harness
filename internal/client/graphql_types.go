package client

type GraphQLResponse struct {
	Data             GraphQLResponseData      `json:"data"`
	Metadata         interface{}              `json:"metadata"`
	Resource         string                   `json:"resource"`
	ResponseMessages []GraphQLResponseMessage `json:"responseMessages"`
	Errors           []GraphQLError           `json:"errors"`
}

type GraphQLError struct {
	Message   string `json:"message"`
	Locations []struct {
		Line   int      `json:"line"`
		Column int      `json:"column"`
		Path   []string `json:"path"`
	} `json:"column"`
}

type GraphQLResponseMessage struct {
	Code         string   `json:"code"`
	Level        string   `json:"level"`
	Message      string   `json:"message"`
	Exception    string   `json:"exception"`
	FailureTypes []string `json:"failureTypes"`
}

type GraphQLResponseData struct {
	Application       *Application              `json:"application"`
	ApplicationByName *Application              `json:"applicationByName"`
	CreateApplication *CreateApplicationPayload `json:"createApplication"`
	DeleteApplication *DeleteApplicationPayload `json:"deleteApplication"`
	UpdateApplication *UpdateApplicationPayload `json:"updateApplication"`
}
