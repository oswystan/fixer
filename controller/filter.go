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
	"strconv"

	"github.com/oswystan/fixer/datastore"
	"github.com/oswystan/fixer/model"
)

func marshalResult(w http.ResponseWriter, v interface{}, code int) {
	bs, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bs = append(bs, '\n')
	w.WriteHeader(code)
	w.Write(bs)
}

func ServeFilterMemberList(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve member list %s", r.URL.Path)
}

func ServeFilterTeamList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var err error

	l := &model.ResultTeamList{}
	tl := datastore.NewStoreTeamList()
	filter := &datastore.FilterTeamList{}

	query := r.URL.Query()

	for k, v := range query {
		switch k {
		case "creatorid":
			filter.LeaderId, _ = strconv.Atoi(v[0])
		case "creator":
			filter.LeaderName = v[0]

		case "memberid":
			filter.MemberId, _ = strconv.Atoi(v[0])
		case "member":
			filter.MemberName = v[0]
		}
	}

	l.Teams, err = tl.GetTeamList(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshalResult(w, l, http.StatusOK)
}

func ServeFilterBugList(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve bug list %s", r.URL.Path)
}

//==================================== END ======================================
