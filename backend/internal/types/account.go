package types

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	CreationDate   time.Time `json:"creation_date"`
}

func (a *Account) HashPassword() (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.HashedPassword), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	a.HashedPassword = string(hashedPassword)

	return nil
}

// Valid fields name for accounts, (will have all the columns in account table)
// Builds the scan aguments to fill in account types with attributes pulled from database
func (a *Account) BuildScanArgs(fields []string) ([]any, error) {
	var scanArgs []any
	for _, field := range fields {
		switch field {
		case "id":
			scanArgs = append(scanArgs, &a.ID)
		case "username":
			scanArgs = append(scanArgs, &a.Username)
		case "hashed_password":
			scanArgs = append(scanArgs, &a.HashedPassword)
		case "email":
			scanArgs = append(scanArgs, &a.Email)
		case "created_at":
			scanArgs = append(scanArgs, &a.CreationDate)
		default:
			return nil, fmt.Errorf("field not found")
		}

	}

	return scanArgs, nil
}

func (a Account) Fields() []string {
	return []string{
		"id",
		"username",
		"hashed_password",
		"email",
		"creation_date",
	}
}
