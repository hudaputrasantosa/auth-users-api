package otp

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateOTP(length ...int) string {
	lengthDefault := 6
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	var otp strings.Builder

	if len(length) > 0 {
		lengthDefault = length[0]
	}

	for i := 0; i < lengthDefault; i++ {
		otp.WriteByte(charset[r.Intn(len(charset))])
	}

	return otp.String()
}
