package utils

import "errors"

var (
	ErrorUserRegister     = errors.New("error, can't process register")
	FailedUserRegister    = errors.New("please use different email or username")
	FailedUserLogin       = errors.New("email or password is wrong")
	ErrorUserVerification = errors.New("error, can't process verification")
)
