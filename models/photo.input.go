package models

type PhotoInput struct {
	Id    string  `json:"id"`
	Photo []uint8 `json:"photo"`
}
