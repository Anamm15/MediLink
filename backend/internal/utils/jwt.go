package utils

import (
	"errors"
	"fmt"
	"time"

	"MediLink/internal/helpers/constants"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Definisi durasi token (bisa dipindah ke config/env)
const tokenDuration = 24 * time.Hour

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// 1. Custom Claims Struct
// Menggunakan struct daripada MapClaims agar Type-Safe dan mencegah typo saat akses key.
type JWTCustomClaims struct {
	UserID uuid.UUID          `json:"user_id"`
	Role   constants.UserRole `json:"role"`
	jwt.RegisteredClaims
}

// 2. Generate Token
func GenerateJWT(userID uuid.UUID, role constants.UserRole) (string, error) {
	secretKey := constants.GetJWTSecret()
	// Setup claims
	claims := JWTCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			// Best Practice: Selalu set Expiration (exp)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
			// Best Practice: Set IssuedAt (iat) untuk audit kapan token dibuat
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// Opsional: Set Issuer (iss) untuk identifikasi aplikasi
			Issuer: "medilink-app",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	return token.SignedString([]byte(secretKey))
}

// 3. Validate & Parse Token
// Mengembalikan *JWTCustomClaims agar data user_id dan role bisa langsung diakses
func ValidateJWT(tokenString string) (*JWTCustomClaims, error) {
	secretKey := constants.GetJWTSecret()
	// ParseWithClaims memvalidasi signature dan exp secara otomatis
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Best Practice: Verifikasi algoritma signing untuk mencegah serangan "None" alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	// Error handling detail
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, err
	}

	// Validasi Claims dan Token
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
