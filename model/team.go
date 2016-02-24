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
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Leader     string    `json:"leader"`
	CreateDate time.Time `json:"create_date"`
	Status     string    `json:"status"`
	BugTab     string    `json:"bug_tab"`
	Goal       string    `json:"goal"`
}

type ResultTeamList struct {
	Error int    `json:"error"`
	Teams []Team `json:"teams"`
}

type ResultTeam struct {
	Error int  `json:"error"`
	Team  Team `json:"team"`
}

//==================================== END ======================================
