package utils

import "golang.org/x/crypto/bcrypt"

type IHashUtils interface {
	Hash(value string) (string, error)
	Compare(rawValue string, hashedValue string) bool
}

type bcryptHashUtils struct{}

func NewBcryptHashUtils() *bcryptHashUtils {
	return &bcryptHashUtils{}
}
func (bhu *bcryptHashUtils) Hash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	return string(bytes), err
}
func (bhu *bcryptHashUtils) Compare(rawValue string, hashedValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(rawValue))
	return err == nil
}
