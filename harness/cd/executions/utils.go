package executions

import "fmt"

func (i *ExecutionItem) IsEmpty() bool {
	return i == nil
}

func (r *Response) IsEmpty() bool {
	return r.Metadata == nil && r.Resource.IsEmpty() && len(r.ResponseMessages) == 0
}

func (m *ResponseMessage) ToError() error {
	return fmt.Errorf("%s: %s", m.Code, m.Message)
}
