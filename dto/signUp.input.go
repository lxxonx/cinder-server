package dto

type SignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
