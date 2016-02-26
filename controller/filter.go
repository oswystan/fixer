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

	"github.com/oswystan/fixer/datastore"
	"github.com/oswystan/fixer/model"
)

func ServeFilterMemberList(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve member list %s", r.URL.Path)
}

func ServeFilterTeamList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var err error

	l := &model.ResultTeamList{}
	tl := datastore.NewStoreTeamList()
	filter := &datastore.FilterTeamList{LeaderId: 1}

	l.Teams, err = tl.GetTeamList(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
