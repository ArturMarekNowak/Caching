package responses

type HttpError struct {
	Message string
}

func NewHttpError(message string) *HttpError {
	he := new(HttpError)
	he.Message = message
	return he
}
