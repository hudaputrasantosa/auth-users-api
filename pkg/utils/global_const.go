package utils

import "errors"

type TokenTypeEnum string

// Token Type
const (
	VerificationToken TokenTypeEnum = "VERIFICATION_TOKEN"
	AccessToken       TokenTypeEnum = "ACCESS_TOKEN"
)

// Global Errors
var (
	ErrorGlobalPublicMessage = errors.New("ops, something went wrong")
)
