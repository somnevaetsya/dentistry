package token

import "github.com/google/uuid"

func GenerateSessionToken() string {
	return uuid.NewString()
}
