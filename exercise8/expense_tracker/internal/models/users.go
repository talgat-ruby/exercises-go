package models

type UserData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
