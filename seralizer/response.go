package seralizer

type Response struct {
	Code   int         `json:"code"`
	Result string      `json:"result"`
	Data   interface{} `json:"data,omitempty"`
}

const (
	Sucess        = 1
	InvaildParams = 2
	NotExistUser  = 3

	ErrorDelete   = 4
	ErrorSave     = 5
	ErrorGet      = 6
	ErrorRegister = 7
	ErrorLogin    = 8
	ErrorPassword = 9

	ExistUser = 10

	ErrorAuth = 11

	ErrorCreatToken = 12

	ErrorCheckToken = 401

	ErrorTokenTimeout = 402
)

func GetResult(code int) string {
	result := "无效code"
	switch code {
	case Sucess:
		result = "成功"
	case InvaildParams:
		result = "参数错误"
	case NotExistUser:
		result = "不存在用户"
	case ErrorDelete:
		result = "删除失败"
	case ErrorSave:
		result = "保存失败"
	case ErrorGet:
		result = "获取失败"
	case ErrorRegister:
		result = "注册失败"
	case ErrorLogin:
		result = "登录失败"
	case ExistUser:
		result = "用户已经存在"
	case ErrorAuth:
		result = "授权失败"
	case ErrorCreatToken:
		result = "创建token失败"
	case ErrorCheckToken:
		result = "检查token失败"
	case ErrorTokenTimeout:
		result = "token已经超时"
	case ErrorPassword:
		result = "密码错误"
	}

	return result
}
