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
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/oswystan/fixer/datastore"
	"github.com/oswystan/fixer/model"
)

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
		Json(w, nil, http.StatusInternalServerError)
		return
	}

	Json(w, ret, http.StatusOK)
}

func ServeFilterMemberList(w http.ResponseWriter, r *http.Request) {
	var err error

	l := &model.ResultMemberList{}
	tl := datastore.NewStoreTeamList()

	query := r.URL.Query()

LOOP:
	for k, v := range query {
		switch k {
		case "team_id":
			id, _ := strconv.Atoi(v[0])
			l.Members, err = tl.GetMemberList(id)
			if err != nil {
				break LOOP
			}
		case "_":
		default:
			err = fmt.Errorf("bad request")
		}
	}

	if err != nil {
		Json(w, nil, http.StatusInternalServerError)
		return
	}

	Json(w, l, http.StatusOK)
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
		Json(w, nil, http.StatusInternalServerError)
		return
	}

	Json(w, l, http.StatusOK)
}

func ServeFilterUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve userdetail %s", r.URL.Path)
	var err error
	us := datastore.NewStoreUser()
	ret := &model.ResultUserDetail{}

	query := r.URL.Query()

LOOP:
	for k, v := range query {
		switch k {
		case "id":
			id, _ := strconv.Atoi(v[0])
			ret.User, err = us.GetUserById(id)
			if err != nil {
				break LOOP
			}
		case "nicky":
			ret.User, err = us.GetUserByNicky(v[0])
			if err != nil {
				break LOOP
			}
		case "_":
		default:
			err = fmt.Errorf("bad request query")
			break LOOP
		}
	}
	if err != nil {
		log.Printf("ERROR: %s", err)
		Json(w, nil, http.StatusInternalServerError)
		return
	}

	Json(w, ret, http.StatusOK)
}

func ServeFilterBugList(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve bug list %s", r.URL.Path)
	var err error
	f := &datastore.BugFilter{}

	query := r.URL.Query()
	bugs := datastore.NewStoreBugs()

LOOP:
	for k, v := range query {
		switch k {
		case "team_id":
			f.TeamId, err = strconv.Atoi(v[0])
			if err != nil {
				break LOOP
			}
		case "priority":
			f.Priority, err = strconv.Atoi(v[0])
			if err != nil {
				break LOOP
			}
		case "handler":
			f.Handler, err = strconv.Atoi(v[0])
			if err != nil {
				break LOOP
			}
		case "created_by":
			f.CreatedBy, err = strconv.Atoi(v[0])
			if err != nil {
				break LOOP
			}
		case "status":
			f.Status, err = strconv.Atoi(v[0])
			if err != nil {
				break LOOP
			}
		case "offset":
			f.Offset, err = strconv.Atoi(v[0])
			if err != nil {
				break LOOP
			}
		case "count":
			f.Count, err = strconv.Atoi(v[0])
			if err != nil {
				break LOOP
			}
		case "date_from":
			f.DateFrom, err = time.Parse("2006-01-02", v[0])
			if err != nil {
				break LOOP
			}
		case "date_to":
			f.DateTo, err = time.Parse("2006-01-02", v[0])
			if err != nil {
				break LOOP
			}
		case "_":
		default:
			err = fmt.Errorf("invalid query string [%s]", k)
		}
	}

	if err != nil {
		log.Printf("ERROR: %s [query=%s]", err, r.URL.RawQuery)
		Json(w, nil, http.StatusBadRequest)
		return
	}

	bl, err := bugs.GetBugs(f)
	if err != nil {
		log.Printf("ERROR: %s", err)
		Json(w, nil, http.StatusInternalServerError)
		return
	}

	Json(w, bl, http.StatusOK)
}

//==================================== END ======================================
