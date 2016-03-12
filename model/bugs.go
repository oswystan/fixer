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
	Id             int       `json:"id"`
	CreatedBy      int32     `json:"created_by"`
	CreatedNicky   string    `json:"created_nicky"`
	Title          string    `json:"title"`
	Priority       int       `json:"priority"`
	Detail         string    `json:"detail"`
	Attachments    string    `json:"attchments"`
	CurrentHandler int       `json:"current_handler"`
	HandlerNicky   string    `json:"handler_nicky"`
	Status         int32     `json:"status"`
	CreatedTime    time.Time `json:"created_time"`
	LastUpdate     time.Time `json:"last_update"`
}

type Buglog struct {
	BugId      int       `json:"bug_id"`
	Who        int       `json:"who"`
	ActionType int       `json:"action_type"`
	ActionTime time.Time `json:"action_time"`
	Action     string    `json:"action"`
}

//==================================== END ======================================
