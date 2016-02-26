//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: team.go
//     description:
//         created: 2016-02-24 20:26:08
//          author: wystan
//
//===============================================================================

package datastore

import "github.com/oswystan/fixer/model"

var sqlJoined = `select t.id, t.name, t.leader_id, u.nicky as leader_name, t.created_date 
				  from user_team as ut inner join team as t on ut.team_id = t.id 
				  inner join users as u on t.leader_id = u.id where ut.user_id = ? order by t.id;`

var sqlCreated = `select t.id, t.name, t.leader_id, u.nicky as leader_name, t.created_date
                  from team as t inner join users as u 
				  on t.leader_id = u.id where t.leader_id = ? order by t.id;`

type FilterTeamList struct {
	LeaderId   int
	LeaderName string
	MemberId   int
	MemberName string
	from       uint
	count      int
}

func (f *FilterTeamList) Reset() {
	f.LeaderId = 0
	f.MemberId = 0
	f.LeaderName = ""
	f.MemberName = ""
}

type StoreTeamList interface {
	GetTeamList(filter *FilterTeamList) ([]model.TeamCreatedOrJoined, error)
}

type teamlist struct {
	user StoreUser
}

func (tl *teamlist) GetTeamList(f *FilterTeamList) ([]model.TeamCreatedOrJoined, error) {
	var ts []model.TeamCreatedOrJoined
	var err error

	db := GetDB()

	if f.LeaderId > 0 {
		_, err = db.pg.Query(&ts, sqlCreated, f.LeaderId)
	} else if f.MemberId > 0 {
		_, err = db.pg.Query(&ts, sqlJoined, f.MemberId)
	}
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func NewStoreTeamList() StoreTeamList {
	return &teamlist{user: NewStoreUser()}
}

//==================================== END ======================================
