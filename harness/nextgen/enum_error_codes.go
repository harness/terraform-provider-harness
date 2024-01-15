package nextgen

type ErrorCode string

var ErrorCodes = struct {
	ResourceNotFound ErrorCode
	EntityNotFound   ErrorCode
}{
	ResourceNotFound: "RESOURCE_NOT_FOUND_EXCEPTION",
	EntityNotFound:   "ENTITY_NOT_FOUND",
}
