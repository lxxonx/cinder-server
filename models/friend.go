package models

type Friend struct {
	Username   string   `json:"username"`
	ActualName string   `json:"actual_name"`
	Uni        string   `json:"uni"`
	Dep        string   `json:"dep"`
	Bio        string   `json:"bio"`
	Gender     string   `json:"gender"`
	BirthYear  int      `json:"birth_year"`
	Avatar     string   `json:"avatar"`
	Pics       []string `json:"pics"`
}
