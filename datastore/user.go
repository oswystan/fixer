//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: user.go
//     description:
//         created: 2016-02-25 22:46:52
//          author: wystan
//
//===============================================================================

package datastore

import (
	"fmt"
	"log"

	"github.com/oswystan/fixer/model"
)

var sqlById = "select * from users where id = ?"
var sqlByNicky = "select * from users where nicky = ?"
var sqlLikeNicky = "select * from users where nicky like '%s%%';"

type StoreUser interface {
	GetUserById(id int) (*model.User, error)
	GetUserByNicky(nicky string) (*model.User, error)
	GetUserLikeNicky(nicky string) ([]model.User, error)

	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
}

type user struct {
}

func (u *user) GetUserById(id int) (*model.User, error) {
	usr := &model.User{}
	_, err := GetDB().pg.QueryOne(usr, sqlById, id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *user) GetUserByNicky(nicky string) (*model.User, error) {
	usr := &model.User{}
	_, err := GetDB().pg.QueryOne(usr, sqlByNicky, nicky)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *user) GetUserLikeNicky(nicky string) ([]model.User, error) {
	var usr []model.User
	sql := fmt.Sprintf(sqlLikeNicky, nicky)
	log.Println(sql)
	_, err := GetDB().pg.Query(&usr, sql)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *user) Create(usr *model.User) error {
	return nil
}
func (u *user) Update(usr *model.User) error {
	return nil
}
func (u *user) Delete(usr *model.User) error {
	return nil
}

func NewStoreUser() StoreUser {
	return &user{}
}

//==================================== END ======================================
