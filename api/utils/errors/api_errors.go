package errors

type ApiErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
