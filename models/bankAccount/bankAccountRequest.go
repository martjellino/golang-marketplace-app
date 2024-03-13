package bankaccount

type BankAccountRequest struct {
	BankName          string `json:"bankName" binding:"required,min=5,max=15" validate:"required,min=5,max=15"`
	BankAccountName   string `json:"bankAccountName" binding:"required,min=5,max=15" validate:"required,min=5,max=15"`
	BankAccountNumber string `json:"bankAccountNumber" binding:"required,min=5,max=15" validate:"required,min=5,max=15"`
}
