package error_code

const (
	CODE_AUTH_TOKEN_EXPIRE = 29
)

var (
	errCodeMessage = map[int]string{
		CODE_AUTH_TOKEN_EXPIRE: "token已过期，请重新登录",
	}
)

func ErrCodeMessage(code int) string  {
	return errCodeMessage[code]
}