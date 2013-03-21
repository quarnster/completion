package content

// A common function for creating error responses.
func NewErrorResponse(s string) *Response {
	r := NewResponse()
	r.Data.Set("error", s)
	return &r
}
