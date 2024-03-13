package services

import (
	"database/sql"
	"fmt"
	bankaccount "golang-marketplace-app/models/bankAccount"
	"log"
	"strconv"
	"time"
)

func CreateBankAccount(Request bankaccount.BankAccountRequest, db *sql.DB) (bankaccount.BankAccountResponse, error) {
	const dummyUserId = 1;
	
	stmt, err := db.Prepare("INSERT INTO bank_accounts (user_id, bank_name, account_name, account_number) VALUES ($1, $2, $3, $4)")
	if err != nil {
			log.Println("Error preparing SQL query:", err)
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(dummyUserId, Request.BankName, Request.BankAccountName, Request.BankAccountNumber)
	if err != nil {
			log.Println("Error executing insert statement:", err)
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error executing insert statement: %v", err)
	}

	var accountID int
	err = db.QueryRow("SELECT LASTVAL()").Scan(&accountID)
	if err != nil {
			log.Println("Error retrieving last inserted ID:", err)
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error retrieving last inserted ID: %v", err)
	}

	parsedAccountId := strconv.Itoa(accountID)
	parsedUserId := strconv.Itoa(dummyUserId)

	return bankaccount.BankAccountResponse {
		AccountID: parsedAccountId,
		BankName: Request.BankName,
		AccountName: Request.BankAccountName,
		AccountNumber: Request.BankAccountNumber,
		UserID: parsedUserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil;
}

func UpdateBankAccountByAccountId(accountId int, Request bankaccount.BankAccountRequest, db *sql.DB) (bankaccount.BankAccountResponse, error) {
	stmt, err := db.Prepare("UPDATE bank_accounts SET bank_name=$1, account_name=$2, account_number=$3, updated_at=$4 WHERE account_id=$5")
	if err != nil {
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(Request.BankName, Request.BankAccountName, Request.BankAccountNumber, time.Now(), accountId)
	if err != nil {
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error executing update statement: %v", err)
	}
	
	return bankaccount.BankAccountResponse{
			AccountID:     strconv.Itoa(accountId),
			BankName:      Request.BankName,
			AccountName:   Request.BankAccountName,
			AccountNumber: Request.BankAccountNumber,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
	}, nil
}

func FindBankAccountByAccountId(accountId int, db *sql.DB) (bankaccount.BankAccountResponse, error) {
	var (
			parsedAccountId int
			bankName        string
			accountName     string
			accountNumber   string
			userId          int
			createdAt       time.Time
			updatedAt       time.Time
	)

	query := fmt.Sprintf("SELECT account_id, user_id, bank_name, account_name, account_number, created_at, updated_at FROM bank_accounts WHERE account_id = %d", accountId)
	fmt.Println("Query:", query)

	err := db.QueryRow(query).Scan(&parsedAccountId, &userId, &bankName, &accountName, &accountNumber, &createdAt, &updatedAt)
	if err != nil {
			log.Println(err)
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error retrieving bank account details: %v", err)
	}

	return bankaccount.BankAccountResponse{
			AccountID:     strconv.Itoa(parsedAccountId),
			BankName:      bankName,
			AccountName:   accountName,
			AccountNumber: accountNumber,
			UserID:        strconv.Itoa(userId),
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
	}, nil
}