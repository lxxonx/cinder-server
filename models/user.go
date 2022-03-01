package models

type User struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Uni      string `json:"uni"`
	Dep      string `json:"dep"`
	Bio      string `json:"bio"`
}
