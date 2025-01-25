package models

type RequestNewUser struct {
	NewUserRequest NewUser `json:"data"`
}
type NewUser struct {
	Amount float64 `json:"budget"`
}
type NewUserResponse struct {
	UserID   int     `json:"user_id"`
	BudgetID int     `json:"budget_id"`
	Amount   float64 `json:"budget"`
}
