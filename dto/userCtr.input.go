package dto

type AuthInput struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type SignUpInput struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Uni      string `json:"uni"`
	Dep      string `json:"dep"`
}
