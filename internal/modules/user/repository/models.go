// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type PreGoAccUserTwoFactor9999TwoFactorAuthType string

const (
	PreGoAccUserTwoFactor9999TwoFactorAuthTypeSMS   PreGoAccUserTwoFactor9999TwoFactorAuthType = "SMS"
	PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL PreGoAccUserTwoFactor9999TwoFactorAuthType = "EMAIL"
	PreGoAccUserTwoFactor9999TwoFactorAuthTypeAPP   PreGoAccUserTwoFactor9999TwoFactorAuthType = "APP"
)

func (e *PreGoAccUserTwoFactor9999TwoFactorAuthType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PreGoAccUserTwoFactor9999TwoFactorAuthType(s)
	case string:
		*e = PreGoAccUserTwoFactor9999TwoFactorAuthType(s)
	default:
		return fmt.Errorf("unsupported scan type for PreGoAccUserTwoFactor9999TwoFactorAuthType: %T", src)
	}
	return nil
}

type NullPreGoAccUserTwoFactor9999TwoFactorAuthType struct {
	PreGoAccUserTwoFactor9999TwoFactorAuthType PreGoAccUserTwoFactor9999TwoFactorAuthType `json:"pre_go_acc_user_two_factor_9999_two_factor_auth_type"`
	Valid                                      bool                                       `json:"valid"` // Valid is true if PreGoAccUserTwoFactor9999TwoFactorAuthType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPreGoAccUserTwoFactor9999TwoFactorAuthType) Scan(value interface{}) error {
	if value == nil {
		ns.PreGoAccUserTwoFactor9999TwoFactorAuthType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PreGoAccUserTwoFactor9999TwoFactorAuthType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPreGoAccUserTwoFactor9999TwoFactorAuthType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PreGoAccUserTwoFactor9999TwoFactorAuthType), nil
}

// pre_go_acc_user_base_9999
type PreGoAccUserBase9999 struct {
	UserID         int32          `json:"user_id"`
	UserAccount    string         `json:"user_account"`
	UserPassword   string         `json:"user_password"`
	UserSalt       string         `json:"user_salt"`
	UserLoginTime  sql.NullTime   `json:"user_login_time"`
	UserLogoutTime sql.NullTime   `json:"user_logout_time"`
	UserLoginIp    sql.NullString `json:"user_login_ip"`
	UserCreatedAt  sql.NullTime   `json:"user_created_at"`
	UserUpdatedAt  sql.NullTime   `json:"user_updated_at"`
	// authentication is enabled for the user base
	IsTwoFactorEnabled sql.NullInt32 `json:"is_two_factor_enabled"`
}

// pre_go_acc_user_info_9999
type PreGoAccUserInfo9999 struct {
	// User ID
	UserID uint64 `json:"user_id"`
	// User account
	UserAccount string `json:"user_account"`
	// User nickname
	UserNickname sql.NullString `json:"user_nickname"`
	// User avatar
	UserAvatar sql.NullString `json:"user_avatar"`
	// User state: 0-Locked, 1-Activated, 2-Not Activated
	UserState uint8 `json:"user_state"`
	// Mobile phone number
	UserMobile sql.NullString `json:"user_mobile"`
	// User gender: 0-Secret, 1-Male, 2-Female
	UserGender sql.NullInt16 `json:"user_gender"`
	// User birthday
	UserBirthday sql.NullTime `json:"user_birthday"`
	// User email address
	UserEmail sql.NullString `json:"user_email"`
	// Authentication status: 0-Not Authenticated, 1-Pending, 2-Authenticated, 3-Failed
	UserIsAuthentication uint8 `json:"user_is_authentication"`
	// Record creation time
	CreatedAt sql.NullTime `json:"created_at"`
	// Record update time
	UpdatedAt sql.NullTime `json:"updated_at"`
}

// pre_go_acc_user_two_factor_9999
type PreGoAccUserTwoFactor9999 struct {
	TwoFactorID         uint32                                     `json:"two_factor_id"`
	UserID              uint32                                     `json:"user_id"`
	TwoFactorAuthType   PreGoAccUserTwoFactor9999TwoFactorAuthType `json:"two_factor_auth_type"`
	TwoFactorAuthSecret string                                     `json:"two_factor_auth_secret"`
	TwoFactorPhone      sql.NullString                             `json:"two_factor_phone"`
	TwoFactorEmail      sql.NullString                             `json:"two_factor_email"`
	TwoFactorIsActive   bool                                       `json:"two_factor_is_active"`
	TwoFactorCreatedAt  sql.NullTime                               `json:"two_factor_created_at"`
	TwoFactorUpdatedAt  sql.NullTime                               `json:"two_factor_updated_at"`
}

// account_user_verify
type PreGoAccUserVerify9999 struct {
	VerifyID        int32         `json:"verify_id"`
	VerifyOtp       string        `json:"verify_otp"`
	VerifyKey       string        `json:"verify_key"`
	VerifyKeyHash   string        `json:"verify_key_hash"`
	VerifyType      sql.NullInt32 `json:"verify_type"`
	IsVerified      sql.NullInt32 `json:"is_verified"`
	IsDeleted       sql.NullInt32 `json:"is_deleted"`
	VerifyCreatedAt sql.NullTime  `json:"verify_created_at"`
	VerifyUpdatedAt sql.NullTime  `json:"verify_updated_at"`
}
