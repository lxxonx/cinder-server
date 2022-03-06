package dto

type AuthInput struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type SignUpInput struct {
	Uid        string `json:"uid"`
	Username   string `json:"username"`
	ActualName string `json:"actual_name"`
	Uni        string `json:"uni"`
	Dep        string `json:"dep"`
	BirthYear  int    `json:"birth_year"`
	Gender     string `json:"gender"`
}
