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

	r.HandleFunc("/filter/team.html", controller.ServeFilterTeam).Methods("GET")
	r.HandleFunc("/filter/user-detail.html", controller.ServeFilterUser).Methods("GET")
	r.HandleFunc("/filter/teamlist.html", controller.ServeFilterTeamList).Methods("GET")
	r.HandleFunc("/filter/memberlist.html", controller.ServeFilterMemberList).Methods("GET")
	r.HandleFunc("/filter/buglist.html", controller.ServeFilterBugList).Methods("GET")

	r.HandleFunc("/users", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/users", controller.ServeUnsupported).Methods("POST")
	r.HandleFunc("/users", controller.ServeUnsupported).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", controller.ServeUnsupported).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", controller.ServeUnsupported).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}/teams/joined", controller.GetUserTeamsJoined).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/teams/created", controller.GetUserTeamsCreated).Methods("GET")

	r.HandleFunc("/teams", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/teams", controller.ServeUnsupported).Methods("POST")
	r.HandleFunc("/teams", controller.ServeUnsupported).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}", controller.ServeUnsupported).Methods("PUT")
	r.HandleFunc("/teams/{id:[0-9]+}", controller.ServeUnsupported).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/users", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/users", controller.ServeUnsupported).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/users/{uid:[0-9]+}", controller.ServeUnsupported).Methods("PUT")
	r.HandleFunc("/teams/{id:[0-9]+}/users/{uid:[0-9]+}", controller.ServeUnsupported).Methods("DELETE")

	r.HandleFunc("/teams/{id:[0-9]+}/bugs", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs", controller.ServeUnsupported).Methods("POST")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs", controller.ServeUnsupported).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}", controller.ServeUnsupported).Methods("PUT")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}", controller.ServeUnsupported).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}/logs", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}/logs", controller.ServeUnsupported).Methods("POST")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}/logs", controller.ServeUnsupported).Methods("DELETE")

	r.HandleFunc("/teams/{id:[0-9]+}/templates", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/templates", controller.ServeUnsupported).Methods("POST")
	r.HandleFunc("/teams/{id:[0-9]+}/templates", controller.ServeUnsupported).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/templates/{tid:[0-9]+}", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/templates/{tid:[0-9]+}", controller.ServeUnsupported).Methods("PUT")
	r.HandleFunc("/teams/{id:[0-9]+}/templates/{tid:[0-9]+}", controller.ServeUnsupported).Methods("DELETE")

	r.HandleFunc("/stats/users", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/stats/teams", controller.ServeUnsupported).Methods("GET")
	r.HandleFunc("/stats/teams/{id:[0-9]+}/bugs", controller.ServeUnsupported).Methods("GET")

	r.HandleFunc("/syslog", controller.ServeUnsupported).Methods("GET")
	return r
}

//==================================== END ======================================
