//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: teams.go
//     description:
//         created: 2016-02-29 00:02:00
//          author: wystan
//
//===============================================================================

package controller

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/oswystan/fixer/datastore"
	"github.com/oswystan/fixer/model"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	f := model.NewFilter()
	err := decodeQuery(r.URL.RawQuery, f)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ds := datastore.NewTeamStore()
	tl, err := ds.GetTeams(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, tl, http.StatusOK)
}

func PostTeam(w http.ResponseWriter, r *http.Request) {
	team := &model.Team{}
	err := decodeBody(r, team)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}
	team.CreatedDate = time.Now()
	str := fmt.Sprintf("%s%s", team.Name, team.CreatedDate)
	hdata := sha1.Sum([]byte(str))
	team.BugTable = fmt.Sprintf("%x", hdata)

	ds := datastore.NewTeamStore()
	newTeam, err := ds.Create(team)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, newTeam, http.StatusOK)
}

func PutTeam(w http.ResponseWriter, r *http.Request) {
	team := &model.Team{}
	err := decodeBody(r, team)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}
	team.Id, _ = strconv.Atoi(mux.Vars(r)["id"])

	ds := datastore.NewTeamStore()
	newTeam, err := ds.Update(team)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, newTeam, http.StatusOK)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewTeamStore()
	err := ds.Delete(id)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, nil, http.StatusOK)
}
func DeleteTeams(w http.ResponseWriter, r *http.Request) {
	ds := datastore.NewTeamStore()
	err := ds.DeleteAll()
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, nil, http.StatusOK)
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	tid, _ := strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewTeamStore()
	t, err := ds.GetTeam(tid)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, t, http.StatusOK)
}

func GetTeamUsers(w http.ResponseWriter, r *http.Request) {
	f := model.NewFilter()
	err := decodeQuery(r.URL.RawPath, f)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	f.TeamId, _ = strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewTeamStore()
	ul, err := ds.GetMembers(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, ul, http.StatusOK)
}

func DeleteTeamUsers(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewTeamStore()
	err := ds.DeleteMembers(id)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, nil, http.StatusOK)
}

func PutTeamUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	uid, _ := strconv.Atoi(mux.Vars(r)["uid"])

	ds := datastore.NewTeamStore()
	err := ds.AddMember(id, uid)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, nil, http.StatusOK)
}

func DeleteTeamUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	uid, _ := strconv.Atoi(mux.Vars(r)["uid"])

	ds := datastore.NewTeamStore()
	err := ds.DeleteMember(id, uid)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, nil, http.StatusOK)
}

//==================================== END ======================================
