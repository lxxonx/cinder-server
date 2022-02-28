package models

import (
	"time"
)

type Group struct {
	Uid        string        `json:"uid"`
	GroupName  string        `json:"groupName"`
	Pics       []interface{} `json:"pics"`
	CreateTime time.Time     `json:"createTime"`
	Bio        string        `json:"bio"`
	Members    []User        `json:"members"`
}
