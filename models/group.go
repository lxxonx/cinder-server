package models

import (
	"time"
)

type Group struct {
	Uid       string        `json:"uid"`
	Groupname string        `json:"groupname"`
	Pics      []interface{} `json:"pics"`
	CreatedAt time.Time     `json:"createdAt"`
	Bio       string        `json:"bio"`
	Members   []User        `json:"members"`
}
