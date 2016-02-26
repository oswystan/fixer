//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: team.go
//     description:
//         created: 2016-02-23 22:28:39
//          author: wystan
//
//===============================================================================

package model

import "time"

type Team struct {
	Id             int32     `json:"id"`
	Name           string    `json:"name"`
	LeaderId       int32     `json:"leader_id"`
	LeaderName     int32     `json:"leader_name"`
	Goal           string    `json:"goal"`
	CreatedDate    time.Time `json:"created_date"`
	BugTable       string    `json:"bug_table"`
	BugTableStatus int32     `json:"bug_table_status"`
	Status         int32     `json:"status"`
	Logo           string    `json:"logo"`
}

type TeamCreatedOrJoined struct {
	Id          int32     `json:"id"`
	Name        string    `json:"name"`
	LeaderId    int32     `json:"leader_id"`
	LeaderName  string    `json:"leader_name"`
	CreatedDate time.Time `json:"created_date"`
}

type ResultTeamList struct {
	Error int                   `json:"error"`
	Teams []TeamCreatedOrJoined `json:"teams"`
}

type ResultTeam struct {
	Error int   `json:"error"`
	Team  *Team `json:"team"`
}

//==================================== END ======================================
