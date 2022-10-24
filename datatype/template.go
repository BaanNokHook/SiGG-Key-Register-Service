package datatype

// Template implements the keyregister interface for a template error
type Template struct {
	AbsPath  *string         `json:"absPath,omitempty"`
	Filename *string         `json:"filename,omitempty"`
	LineNo   *int            `json:"lineNo,omitempty"`
	Context  *[]FrameContext `json:"context,omitempty"`
}
