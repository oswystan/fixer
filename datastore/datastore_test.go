//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: datastore_test.go
//     description:
//         created: 2016-02-26 15:44:42
//          author: wystan
//
//===============================================================================

package datastore

import "testing"

func openDB() error {
	db := GetDB()
	err := db.Open("pgtest", "123456", "fixer")
	return err
}
func closeDB() error {
	db := GetDB()
	err := db.Close()
	return err
}

func TestOpen(t *testing.T) {
	if err := openDB(); err != nil {
		t.Errorf("fail to open db [%s]", err)
		return
	}
	closeDB()
}

//==================================== END ======================================
