package res

import (
	"github.com/gin-gonic/gin"
	"github.com/mohae/deepcopy"
)

const (
	// SuccessCode 请求成功
	SuccessCode = "2001"
	// UserAlreadyExistsCode 用户已经存在
	UserAlreadyExistsCode = "2002"
	// DuplicateInformationCode 重复信息
	DuplicateInformationCode = "2003"
	// UserNotExistsCode 用户不存在或者密码错误
	UserNotExistsCode = "2004"
	// WrongPasswordCode 密码错误
	WrongPasswordCode = "2005"
	// MissingRequiredParameterCode 缺少必要参数或者参数格式不正确
	MissingRequiredParameterCode = "4000"
	// InvalidTokenCode 缺少token或者token无效
	InvalidTokenCode = "4001"
	// SystemErrorCode 系统错误
	SystemErrorCode = "5000"
)

// ResMap 返回的提示信息
var ResMap = map[string]gin.H{
	SuccessCode: { "code": 2001, "message": "success" },
	UserAlreadyExistsCode: { "code": 2002, "message": "user already exists" },
	DuplicateInformationCode: { "code": 2003, "message": "duplicate information" },
	UserNotExistsCode: { "code": 2004, "message": "the user does not exist" },
	WrongPasswordCode: { "code": 2005, "message": "the password is wrong" },
	MissingRequiredParameterCode: { "code": 4000, "message": "missing required parameter" },
	InvalidTokenCode: { "code": 4001, "message": "missing token or invalid token" },
	SystemErrorCode: { "code": 5000, "message": "system error" },
}

// Resust 响应信息封装函数
func Resust(code string, result interface{}) gin.H {
	var r gin.H = deepcopy.Copy(ResMap[code]).(gin.H)

	r["result"] = result

	return r
}

// Success 成功响应信息
func Success(result interface{}) gin.H {
	var r gin.H = deepcopy.Copy(ResMap[SuccessCode]).(gin.H)

	r["result"] = result

	return r
}

// Error 错误响应信息
func Error(code string) gin.H {
	return ResMap[code]
}
