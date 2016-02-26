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
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/oswystan/fixer/datastore"
	"github.com/oswystan/fixer/model"
)

func marshalResult(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if v == nil {
		w.WriteHeader(code)
		return
	}

	bs, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bs = append(bs, '\n')
	w.WriteHeader(code)
	w.Write(bs)
}

func ServeFilterTeam(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve team %s", r.URL.Path)

	var err error

	ret := &model.ResultTeam{}
	tl := datastore.NewStoreTeamList()
	query := r.URL.Query()
	for k, v := range query {
		switch k {
		case "id":
			id, _ := strconv.Atoi(v[0])
			ret.Team, err = tl.GetTeamById(id)
		case "name":
			ret.Team, err = tl.GetTeamByName(v[0])
		}
	}

	if err != nil {
		log.Printf("ERROR: %s", err)
		marshalResult(w, nil, http.StatusInternalServerError)
		return
	}

	marshalResult(w, ret, http.StatusOK)
}

func ServeFilterMemberList(w http.ResponseWriter, r *http.Request) {
	var err error

	l := &model.ResultMemberList{}
	tl := datastore.NewStoreTeamList()

	query := r.URL.Query()

	for k, v := range query {
		switch k {
		case "team_id":
			id, _ := strconv.Atoi(v[0])
			l.Members, err = tl.GetMemberList(id)
		case "_":
		default:
			err = fmt.Errorf("bad request")
		}
	}

	if err != nil {
		marshalResult(w, nil, http.StatusInternalServerError)
		return
	}

	marshalResult(w, l, http.StatusOK)
}

func ServeFilterTeamList(w http.ResponseWriter, r *http.Request) {
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
		marshalResult(w, nil, http.StatusInternalServerError)
		return
	}

	marshalResult(w, l, http.StatusOK)
}

func ServeFilterBugList(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve bug list %s", r.URL.Path)
}
func ServeUserDetail(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve userdetail %s", r.URL.Path)
	var err error
	us := datastore.NewStoreUser()
	ret := &model.ResultUserDetail{}

	query := r.URL.Query()
	for k, v := range query {
		switch k {
		case "id":
			id, _ := strconv.Atoi(v[0])
			ret.User, err = us.GetUserById(id)
		case "nicky":
			ret.User, err = us.GetUserByNicky(v[0])
		case "_":
		default:
			err = fmt.Errorf("bad request query")
		}
	}
	if err != nil {
		log.Printf("ERROR: %s", err)
		marshalResult(w, nil, http.StatusInternalServerError)
		return
	}

	marshalResult(w, ret, http.StatusOK)
}

//==================================== END ======================================
