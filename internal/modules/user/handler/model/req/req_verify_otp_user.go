package req

type UserReqVerifyOTPRequest struct {
	VerifyKey  string `json:"verify_key"`
	VerifyCode string `json:"verify_code"`
}
