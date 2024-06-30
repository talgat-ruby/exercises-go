package todo

type CreateModels struct {
	Description string `json:"description"`
}

type UpdateModels struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
