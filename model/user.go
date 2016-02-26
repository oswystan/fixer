//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: user.go
//     description:
//         created: 2016-02-25 22:42:02
//          author: wystan
//
//===============================================================================
package model

import "time"

type User struct {
	Id            int32     `json: "id"`
	Nicky         string    `json: "nicky"`
	Email         string    `json: "email"`
	Pwd           string    `json: "-"`
	Portrait      string    `json: "portrait"`
	RegisterDate  time.Time `json: "register_date"`
	LastLoginTime time.Time `json: "last_login_time"`
}

//==================================== END ======================================
