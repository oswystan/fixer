//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: datastore.go
//     description:
//         created: 2016-02-24 19:02:59
//          author: wystan
//
//===============================================================================

package datastore

import (
	"fmt"

	"gopkg.in/pg.v3"
)

type Database struct {
	pg *pg.DB
}

var database = &Database{pg: nil}

func (db *Database) Open(user, pwd string, dbname string) error {
	if db.pg != nil {
		return fmt.Errorf("already connect to %s", dbname)
	}

	opt := &pg.Options{}
	if len(user) != 0 {
		opt.User = user
		if len(pwd) != 0 {
			opt.Password = pwd
		}
	}
	opt.Database = dbname
	db.pg = pg.Connect(opt)
	if db.pg == nil {
		return fmt.Errorf("fail to connect database")
	}

	return nil
}

func (db *Database) Close() error {
	if db.pg == nil {
		return fmt.Errorf("database already closed")
	}

	db.pg.Close()
	db.pg = nil
	return nil
}

func GetDB() *Database {
	return database
}

//==================================== END ======================================
