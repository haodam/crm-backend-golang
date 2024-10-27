package random

import "testing"

func TestGenerateSixDigOtp(t *testing.T) {
	for i := 0; i < 1000; i++ {
		otp := GenerateSixDigOtp()

		if otp < 100000 || otp > 999999 {
			t.Errorf("OTP %d không nằm trong khoảng 100000-999999", otp)
		}
	}
}
