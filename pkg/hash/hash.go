package hash

import (
	customErrors "backend/pkg/errors"
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/mail"
	"unicode"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckPassword(pass string) error {
	if len(pass) <= 6 {
		return customErrors.ErrShortPassword
	}

	for i := 0; i < len(pass); i++ {
		if pass[i] > unicode.MaxASCII {
			return customErrors.ErrLatinPassword
		}
	}
	return nil
}

func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
