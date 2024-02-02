package res

type CommonResponse struct {
	Code    int
	Status  string
	Message string
	Data    any
}

type ErrorResponse struct {
	Code         int
	Status       string
	ErrorMessage string
}
