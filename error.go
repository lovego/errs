package errs

type Error struct {
	code, message string
	data          interface{}
	logError      bool
}

func New(code, message string) Error {
	return Error{code: code, message: message}
}

func New2(code, message string) Error {
	return Error{code: code, message: message, logError: true}
}

func Make(code, message string, data interface{}) Error {
	return Error{code: code, message: message, data: data}
}

func (err Error) Code() string {
	return err.code
}

func (err Error) Message() string {
	return err.message
}

func (err Error) Data() interface{} {
	return err.data
}

func (err *Error) LogError() bool {
	return err.logError
}

func (err Error) Error() string {
	return err.code + `: ` + err.message
}

func (err *Error) SetData(data interface{}) {
	err.data = data
}
