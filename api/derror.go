package api



type APIError struct {
	ErrorMessage string
	StatusCode int
}

func (err APIError) Error() string {
	return string(err.StatusCode)+": "+err.ErrorMessage
}

func (err APIError) IsClientError() bool {
	return err.StatusCode >= 400 && err.StatusCode < 500
}

func (err APIError) IsClientError() bool {
	return err.StatusCode >= 500 && err.StatusCode < 600
}
