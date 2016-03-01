//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: users.go
//     description:
//         created: 2016-02-28 19:05:49
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	f := model.NewFilter()
	err := decodeQuery(r.URL.RawQuery, f)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ds := datastore.NewStoreUsers()
	ul, err := ds.GetUsers(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, ul, http.StatusOK)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	ds := datastore.NewStoreUsers()
	if err := ds.DeleteUsers(); err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, nil, http.StatusOK)
}

func PostUser(w http.ResponseWriter, r *http.Request) {

}

func PutUser(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewStoreUsers()
	user, err := ds.GetUser(id)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, user, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewStoreUsers()
	err := ds.Delete(id)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, nil, http.StatusOK)
}

func GetUserTeamsJoined(w http.ResponseWriter, r *http.Request) {
	f := model.NewFilter()

	err := decodeQuery(r.URL.RawPath, f)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	f.UserId, _ = strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewStoreTeams()
	tl, err := ds.GetTeamsJoined(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, tl, http.StatusOK)
}

func GetUserTeamsCreated(w http.ResponseWriter, r *http.Request) {
	f := model.NewFilter()

	err := decodeQuery(r.URL.RawPath, f)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	f.UserId, _ = strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewStoreTeams()
	tl, err := ds.GetTeamsCreated(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, tl, http.StatusOK)
}

//==================================== END ======================================
