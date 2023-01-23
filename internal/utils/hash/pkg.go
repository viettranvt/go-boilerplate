package hash_util

import (
	"crypto/sha256"
	"fmt"
	configs "parishioner_management/internal/configs"

	"golang.org/x/crypto/bcrypt"
)

func SHA256(password string) string {
	hasher := sha256.New()
	if _, err := hasher.Write([]byte(password)); err != nil {
		return ""
	}
	//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	sha := fmt.Sprintf("%x", hasher.Sum(nil))
	return sha
}
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), configs.DefaultBcryptCost)
	return string(bytes), err
}

func IsValidPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
