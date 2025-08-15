package db

import (
	"github.com/DustinMeyer1010/livechat/internal/types"
)

// TODO
// Validates the inputed values are valid for account table
func validateFields(fields []string, model types.Database) bool {
	if len(fields) > len(model.Fields()) {
		return false
	}

	return true

}
