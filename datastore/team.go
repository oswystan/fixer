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

import (
	"fmt"

	"github.com/oswystan/fixer/model"
)

var teamJoined = `select t.id, t.name, t.leader_id, u.nicky as leader_name, t.created_date 
				  from user_team as ut inner join team as t on ut.user_id = ? and ut.team_id = t.id
				  inner join users as u on t.leader_id = u.id order by t.id;`

var teamCreated = `select t.id, t.name, t.leader_id, u.nicky as leader_name, t.created_date
                  from team as t inner join users as u 
				  on t.leader_id = ? and t.leader_id = u.id order by t.id;`
var teamById = `select * from team where id = ?`
var teamByName = `select * from team where name = ?`
var teamLikeName = `select * from team where name like '%s%%'`

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
	GetTeamById(id int) (*model.Team, error)
	GetTeamByName(name string) (*model.Team, error)
	GetTeamLikeName(name string) ([]model.Team, error)
}

type teamlist struct {
	user StoreUser
}

func (tl *teamlist) GetTeamById(id int) (*model.Team, error) {
	t := &model.Team{}
	_, err := GetDB().pg.QueryOne(t, teamById, id)
	return t, err
}

func (tl *teamlist) GetTeamByName(name string) (*model.Team, error) {
	t := &model.Team{}
	_, err := GetDB().pg.QueryOne(t, teamByName, name)
	return t, err
}

func (tl *teamlist) GetTeamLikeName(name string) ([]model.Team, error) {
	var ts []model.Team
	sql := fmt.Sprintf(teamLikeName, name)
	_, err := GetDB().pg.Query(&ts, sql)
	return ts, err
}

func (tl *teamlist) GetTeamList(f *FilterTeamList) ([]model.TeamCreatedOrJoined, error) {
	var ts []model.TeamCreatedOrJoined
	var err error

	db := GetDB()

	if f.LeaderId > 0 {
		_, err = db.pg.Query(&ts, teamCreated, f.LeaderId)
	} else if f.MemberId > 0 {
		_, err = db.pg.Query(&ts, teamJoined, f.MemberId)
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
