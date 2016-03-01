//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: teams.go
//     description:
//         created: 2016-02-29 15:53:32
//          author: wystan
//
//===============================================================================

package datastore

import "github.com/oswystan/fixer/model"

var sqlTeamJoined = `select t.id, t.name, t.leader_id, get_nicky(t.leader_id) as leader_name, t.created_date 
				  from user_team as ut inner join team as t on ut.user_id = ? and ut.team_id = t.id
				  order by t.id offset ? limit ?;`
var sqlTeamCreated = `select t.id, t.name, t.leader_id, get_nicky(t.leader_id) as leader_name, t.created_date
                  from team as t where t.leader_id = ? order by t.id offset ? limit ?;`

type StoreTeams interface {
	GetTeamsJoined(f *model.Filter) ([]model.TeamCreatedOrJoined, error)
	GetTeamsCreated(f *model.Filter) ([]model.TeamCreatedOrJoined, error)
	GetMembers(tid int) ([]model.User, error)
	GetTeams(f *model.Filter) ([]model.Team, error)
	GetTeam(tid int) (*model.Team, error)
	Create(t *model.Team) (*model.Team, error)
	Update(t *model.Team) (*model.Team, error)
	Delete(t *model.Team) error
}

type storeTeams struct {
}

func (ds *storeTeams) GetTeamsJoined(f *model.Filter) ([]model.TeamCreatedOrJoined, error) {
	var tl []model.TeamCreatedOrJoined

	db := GetDB()
	_, err := db.pg.Query(&tl, sqlTeamJoined, f.UserId, f.Offset, f.Limit)
	if err != nil {
		return nil, err
	}

	return tl, nil
}

func (ds *storeTeams) GetTeamsCreated(f *model.Filter) ([]model.TeamCreatedOrJoined, error) {
	var tl []model.TeamCreatedOrJoined

	db := GetDB()
	_, err := db.pg.Query(&tl, sqlTeamCreated, f.UserId, f.Offset, f.Limit)
	if err != nil {
		return nil, err
	}
	return tl, nil
}
func (ds *storeTeams) GetMembers(tid int) ([]model.User, error) {
	return nil, nil
}
func (ds *storeTeams) GetTeams(f *model.Filter) ([]model.Team, error) {
	return nil, nil
}

func (ds *storeTeams) GetTeam(tid int) (*model.Team, error) {
	return nil, nil
}
func (ds *storeTeams) Create(t *model.Team) (*model.Team, error) {
	return nil, nil
}
func (ds *storeTeams) Update(t *model.Team) (*model.Team, error) {
	return nil, nil
}
func (ds *storeTeams) Delete(t *model.Team) error {
	return nil
}

func NewStoreTeams() StoreTeams {
	return &storeTeams{}
}

//==================================== END ======================================
