package domain

type RestError struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}
