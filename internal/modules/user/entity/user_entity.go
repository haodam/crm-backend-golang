package entity

type RegisterInput struct {
	VerifyKey     string `json:"verify_key"`
	VerifyType    int    `json:"verify_type"`
	VerifyPurpose string `json:"verify_purpose"`
}

type LoginInput struct {
	UserAccount  string `json:"user_account"`
	UserPassword string `json:"user_password"`
}

type LoginOutput struct {
}
