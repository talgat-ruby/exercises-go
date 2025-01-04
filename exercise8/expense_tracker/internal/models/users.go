package models

type UserData struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type DBModelUser struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Salt         string `json:"salt"`
	Created      string `json:"created"`
	Update       string `json:"update"`
}
