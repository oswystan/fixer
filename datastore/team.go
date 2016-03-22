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

var teamJoined = `select t.id, t.name, t.leader_id, get_nicky(t.leader_id) as leader_name, t.created_date 
				  from user_team as ut inner join team as t on ut.user_id = ? and ut.team_id = t.id
				  order by t.id;`

var teamCreated = `select t.id, t.name, t.leader_id, get_nicky(t.leader_id) as leader_name, t.created_date
                  from team as t where t.leader_id = ? order by t.id;`

var teamById = `select t.*, u.nicky as leader_name from team as t inner join users as u 
					on t.id = ? and u.id = t.leader_id;`
var teamByName = `select t.*, u.nicky as leader_name from team as t inner join users as u 
					on t.name = ? and u.id = t.leader_id;`
var teamLikeName = `select t.*, u.nicky as leader_name from team as t inner join users as u 
                    on t.name like '%s%%' and u.id = t.leader_id order by t.id;`
var memberById = `select u.* from user_team as ut inner join users as u on ut.team_id = ? and ut.user_id = u.id order by u.id;`

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
	GetMemberList(id int) ([]model.User, error)
	GetTeamById(id int) (*model.Team, error)
	GetTeamByName(name string) (*model.Team, error)
	GetTeamLikeName(name string) ([]model.Team, error)
}

type teamlist struct {
	user StoreUser
}

func (tl *teamlist) GetMemberList(id int) ([]model.User, error) {
	var ul []model.User
	_, err := GetDB().pg.Query(&ul, memberById, id)
	return ul, err
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
	var u *model.User

	db := GetDB()

	if f.LeaderId > 0 {
		_, err = db.pg.Query(&ts, teamCreated, f.LeaderId)
	} else if f.MemberId > 0 {
		_, err = db.pg.Query(&ts, teamJoined, f.MemberId)
	} else if len(f.LeaderName) > 0 {
		u, err = tl.user.GetUserByNicky(f.LeaderName)
		if err == nil && u != nil {
			_, err = db.pg.Query(&ts, teamCreated, u.Id)
		}
	} else if len(f.MemberName) > 0 {
		u, err = tl.user.GetUserByNicky(f.MemberName)
		if err == nil && u != nil {
			_, err = db.pg.Query(&ts, teamJoined, u.Id)
		}
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
