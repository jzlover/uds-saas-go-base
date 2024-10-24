package common

import (
	"time"
)

type Res[T any] struct {
	Data T      `json:"data"`
	Msg  string `json:"msg"`
	Code int16  `json:"code"`
}

type CommonFields struct {
	Id        uint64    `db:"id"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
	Status    int16     `db:"status"`
}

const (
	Password      string = "0"
	Tel           string = "1"
	Email         string = "2"
	ValicodeEmail string = "3"
	Wechat        string = "4"
)

const (
	Role_User    string = "User"
	Role_Company string = "Company"
	Role_Admin   string = "Admin"
)
