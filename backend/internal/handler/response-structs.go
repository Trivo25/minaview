package handler

// defines the structure of the response with some meta data and the actual data
type Response struct {
	Error string
	Data  interface{}
}
