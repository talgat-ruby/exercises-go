package models

type EditUser struct {
	IdUser          int     `json:"id_user"`
	IdBudget        int     `json:"id_budget"`
	Expense         float64 `json:"expense"`
	Comment         string  `json:"comment"`
	ExpenseCategory string  `json:"expense_category"`
}
