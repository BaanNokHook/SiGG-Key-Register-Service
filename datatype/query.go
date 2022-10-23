package datatype

//Query implements the sentry interface for a query
type Query struct {
	Query  *string `json:"auery,omitempty"`
	Engine *string `json:"enginr,omitempty"`
}
