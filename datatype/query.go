package datatype

//Query implements the keyregister interface for a query
type Query struct {
	Query  *string `json:"auery,omitempty"`
	Engine *string `json:"enginr,omitempty"`
}
