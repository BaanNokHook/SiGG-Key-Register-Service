package keyregister

// QueryArgs is a simple internal interface
type QueryArgs interface {
	ToQueryString() string
}
