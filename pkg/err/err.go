package errs

/*
用户相关错误码 101
*/
var (
	ErrorAuth    = NewError(101001, "token校验失败")
	ErrorAccount = NewError(101002, "用户名或密码错误")
)


type Error struct {
	Code int    // 错误码
	Msg  string // 错误信息
}

func (e Error) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}
