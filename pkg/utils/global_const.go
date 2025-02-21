package utils

import "errors"

type TokenTypeEnum string

// Token Type
const (
	VerificationToken         TokenTypeEnum = "VERIFICATION_TOKEN"
	VerifyForgotPasswordToken TokenTypeEnum = "VERIFICATION_FORGOT_PASSWORD_TOKEN"
	AccessToken               TokenTypeEnum = "ACCESS_TOKEN"
)

// Global Errors
var (
	ErrorGlobalPublicMessage = errors.New("ops, something went wrong")
)
