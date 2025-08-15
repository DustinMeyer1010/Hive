package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DustinMeyer1010/hive/internal/types"
)

// Given a account it will attempt to add the account to the system
func CreateAccount(account types.Account) error {
	query, err := os.ReadFile("internal/db/queries/account/INSERT_ACCOUNT.SQL")

	if err != nil {
		return fmt.Errorf("sql query not found: %w", err)
	}

	if err = account.HashPassword(); err != nil {
		return fmt.Errorf("failed to hash password %w", err)
	}

	_, err = pool.Exec(context.Background(), string(query), account.Username, account.HashedPassword, account.Email, time.Now())

	if err != nil {
		return err
	}

	return nil
}
