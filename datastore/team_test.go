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

func TestGetTeamList(t *testing.T) {
	if err := openDB(); err != nil {
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

func TestTeamGetBy(t *testing.T) {
	if err := openDB(); err != nil {
		t.Fatal("fail to connect db")
	}
	defer closeDB()

	tl := NewStoreTeamList()
	tm, err := tl.GetTeamById(1)
	if err != nil || tm == nil {
		t.Fatalf("fail to get team by id [%s]", err)
	}

	tm, err = tl.GetTeamByName("john-frog")
	if err != nil || tm == nil {
		t.Fatalf("fail to get team by name [%s]", err)
	}
	ts, err := tl.GetTeamLikeName("john")
	if err != nil || len(ts) == 0 {
		t.Fatalf("fail to get team like name [%s]", err)
	}
}

//==================================== END ======================================
