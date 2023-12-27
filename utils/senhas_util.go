package utils

import "golang.org/x/crypto/bcrypt"

// Para criar um hash de senha
func GerarHashSenha(senha string) (string, error) {
	hashedSenha, err := bcrypt.GenerateFromPassword([]byte(senha), 12)
	return string(hashedSenha), err
}
