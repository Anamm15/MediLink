package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"MediLink/internal/helpers/constants"
)

func BuildOTPEmailBody(name, otp string) string {
	return fmt.Sprintf(constants.EmailOTPTemplate, name, otp)
}

func GenerateOTP(length int) (string, error) {
	max := big.NewInt(1)
	for i := 0; i < length; i++ {
		max.Mul(max, big.NewInt(10))
	}

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%0*d", length, n), nil
}
