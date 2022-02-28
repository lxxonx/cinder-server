package dto

type AuthInput struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type SignUpInput struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Uni      string `json:"uni"`
	Dep      string `json:"dep"`
}
