package nextgen

type ErrorCode string

var ErrorCodes = struct {
	ResourceNotFound ErrorCode
}{
	ResourceNotFound: "RESOURCE_NOT_FOUND_EXCEPTION",
}
