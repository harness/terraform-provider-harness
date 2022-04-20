package graphql

type PageInfo struct {
	HasMore bool
	Limit   int
	Offset  int
	Total   int
}
