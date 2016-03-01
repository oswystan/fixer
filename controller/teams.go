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
	"net/http"
	"strconv"

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

	ds := datastore.NewStoreTeams()
	tl, err := ds.GetTeams(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, tl, http.StatusOK)
}

func PostTeam(w http.ResponseWriter, r *http.Request) {

}

func DeleteTeams(w http.ResponseWriter, r *http.Request) {

}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	tid, _ := strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewStoreTeams()
	t, err := ds.GetTeam(tid)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, t, http.StatusOK)
}

func PutTeam(w http.ResponseWriter, r *http.Request) {

}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {

}

func GetTeamUsers(w http.ResponseWriter, r *http.Request) {
	f := model.NewFilter()
	err := decodeQuery(r.URL.RawPath, f)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	f.TeamId, _ = strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewStoreTeams()
	ul, err := ds.GetMembers(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, ul, http.StatusOK)
}

func DeleteTeamUsers(w http.ResponseWriter, r *http.Request) {

}

func PutTeamUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteTeamUser(w http.ResponseWriter, r *http.Request) {

}

//==================================== END ======================================
