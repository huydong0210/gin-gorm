package error

type AppError struct {
	Message string
}

func (err *AppError) Error() string {
	return err.Message
}
