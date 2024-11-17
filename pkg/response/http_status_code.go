package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001 // token is invalid
	ErrInvalidOTP   = 30002
	ErrSendEmailOtp = 30003
	// ErrCodeAuthFailed User Authentication
	ErrCodeAuthFailed = 40005
	// ErrCodeUserHasExists Register Code
	ErrCodeUserHasExists = 50001 // user has already registered

	// ErrCodeOtpNotExists Err Login
	ErrCodeOtpNotExists     = 60009
	ErrCodeUserOtpNotExists = 60008

	// ErrCodeTwoFactorAuthSetupFailed Two-Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed  = 80001
	ErrCodeTwoFactorAuthVerifyFailed = 80002
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "token is invalid",
	ErrInvalidOTP:       "Otp error",
	ErrSendEmailOtp:     "Failed to send email OTP",

	ErrCodeUserHasExists: "user has already registered",

	ErrCodeOtpNotExists:     "OTP exists but not registered",
	ErrCodeUserOtpNotExists: "User OTP not exists",
	ErrCodeAuthFailed:       "Authentication failed",

	// Two-Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed:  "Two Factor Authentication setup failed",
	ErrCodeTwoFactorAuthVerifyFailed: "Two Factor Authentication verify failed",
}
