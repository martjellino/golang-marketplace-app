package main

import (
	"database/sql"
	"errors"
	bankaccount "golang-marketplace-app/models/bankAccount"
	"golang-marketplace-app/services"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func beforeEach(t *testing.T) (sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database connection:", err)
	}

	return mock, db;
}

func TestShouldReturnBankAccountResponseAndErrorNil_whenInsertSuccess(t *testing.T) {
	var mock, db = beforeEach(t)
	mockBankName := "MockBank"
	mockAccountName := "MockAccount"
	mockAccountNumber := "1234567890"
	mockUserId := 1
	mock.ExpectPrepare("INSERT INTO bank_accounts").
		ExpectExec().
		WithArgs(mockUserId, mockBankName, mockAccountName, mockAccountNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT LASTVAL()").
		WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow(1))

	// Call the function under test
	response, err := services.CreateBankAccount(mockUserId, bankaccount.BankAccountRequest{
		BankName:          mockBankName,
		BankAccountName:   mockAccountName,
		BankAccountNumber: mockAccountNumber,
	}, db)

	// Check for errors
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, response.AccountID)
	assert.Equal(t, mockBankName, response.BankName)
	assert.Equal(t, mockAccountName, response.AccountName)
	assert.Equal(t, mockAccountNumber, response.AccountNumber)
	assert.Equal(t, 1, response.UserID)
	assert.WithinDuration(t, time.Now(), response.CreatedAt, 1*time.Second)
	assert.WithinDuration(t, time.Now(), response.UpdatedAt, 1*time.Second)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	defer db.Close()
}

func TestShouldReturnBankAccountResponseAndError_whenInsertExecutionFailed(t *testing.T) {
	var mock, db = beforeEach(t);
	mockBankName := "MockBank"
	mockAccountName := "MockAccount"
	mockAccountNumber := "1234567890"
	mockUserId := 1
	mock.ExpectPrepare("INSERT INTO bank_accounts").
		ExpectExec().
		WithArgs(mockUserId, mockBankName, mockAccountName, mockAccountNumber).
		WillReturnError(errors.New("SQL query preparation failed"))

	// Call the function under test
	response, err := services.CreateBankAccount(mockUserId, bankaccount.BankAccountRequest{
		BankName:          mockBankName,
		BankAccountName:   mockAccountName,
		BankAccountNumber: mockAccountNumber,
	}, db)

	assert.Error(t, err)
	assert.NotNil(t, response)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	defer db.Close()
}

func TestShouldReturnBankAccountResponseAndError_whenGetLastInsertedValueFailed(t *testing.T) {
	var mock, db = beforeEach(t);
	mockBankName := "MockBank"
	mockAccountName := "MockAccount"
	mockAccountNumber := "1234567890"
	mockUserId := 1
	mock.ExpectPrepare("INSERT INTO bank_accounts").
		ExpectExec().
		WithArgs(mockUserId, mockBankName, mockAccountName, mockAccountNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT LASTVAL()").
		WillReturnError(errors.New("mock QueryRow error"))

	// Call the function under test
	response, err := services.CreateBankAccount(mockUserId, bankaccount.BankAccountRequest{
		BankName:          mockBankName,
		BankAccountName:   mockAccountName,
		BankAccountNumber: mockAccountNumber,
	}, db)

	assert.Error(t, err)
	assert.NotNil(t, response)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	defer db.Close()
}