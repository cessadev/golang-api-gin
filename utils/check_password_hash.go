package utils

import "golang.org/x/crypto/bcrypt"

// CheckPasswordHash compara la contraseña en texto plano con el hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
