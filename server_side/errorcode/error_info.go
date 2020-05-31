package errorcode

// ErrorInfo ...
type ErrorInfo struct {
	Code    int    // HTTPコード
	Message string // ユーザーに表示するためのメッセージ
}

// GetUnknownErrorInfo 主に原因が不明であるときに使用
func GetUnknownErrorInfo() ErrorInfo {
	ecm := GetErrorCodeMapWithCommon()
	return ecm[C0001]
}

// GetErrorInfoFromErrorCode ...
func GetErrorInfoFromErrorCode(errorCode string) ErrorInfo {

	commonErrorCodeMap := GetErrorCodeMapWithCommon()
	if v, ok := commonErrorCodeMap[errorCode]; ok {
		return v
	}

	return GetUnknownErrorInfo()
}
