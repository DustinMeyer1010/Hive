package db

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/DustinMeyer1010/hive/internal/types"
)

// Provided a username and the feilds it will return the account that was found with the username.
// Returns the account with those values filled will the fields provided all other fileds will be blank
func GetAccountByUsername(username string, fields ...string) (*types.Account, error) {
	var acc types.Account

	if !validateFields(fields, &acc) {
		return nil, fmt.Errorf("invalid fields for accounts")
	}
	queryTemplate, err := os.ReadFile("internal/db/queries/account/SELECT_ACCOUNT_BY_USERNAME.SQL")

	if err != nil {
		return nil, fmt.Errorf("sql query not found: %w", err)
	}

	query := strings.Replace(string(queryTemplate), "{{fields}}", strings.Join(fields, ", "), 1)

	row := pool.QueryRow(context.Background(), string(query), username)

	scanArgs, err := acc.BuildScanArgs(fields)

	if err != nil {
		return nil, err
	}

	row.Scan(scanArgs...)

	if err != nil {
		return nil, fmt.Errorf("account not found: %v", err)
	}

	return &acc, nil
}
