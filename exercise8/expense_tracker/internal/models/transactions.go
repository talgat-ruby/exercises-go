package models

import "time"

type TransactionsReponse struct {
	IdUser       int           `json:"id_user"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Expense         float64   `json:"expense"`
	Comment         string    `json:"comment"`
	ExpenseCategory string    `json:"expense_category"`
	CreatedAT       time.Time `json:"created"`
}
