//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: bugs.go
//     description:
//         created: 2016-02-29 00:02:29
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

func GetBugs(w http.ResponseWriter, r *http.Request) {
	f := model.NewFilter()
	err := decodeQuery(r.URL.RawQuery, f)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ds := datastore.NewBugStore()
	tid, _ := strconv.Atoi(mux.Vars(r)["id"])
	err = ds.SetTeam(tid)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	list, err := ds.GetBugs(f)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, list, http.StatusOK)
}

func PostBug(w http.ResponseWriter, r *http.Request) {
	bug := &model.Bug{}
	err := decodeBody(r, bug)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ds := datastore.NewBugStore()
	tid, _ := strconv.Atoi(mux.Vars(r)["id"])
	err = ds.SetTeam(tid)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newBug, err := ds.Create(bug)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, newBug, http.StatusOK)
}

func DeleteBugs(w http.ResponseWriter, r *http.Request) {
	ds := datastore.NewBugStore()
	tid, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := ds.SetTeam(tid)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = ds.DeleteBugs()
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, nil, http.StatusOK)
}

func GetBug(w http.ResponseWriter, r *http.Request) {
	ds := datastore.NewBugStore()
	tid, _ := strconv.Atoi(mux.Vars(r)["id"])
	bid, _ := strconv.Atoi(mux.Vars(r)["bid"])
	err := ds.SetTeam(tid)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	bug, err := ds.GetBug(bid)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, bug, http.StatusOK)
}

func PutBug(w http.ResponseWriter, r *http.Request) {
	bug := &model.Bug{}
	err := decodeBody(r, bug)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ds := datastore.NewBugStore()
	tid, _ := strconv.Atoi(mux.Vars(r)["id"])
	bug.Id, _ = strconv.Atoi(mux.Vars(r)["bid"])
	err = ds.SetTeam(tid)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newBug, err := ds.Update(bug)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, newBug, http.StatusOK)
}

func DeleteBug(w http.ResponseWriter, r *http.Request) {
	ds := datastore.NewBugStore()
	tid, _ := strconv.Atoi(mux.Vars(r)["id"])
	bid, _ := strconv.Atoi(mux.Vars(r)["bid"])
	err := ds.SetTeam(tid)
	if err != nil {
		JsonErr(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = ds.Delete(bid)
	if err != nil {
		JsonErr(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, nil, http.StatusOK)
}

func GetBuglog(w http.ResponseWriter, r *http.Request) {

}

func PostBuglog(w http.ResponseWriter, r *http.Request) {

}

func DeleteBuglog(w http.ResponseWriter, r *http.Request) {

}

//==================================== END ======================================
