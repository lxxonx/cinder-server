package models

type User struct {
	Username   string   `json:"username"`
	ActualName string   `json:"actual_name"`
	Uni        string   `json:"uni"`
	Dep        string   `json:"dep"`
	Bio        string   `json:"bio"`
	Gender     string   `json:"gender"`
	BirthYear  int      `json:"birth_year"`
	IsVerified bool     `json:"is_verified"`
	Status     bool     `json:"status"`
	Pics       []string `json:"pics"`
}
