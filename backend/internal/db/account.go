package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DustinMeyer1010/livechat/internal/types"
)

func CreateAccount(account types.Account) error {
	query, err := os.ReadFile("internal/db/queries/account/INSERT_ACCOUNT.SQL")

	if err != nil {
		return fmt.Errorf("sql query not found: %w", err)
	}

	if err = account.HashPassword(); err != nil {
		return fmt.Errorf("failed to hash password %w", err)
	}

	_, err = Pool.Exec(context.Background(), string(query), account.Username, account.HashedPassword, account.Email, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func GetAccountByUsername(username string) (*types.Account, error) {
	var account types.Account
	query, err := os.ReadFile("internal/db/queries/account/SELECT_ACCOUNT_BY_USERNAME.SQL")

	if err != nil {
		return nil, fmt.Errorf("sql query not found: %w", err)
	}

	row := Pool.QueryRow(context.Background(), string(query), username)

	err = row.Scan(&account.Username, &account.HashedPassword)

	fmt.Println(err)

	return &account, nil
}
