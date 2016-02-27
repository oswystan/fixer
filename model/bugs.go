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
	BugId     int32     `json:"bug_id"`
	Who       int32     `json:"who"`
	Type      int32     `json:"type"`
	Operation string    `json:"operation"`
	When      time.Time `json:"when"`
}

//==================================== END ======================================
