package xerr

import "github.com/pkg/errors"

var message map[int64]string
var code map[string]int64

func init() {
	message = make(map[int64]string)
	message[SystemError] = "系统报错"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "参数不足"
	message[DBError] = "数据库异常"
	message[DBUpdateAffectedZeroError] = "数据库更新失败"
	message[ErrorReasonDataNotFound] = "数据查找失败"
	message[ErrorReasonRequestParamMissing] = "请求参数缺失"
	message[ErrorReasonRepeatedDo] = "重复操作"
	message[ErrorReasonIllegalOperation] = "非法操作"
	message[ErrorReasonBeyondLimitRange] = "超出限制范围"
	message[ErrorAppId] = "appId无效"
	message[ErrorSecretKey] = "SecretKey错误"
	message[ErrorAuth] = "签名错误"
	code = make(map[string]int64)
	for errCode, s := range message {
		code[s] = errCode
	}
}

func MapErrMsg(errCode int64) string {
	if msg, ok := message[errCode]; ok {
		return msg
	}
	return message[ServerCommonError]
}

func MapErrCode(msg string) int64 {
	if c, ok := code[msg]; ok {
		return c
	}
	return ServerCommonError
}

func IsCodeErr(errCode int64) bool {
	_, ok := message[errCode]
	return ok
}

func NewError(errCode int64) error {
	if v, ok := message[errCode]; ok {
		return errors.New(v)
	}
	return errors.New(message[ServerCommonError])
}
