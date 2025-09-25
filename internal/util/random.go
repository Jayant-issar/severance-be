package util

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	usernameLength = 10
	passwordLength = 12
)

var letters = "abcdefghijklmnopqrstuvwxyz"
var digits = "0123456789"
var symbols = "!@#$%^&*"

// RandomUUID generates a random UUID string
func RandomUUID() string {
	return uuid.New().String()
}

// RandomUsername generates a random username consisting of lowercase letters
func RandomUsername() string {
	return randomString(usernameLength, letters)
}

// RandomEmail generates a random email address
func RandomEmail() string {
	username := RandomUsername()
	return fmt.Sprintf("%s@example.com", username)
}

// RandomPassword generates a random password with letters, digits, and symbols
func RandomPassword() string {
	charset := letters + strings.ToUpper(letters) + digits + symbols
	return randomString(passwordLength, charset)
}

// randomString generates a random string of given length from the charset
func randomString(length int, charset string) string {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))
	for i := range result {
		randomIndex, _ := rand.Int(rand.Reader, charsetLen)
		result[i] = charset[randomIndex.Int64()]
	}
	return string(result)
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns nil on success, or an error if not matching.
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
