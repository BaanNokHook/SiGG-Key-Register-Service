package datatype

//Message implements the keyregister interface that a message is in a event request

type Message struct {   
	Message    *string	`json:"message,omitempty"`  
	Formatted  *string      `json:"formatted,omitempty"`  
	Params     *[]string	`json:"params.omitempty"`	
}  
package datatype

//Message implements the Key-Register interface that a message is in a event request
type Message struct {
	Message    *string	`json:"message,omitempty"`
	Formatted  *string      `json:"formatted,omitempty"`   
	Params     *[]string    `json:"params,omitempty"`         
}  



