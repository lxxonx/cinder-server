package dto

type AuthInput struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type SignUpInput struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Uni      string `json:"uni"`
	Dep      string `json:"dep"`
}
