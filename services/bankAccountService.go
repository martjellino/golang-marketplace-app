package services

import (
	"fmt"
	"golang-marketplace-app/database"
	bankaccount "golang-marketplace-app/models/bankAccount"
	"log"
	"strconv"
	"time"
)

func CreateBankAccount(userId int, Request bankaccount.BankAccountRequest) (bankaccount.BankAccountResponse, error) {
	var accountID int
	err := database.DB.QueryRow("SELECT nextval('bank_accounts_account_id_seq')").Scan(&accountID)
	if err != nil {
			log.Println("Error retrieving next sequence value:", err)
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error retrieving next sequence value: %v", err)
	}
	
	stmt, err := database.DB.Prepare("INSERT INTO bank_accounts (user_id, bank_name, account_name, account_number) VALUES ($1, $2, $3, $4)")
	if err != nil {
			log.Println("Error preparing SQL query:", err)
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, Request.BankName, Request.BankAccountName, Request.BankAccountNumber)
	if err != nil {
			log.Println("Error executing insert statement:", err)
			return bankaccount.BankAccountResponse{}, fmt.Errorf("error executing insert statement: %v", err)
	}

	parsedAccountId := strconv.Itoa(accountID)
	parsedUserId := strconv.Itoa(userId)

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

func UpdateBankAccountByAccountId(accountId int, Request bankaccount.BankAccountRequest) (bankaccount.BankAccountResponse, error) {
	stmt, err := database.DB.Prepare("UPDATE bank_accounts SET bank_name=$1, account_name=$2, account_number=$3, updated_at=$4 WHERE account_id=$5")
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

func FindBankAccountByAccountId(accountId int) (bankaccount.BankAccountResponse, error) {
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

	err := database.DB.QueryRow(query).Scan(&parsedAccountId, &userId, &bankName, &accountName, &accountNumber, &createdAt, &updatedAt)
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

func DeleteBankAccountByAccountId(accountId int) error {
	query := "DELETE FROM bank_accounts WHERE account_id = $1"

	_, err := database.DB.Exec(query, accountId)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error deleting bank account: %v", err)
	}

	return nil
}

func GetBankAccountsByUserId(userId int) ([]bankaccount.BankAccountResponse, error) {
    query := `
        SELECT account_id, bank_name, account_name, account_number, created_at, updated_at 
        FROM bank_accounts 
        WHERE user_id = $1`
    stmt, err := database.DB.Prepare(query)
    if err != nil {
        log.Println(err)
        return nil, fmt.Errorf("error preparing SQL query: %v", err)
    }
    defer stmt.Close()
    
    rows, err := stmt.Query(userId)
    if err != nil {
        log.Println("Error executing SQL query:", err)
        return nil, fmt.Errorf("error executing SQL query: %v", err)
    }
    defer rows.Close()
    
    var accounts []bankaccount.BankAccountResponse
    for rows.Next() {
        var account bankaccount.BankAccountResponse
        if err := rows.Scan(&account.AccountID, &account.BankName, &account.AccountName, &account.AccountNumber, &account.CreatedAt, &account.UpdatedAt); err != nil {
            log.Println(err)
            return nil, fmt.Errorf("error scanning row: %v", err)
        }
        accounts = append(accounts, account)
    }
    if err := rows.Err(); err != nil {
        log.Println(err)
        return nil, fmt.Errorf("error iterating over rows: %v", err)
    }
    
    return accounts, nil
}