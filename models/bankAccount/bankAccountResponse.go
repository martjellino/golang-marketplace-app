package bankaccount

import "time"

type BankAccountResponse struct {
	AccountID     string    `json:"bankAccountId"`
	UserID        string    `json:"user_id,omitempty"`
	BankName      string    `json:"bankName"`
	AccountName   string    `json:"bankAccountName"`
	AccountNumber string    `json:"bankAccountNumber"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
