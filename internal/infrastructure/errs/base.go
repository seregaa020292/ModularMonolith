package errs

type (
	ErrorCustomer interface {
		error
		OriginalError() error
		StatusCode() int
	}
	BaseError struct {
		msg  string
		err  error
		code int
	}
)

func NewBaseError(msg string, err error, statusCode int) error {
	return &BaseError{
		msg:  msg,
		err:  err,
		code: statusCode,
	}
}

func (e *BaseError) Error() string {
	return e.msg
}

func (e *BaseError) OriginalError() error {
	return e.err
}

func (e *BaseError) StatusCode() int {
	return e.code
}
