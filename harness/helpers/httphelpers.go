package helpers

type HTTPHeader string

func (h HTTPHeader) String() string {
	return string(h)
}

var HTTPHeaders = struct {
	Accept          HTTPHeader
	ApplicationJson HTTPHeader
	ContentType     HTTPHeader
	UserAgent       HTTPHeader
	ApiKey          HTTPHeader
	Authorization   HTTPHeader
}{
	Accept:          "Accept",
	ApplicationJson: "application/json; charset=utf-8",
	ContentType:     "Content-Type",
	UserAgent:       "User-Agent",
	ApiKey:          "X-Api-Key",
	Authorization:   "Authorization",
}

type QueryParameter string

func (q QueryParameter) String() string {
	return string(q)
}

var QueryParameters = struct {
	AccountId     QueryParameter
	ApplicationId QueryParameter
	FilePaths     QueryParameter
	Limit         QueryParameter
	Offset        QueryParameter
	Type          QueryParameter
	YamlFilePath  QueryParameter
}{
	AccountId:     "accountId",
	ApplicationId: "appId",
	FilePaths:     "filePaths",
	Limit:         "limit",
	Offset:        "offset",
	Type:          "type",
	YamlFilePath:  "yamlFilePath",
}

var QueryParametersExecutions = struct {
	ApplicationId QueryParameter
	EnvironmentId QueryParameter
	Limit         QueryParameter
	SortField     QueryParameter
	SortDirection QueryParameter
	RoutingId     QueryParameter
}{
	ApplicationId: "appId",
	EnvironmentId: "envId",
	Limit:         "limit",
	SortField:     "sort[0][field]",
	SortDirection: "sort[0][direction]",
	RoutingId:     "routingId",
}
