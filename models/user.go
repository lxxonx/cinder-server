package models

import "cloud.google.com/go/firestore"

type User struct {
	Username string        `json:"username"`
	Uni      string        `json:"uni"`
	Dep      string        `json:"dep"`
	Bio      string        `json:"bio"`
	Pics     []interface{} `json:"pics"`
}

func (g User) Create(f *firestore.DocumentSnapshot) User {
	username := f.Data()["username"]
	if username == nil {
		username = ""
	}
	uni := f.Data()["uni"]
	if uni == nil {
		uni = ""
	}
	dep := f.Data()["dep"]
	if dep == nil {
		dep = ""
	}
	bio := f.Data()["bio"]
	if bio == nil {
		bio = ""
	}
	pics := f.Data()["pics"]
	if pics == nil {
		pics = []interface{}{}
	}
	user := User{
		Username: username.(string),
		Uni:      uni.(string),
		Dep:      dep.(string),
		Bio:      bio.(string),
		Pics:     pics.([]interface{}),
	}
	return user
}
