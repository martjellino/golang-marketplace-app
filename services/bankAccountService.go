package services

import (
	"database/sql"
	bankaccount "golang-marketplace-app/models/bankAccount"
	"log"
	"time"
)

func CreateBankAccount(Request bankaccount.BankAccountRequest, db *sql.DB) (bankaccount.BankAccountResponse, error) {
	const dummyUserId = 1;
	
	stmt, err := db.Prepare("INSERT INTO bank_accounts (user_id, bank_name, account_name, account_number) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal("Error preparing SQL Query:", err)
		return bankaccount.BankAccountResponse{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(dummyUserId, Request.BankName, Request.BankAccountName, Request.BankAccountNumber)
	if err != nil {
			log.Fatal("Error executing insert statement:", err)
			return bankaccount.BankAccountResponse{}, err
	}

	var accountID int
	err = db.QueryRow("SELECT LASTVAL()").Scan(&accountID)
	if err != nil {
			log.Fatal("Error retrieving last inserted ID:", err)
			return bankaccount.BankAccountResponse{}, err
	}

	return bankaccount.BankAccountResponse {
		AccountID: accountID,
		BankName: Request.BankName,
		AccountName: Request.BankAccountName,
		AccountNumber: Request.BankAccountNumber,
		UserID: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil;
}