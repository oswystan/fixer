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

import (
	"fmt"

	"github.com/oswystan/fixer/model"
)

var sqlTeamJoined = `select t.id, t.name, t.leader_id, get_nicky(t.leader_id) as leader_name, t.created_date 
				from user_team as ut inner join team as t on ut.user_id = ? and ut.team_id = t.id
				order by t.id offset ? limit ?;`
var sqlTeamCreated = `select t.id, t.name, t.leader_id, get_nicky(t.leader_id) as leader_name, t.created_date
				from team as t where t.leader_id = ? order by t.id offset ? limit ?;`
var sqlTeamMembers = `select u.id, u.nicky, u.portrait, u.email, u.last_login_time, u.register_date 
				from users as u inner join user_team as ut on ut.team_id = ? and u.id = ut.user_id offset ? limit ?;`
var sqlTeam = `select *, get_nicky(t.leader_id) as leader_name from team as t where id = ?`
var sqlTeams = `select *, get_nicky(t.leader_id) as leader_name from team as t where name like '%s%%' order by id offset ? limit ?`
var sqlCreateTeam = `insert into team(name, leader_id, goal, created_date, bug_table, bug_table_status, status, logo)
				values(?name, ?leader_id, ?goal, ?created_date, ?bug_table, ?bug_table_status, ?status, ?logo) returning *`
var sqlUpdateTeam = `update team set name=?name, leader_id=?leader_id, goal=?goal, created_date=?created_date, 
				bug_table=?bug_table, bug_table_status=?bug_table_status, status=?status, logo=?logo) returning *`
var sqlDelTeam = `delete from team where id=?`
var sqlDelTeams = `delete from team`

type TeamStore interface {
	GetTeamsJoined(f *model.Filter) ([]model.TeamCreatedOrJoined, error)
	GetTeamsCreated(f *model.Filter) ([]model.TeamCreatedOrJoined, error)
	GetMembers(f *model.Filter) ([]model.User, error)
	GetTeams(f *model.Filter) ([]model.Team, error)
	GetTeam(tid int) (*model.Team, error)
	Create(t *model.Team) (*model.Team, error)
	Update(t *model.Team) (*model.Team, error)
	Delete(int) error
	DeleteAll() error
}

type teamstore struct {
}

func (ds *teamstore) GetTeamsJoined(f *model.Filter) ([]model.TeamCreatedOrJoined, error) {
	var tl []model.TeamCreatedOrJoined

	db := GetDB()
	_, err := db.pg.Query(&tl, sqlTeamJoined, f.UserId, f.Offset, f.Limit)
	if err != nil {
		return nil, err
	}

	return tl, nil
}

func (ds *teamstore) GetTeamsCreated(f *model.Filter) ([]model.TeamCreatedOrJoined, error) {
	var tl []model.TeamCreatedOrJoined

	db := GetDB()
	_, err := db.pg.Query(&tl, sqlTeamCreated, f.UserId, f.Offset, f.Limit)
	if err != nil {
		return nil, err
	}
	return tl, nil
}
func (ds *teamstore) GetMembers(f *model.Filter) ([]model.User, error) {
	var ul []model.User

	db := GetDB()
	_, err := db.pg.Query(&ul, sqlTeamMembers, f.TeamId, f.Offset, f.Limit)
	if err != nil {
		return nil, err
	}
	return ul, nil
}
func (ds *teamstore) GetTeams(f *model.Filter) ([]model.Team, error) {
	var tl []model.Team
	db := GetDB()

	sql := fmt.Sprintf(sqlTeams, f.Q)
	_, err := db.pg.Query(&tl, sql, f.Offset, f.Limit)
	if err != nil {
		return nil, err
	}
	return tl, nil
}

func (ds *teamstore) GetTeam(tid int) (*model.Team, error) {
	t := &model.Team{}
	db := GetDB()
	_, err := db.pg.QueryOne(t, sqlTeam, tid)
	return t, err
}
func (ds *teamstore) Create(t *model.Team) (*model.Team, error) {
	db := GetDB()
	_, err := db.pg.QueryOne(t, sqlCreateTeam, t)
	return t, err
}
func (ds *teamstore) Update(t *model.Team) (*model.Team, error) {
	db := GetDB()
	_, err := db.pg.QueryOne(t, sqlCreateTeam, t)
	return t, err
}
func (ds *teamstore) Delete(tid int) error {
	db := GetDB()
	_, err := db.pg.Exec(sqlDelTeam, tid)
	return err
}
func (ds *teamstore) DeleteAll() error {
	db := GetDB()
	_, err := db.pg.Exec(sqlDelTeams)
	return err
}

func NewTeamStore() TeamStore {
	return &teamstore{}
}

//==================================== END ======================================
