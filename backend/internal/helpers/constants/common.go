package constants

import "os"

const PAGE_LIMIT_DEFAULT = 10

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "BOIDSGBOABSDGIASBDIOBASDBOBIDSBOFUAB"
	}
	return secret
}
