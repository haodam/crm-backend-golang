package req

type UserRegistrationRequest struct {
	VerifyKey     string `json:"verify_key" binding:"required,email"`                            // Trường này là email
	VerifyType    int    `json:"verify_type" binding:"required,oneof=1 2"`                       // 1: đăng ký bằng email, 2: đăng ký bằng số điện thoại
	VerifyPurpose string `json:"verify_purpose" binding:"required,oneof=TEST_USER PRODUCT_USER"` // TEST_USER cho phát triển, PRODUCT_USER cho sản phẩm
}
