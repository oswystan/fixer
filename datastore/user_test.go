//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: user_test.go
//     description:
//         created: 2016-02-25 23:04:22
//          author: wystan
//
//===============================================================================

package datastore

import "testing"

func TestGetById(t *testing.T) {
	if err := openDB(); err != nil {
		t.Error("fail to connect db")
	}
	defer closeDB()

	u := NewStoreUser()
	one, err := u.GetUserById(1)
	if err != nil {
		t.Errorf("fail to get user by id [%s]", err)
		return
	}
	if one.Id != 1 {
		t.Errorf("user information is not correct %v", one)
		return
	}
	one, err = u.GetUserByNicky("john")
	if err != nil {
		t.Errorf("fail to get user by id [%s]", err)
		return
	}
	if one.Nicky != "john" {
		t.Errorf("fail to get user by nicky[%s]", err)
		return
	}

	us, err := u.GetUserLikeNicky("jo")
	if err != nil || len(us) == 0 {
		t.Errorf("fail to get user by like [%s]", err)
		return
	}

}

//==================================== END ======================================
