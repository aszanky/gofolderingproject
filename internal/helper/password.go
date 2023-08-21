package helper

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func PasswordGenerator(passlength int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, passlength)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:passlength]
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
