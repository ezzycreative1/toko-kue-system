package mid

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

// Helper function to check if the slice contains a string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateRandString ..
func GenerateRandString(lg int) string {
	var letterRunes = []rune("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, lg)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Generate code
func GenerateCode(code string, year, counter int) string {
	counterStr := fmt.Sprintf("%04d", counter)
	return fmt.Sprintf("%s-%d%s%04s", code, year, "000", counterStr)
}

// Generate Product code
func GenerateProductCode(code string, year, counter int) string {
	counterStr := fmt.Sprintf("%04d", counter)
	return fmt.Sprintf("%s-%d%04s", code, year, counterStr)
}
