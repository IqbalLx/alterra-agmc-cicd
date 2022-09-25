package errors

type InternalServerError struct {
	Origin  string
	Message string
}

type ClientError struct {
	Code    int
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}

func (e *ClientError) Error() string {
	return e.Message
}
