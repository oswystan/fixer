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

var userById = "select * from users where id = ?"
var userByNicky = "select * from users where nicky = ?"
var userLikeNicky = "select * from users where nicky like '%s%%';"

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
	_, err := GetDB().pg.QueryOne(usr, userById, id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *user) GetUserByNicky(nicky string) (*model.User, error) {
	usr := &model.User{}
	_, err := GetDB().pg.QueryOne(usr, userByNicky, nicky)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *user) GetUserLikeNicky(nicky string) ([]model.User, error) {
	var usr []model.User
	sql := fmt.Sprintf(userLikeNicky, nicky)
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
