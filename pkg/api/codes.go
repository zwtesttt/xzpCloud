package api

import "errors"

const (

	//通用模块
	Failed       = 0 //请求失败
	NotFound     = 0
	Success      = 1    //请求成功
	InvalidToken = 1001 // 无效token
	Unauthorized = 401

	//User 模块
	UserLoginPasswordIncorrect = 10001 // 用户名或密码错误
	UserExists                 = 10002 // 用户已存在

	//Order 模块
	OrderStatusInvalid = 20001 // 订单状态错误
)

var codes = map[int]string{
	NotFound:                   StrNotFound,
	UserLoginPasswordIncorrect: StrUserLoginPasswordIncorrect,
	InvalidToken:               StrInvalidToken,
	UserExists:                 StrUserExists,
	Unauthorized:               StrUnauthorized,
	OrderStatusInvalid:         StrOrderStatusInvalid,
}

var ierrors = map[string]int{
	StrNotFound:                   NotFound,
	StrUserLoginPasswordIncorrect: UserLoginPasswordIncorrect,
	StrInvalidToken:               InvalidToken,
	StrUserExists:                 UserExists,
	StrUnauthorized:               Unauthorized,
	StrOrderStatusInvalid:         OrderStatusInvalid,
}

const (
	StrNotFound     = "not found"
	StrInvalidToken = "无效token"
	StrUnauthorized = "Unauthorized"

	// User 模块
	StrUserExists                 = "用户已存在"
	StrUserLoginPasswordIncorrect = "用户名或密码错误"

	// Order 模块
	StrOrderStatusInvalid = "订单状态错误"
)

func FindCodeError(code int) error {
	if msg, ok := codes[code]; ok {
		return errors.New(msg)
	}
	return errors.New("code not found")
}

func FindErrorsCode(err error) int {
	if code, ok := ierrors[err.Error()]; ok {
		return code
	}
	return NotFound
}
