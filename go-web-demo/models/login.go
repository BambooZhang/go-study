package models

type Login struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type User struct {
	Uid   int64  `json:"uid"`
	Mid   int64  `json:"mid"`
	Phone string `json:"phone"`
}
