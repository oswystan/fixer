//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: bugs.go
//     description:
//         created: 2016-02-26 22:04:19
//          author: wystan
//
//===============================================================================

package model

import "time"

type Bug struct {
	Id             int32     `json:"id"`
	CreatedBy      int32     `json:"created_by"`
	Title          string    `json:"title"`
	Priority       int32     `json:"priority"`
	Desc           string    `json:"desc"`
	Attachments    string    `json:"attchments"`
	CurrentHandler string    `json:"current_handler"`
	Status         int32     `json:"status"`
	LastUpdate     time.Time `json:"last_update"`
}

type Buglog struct {
	BugId     int32     `json:"bug_id"`
	Who       int32     `json:"who"`
	Type      int32     `json:"type"`
	Operation string    `json:"operation"`
	When      time.Time `json:"when"`
}

//==================================== END ======================================
