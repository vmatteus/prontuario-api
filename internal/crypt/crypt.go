package crypt

import "golang.org/x/crypto/bcrypt"

type Crypt struct{}

func NewCrypt() Crypt {
	return Crypt{}
}

func (Crypt) HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (Crypt) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
