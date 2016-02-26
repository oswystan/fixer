//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: team_test.go
//     description:
//         created: 2016-02-26 10:10:17
//          author: wystan
//
//===============================================================================

package datastore

import "testing"

//func connectDB() error {
//    db := GetDB()
//    err := db.Open("pgtest", "123456", "fixer")
//    return err
//}
//func closeDB() error {
//    db := GetDB()
//    err := db.Close()
//    return err
//}

func TestGetTeamList(t *testing.T) {
	if err := connectDB(); err != nil {
		t.Error("fail to connect db")
	}
	defer closeDB()

	filter := &FilterTeamList{}
	filter.LeaderId = 1
	tl := NewStoreTeamList()

	ts, err := tl.GetTeamList(filter)
	if err != nil {
		t.Errorf("fail to get team list created [%s]", err)
		return
	}
	if len(ts) <= 0 {
		t.Errorf("none created team list returned")
		return
	}

	filter.Reset()
	filter.MemberId = 1
	ts, err = tl.GetTeamList(filter)
	if err != nil {
		t.Errorf("fail to get team list joined [%s]", err)
		return
	}
	if len(ts) <= 0 {
		t.Errorf("none joined team list returned")
		return
	}
}

//==================================== END ======================================
