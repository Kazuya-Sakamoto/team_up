package errorcode

import "errors"

// C0001 詳細踏不明なエラー
const C0001 = ":::C0001:::詳細不明"

// GetUnknownError ...
func GetUnknownError() error {
	return errors.New(C0001)
}

// C0002 楽観ロック
const C0002 = ":::C0002:::楽観ロック"

// GetOptimalLockError ...
func GetOptimalLockError() error {
	return errors.New(C0002)
}

// C0003 更新タイプエラー
const C0003 = ":::C0003:::適切なUpdateTypeではない"

// GetInvalidUpdateTypeError ...
func GetInvalidUpdateTypeError() error {
	return errors.New(C0002)
}

// GetErrorCodeMapWithCommon ...
func GetErrorCodeMapWithCommon() map[string]ErrorInfo {

	return map[string]ErrorInfo{
		C0001: {
			Code:    500,
			Message: "予期せぬエラーが発生いたしました。リロードを行い再度実行するか、システム管理者にご連絡ください。",
		},
		C0002: {
			Code:    423,
			Message: "既にデータが更新されています。リロードして再度入力してください。",
		},
		C0003: {
			Code:    422,
			Message: "更新方法が適切ではありません。システム管理者にご連絡ください。",
		},
	}
}
