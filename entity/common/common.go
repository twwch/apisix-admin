package common

import "errors"

var ParamsError = errors.New("参数错误")
var SystemError = errors.New("系统错误，请稍后重试！")

type Empty struct {

}
