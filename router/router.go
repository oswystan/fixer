//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: router.go
//     description:
//         created: 2016-02-23 17:02:37
//          author: wystan
//
//===============================================================================
package router

import (
	"github.com/gorilla/mux"
	"github.com/oswystan/fixer/controller"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/login.html", controller.ShowLogin).Methods("GET")
	r.HandleFunc("/user.html", controller.ServeUser).Methods("POST")
	r.PathPrefix("/static/").HandlerFunc(controller.ServeStaticFile).Methods("GET")
	return r
}

//==================================== END ======================================
