package xerr

// OK 成功返回
const OK int64 = 0

// **(前3位代表业务,后三位代表具体功能)**/
// 全局错误码

const SystemError int64 = -1
const ServerCommonError int64 = 100000
const RequestParamError int64 = 100002
const DBError int64 = 100005
const DBUpdateAffectedZeroError int64 = 100006

// NOT FOUND

const ErrorReasonDataNotFound int64 = 101

// 请求参数缺失

const ErrorReasonRequestParamMissing int64 = 103

// 重复操作

const ErrorReasonRepeatedDo int64 = 105

// 非法操作

const ErrorReasonIllegalOperation int64 = 106

// 超出限制范围

const ErrorReasonBeyondLimitRange int64 = 107

// appId无效

const ErrorAppId int64 = 1001

// SecretKey错误

const ErrorSecretKey int64 = 1002

// 签名错误

const ErrorAuth int64 = 1003

// 重新登陆

const ErrorReload int64 = -1506
