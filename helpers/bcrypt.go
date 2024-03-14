package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	cost := 12
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, cost)

	return string(hash)
}

func ComparePassword(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
