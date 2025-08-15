package types

import "golang.org/x/crypto/bcrypt"

type Account struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
}

func (a *Account) HashPassword() (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.HashedPassword), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	a.HashedPassword = string(hashedPassword)

	return nil

}
