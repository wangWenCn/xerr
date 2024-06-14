package xerr

// OK 成功返回
const OK uint32 = 0

// **(前3位代表业务,后三位代表具体功能)**/
// 全局错误码

const ServerCommonError uint32 = 100000
const RequestParamError uint32 = 100002
const TokenExpireError uint32 = 100003
const TokenGenerateError uint32 = 100004
const DBError uint32 = 100005
const DBUpdateAffectedZeroError uint32 = 100006

// NOT FOUND

const ErrorReasonDataNotFound uint32 = 101

// 认证失败

const ErrorReasonAuthenticationFail uint32 = 102

// 请求参数缺失

const ErrorReasonRequestParamMissing uint32 = 103

// 权限不足

const ErrorReasonRequestForbidden uint32 = 104

// 重复do

const ErrorReasonRepeatedDo uint32 = 105

// 非法操作

const ErrorReasonIllegalOperation uint32 = 106

// 超出限制范围

const ErrorReasonBeyondLimitRange uint32 = 107

// appId无效

const ErrorAppId uint32 = 1001

// SecretKey错误

const ErrorSecretKey uint32 = 1002

// 签名错误

const ErrorAuth uint32 = 1003
