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

type StoreUsers interface {
}

type storeusers struct {
}

func NewStoreUsers() StoreUsers {
	return &storeusers{}
}

//==================================== END ======================================
