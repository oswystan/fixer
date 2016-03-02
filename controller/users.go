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
	"time"

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

	ds := datastore.NewUserStore()
	ul, err := ds.GetUsers(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, ul, http.StatusOK)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	ds := datastore.NewUserStore()
	if err := ds.DeleteUsers(); err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, nil, http.StatusOK)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	pwd := &model.UserPwd{}
	err := decodeBody(r, user, pwd)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}
	user.RegisterDate = time.Now()
	user.Pwd = pwd.Pwd

	ds := datastore.NewUserStore()
	newUser, err := ds.Create(user)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, newUser, http.StatusOK)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	pwd := &model.UserPwd{}
	err := decodeBody(r, user, pwd)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}
	user.RegisterDate = time.Now()
	if len(pwd.Pwd) != 0 {
		user.Pwd = pwd.Pwd
	}
	userId, _ := strconv.Atoi(mux.Vars(r)["id"])
	user.Id = userId

	ds := datastore.NewUserStore()
	newUser, err := ds.Update(user)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, newUser, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewUserStore()
	user, err := ds.GetUser(id)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, user, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ds := datastore.NewUserStore()
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
	ds := datastore.NewTeamStore()
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
	ds := datastore.NewTeamStore()
	tl, err := ds.GetTeamsCreated(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, tl, http.StatusOK)
}

//==================================== END ======================================
