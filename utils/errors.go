package utils

type Error struct {
	code    string
	message string
}

type ErrorList struct {
	SignupSuccessfully *Error
	SigninSuccessfully *Error
	UsernameDuplicate  *Error
	PasswordEmpty      *Error
	UsernameEmpty      *Error
	UsernameNotExisted *Error
	PasswordNotMatched *Error
	SystemError        *Error
	Success            *Error
	Failed             *Error
	NotFound           *Error
}

var ErrorConstants = NewErrorList()

func (err Error) Code() string {
	return err.code
}

func (err Error) Message() string {
	return err.message
}

func NewErrorList() *ErrorList {
	user0001 := &Error{"USER001", "Sign-up successfully"}
	user0002 := &Error{"USER002", "Username is duplicate"}
	user0003 := &Error{"USER003", "Password is empty"}
	user0004 := &Error{"USER004", "Username is empty"}
	user0005 := &Error{"USER005", "Username is not existed"}
	user0006 := &Error{"USER006", "Password is not matched"}
	user0007 := &Error{"USER007", "Sign-in successfully"}
	sys0001 := &Error{"SYS001", ""}
	success := &Error{"Success", ""}
	failed := &Error{"Failed", ""}
	notFound := &Error{"NotFound", "URL is not found"}

	return &ErrorList{
		SignupSuccessfully: user0001,
		SigninSuccessfully: user0007,
		UsernameDuplicate:  user0002,
		PasswordEmpty:      user0003,
		UsernameEmpty:      user0004,
		UsernameNotExisted: user0005,
		PasswordNotMatched: user0006,
		SystemError:        sys0001,
		Success:            success,
		Failed:             failed,
		NotFound:           notFound,
	}
}
