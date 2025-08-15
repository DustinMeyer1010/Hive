package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DustinMeyer1010/livechat/internal/db"
	"github.com/DustinMeyer1010/livechat/internal/types"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account types.Account

	defer r.Body.Close()

	if VerifyPostRequest(r.Method) {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	err := db.CreateAccount(account)

	fmt.Println(err)

	if err != nil {
		http.Error(w, "account not added", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var account *types.Account

	defer r.Body.Close()

	if VerifyPostRequest(r.Method) {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
	}

	account, err := db.GetAccountByUsername("dustinmeyer")

	fmt.Println(err)

	fmt.Println(account.Username, account.HashedPassword)
}
