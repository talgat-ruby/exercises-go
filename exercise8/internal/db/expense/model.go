package expense

type ModelExpense struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CreateModelExpense struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
