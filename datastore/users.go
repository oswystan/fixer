//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: users.go
//     description:
//         created: 2016-02-29 18:00:40
//          author: wystan
//
//===============================================================================

package datastore

import (
	"fmt"

	"github.com/oswystan/fixer/model"
)

var sqlUsers = `select * from users where nicky like '%s%%' offset ? limit ?`
var sqlUser = `select * from users where id = ?`
var sqlDelUser = `delete from users where id = ?`
var sqlDelUsers = `delete from users`
var sqlCreate = `insert into users(nicky, email, pwd, portrait, register_date) 
				values (?nicky, ?email, ?pwd, ?portrait, ?register_date) returning *`
var sqlUpdate = `update users set nicky=?nicky, email=?email, pwd=?pwd, portrait=?portrait where id=?id returning *`

type UserStore interface {
	GetUsers(f *model.Filter) ([]model.User, error)
	GetUser(uid int) (*model.User, error)

	Create(u *model.User) (*model.User, error)
	Update(u *model.User) (*model.User, error)
	Delete(id int) error
	DeleteUsers() error
}

type usertore struct {
}

func (us *usertore) GetUsers(f *model.Filter) ([]model.User, error) {
	var ul []model.User
	db := GetDB()

	sql := fmt.Sprintf(sqlUsers, f.Q)
	_, err := db.pg.Query(&ul, sql, f.Offset, f.Limit)
	if err != nil {
		return nil, err
	}

	return ul, nil
}

func (us *usertore) GetUser(uid int) (*model.User, error) {
	user := &model.User{}
	db := GetDB()
	_, err := db.pg.QueryOne(user, sqlUser, uid)
	return user, err
}

func (us *usertore) Create(u *model.User) (*model.User, error) {
	db := GetDB()
	_, err := db.pg.QueryOne(u, sqlCreate, u)
	return u, err
}

func (us *usertore) Update(u *model.User) (*model.User, error) {
	db := GetDB()
	_, err := db.pg.QueryOne(u, sqlUpdate, u)
	return u, err
}

func (us *usertore) Delete(id int) error {
	db := GetDB()
	_, err := db.pg.Exec(sqlDelUser, id)
	return err
}

func (us *usertore) DeleteUsers() error {
	db := GetDB()
	_, err := db.pg.Exec(sqlDelUsers)
	return err
}

func NewUserStore() UserStore {
	return &usertore{}
}

//==================================== END ======================================
