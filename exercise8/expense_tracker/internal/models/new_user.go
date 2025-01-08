package models

type NewUser struct {
	UserID int `json:"user_id"`
	Amount int `json:"budget"`
}
type NewUserResponse struct {
	UserID   int `json:"user_id"`
	BudgetID int `json:"budget_id"`
	Amount   int `json:"budget"`
}
