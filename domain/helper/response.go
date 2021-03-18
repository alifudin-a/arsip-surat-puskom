package helper

// Response custom
type Response struct {
	Code    int64                  `json:"code"`
	Message string                 `json:"message,omitempty"`
	Body    map[string]interface{} `json:"body,omitempty"`
}
