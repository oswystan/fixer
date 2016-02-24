//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: filter.go
//     description:
//         created: 2016-02-23 22:14:27
//          author: wystan
//
//===============================================================================

package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/oswystan/fixer/model"
)

func ServeFilterMemberList(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve member list %s", r.URL.Path)
}

func ServeFilterTeamList(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve team list %s", r.URL.Path)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	l := &model.ResultTeamList{}
	l.Teams = make([]model.Team, 1)
	l.Teams[0].Id = 1
	l.Teams[0].Name = "jugar"
	l.Teams[0].Leader = "wystan"
	l.Teams[0].Status = "ACTIVE"
	l.Teams[0].BugTab = "ef001233"
	l.Teams[0].Goal = "this is a team for wystan."
	l.Teams[0].CreateDate = time.Now()

	bs, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bs = append(bs, '\n')
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func ServeFilterBugList(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve bug list %s", r.URL.Path)
}

//==================================== END ======================================
