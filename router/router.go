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

	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users", controller.PostUser).Methods("POST")
	r.HandleFunc("/users", controller.DeleteUsers).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", controller.PutUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", controller.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}/teams/joined", controller.GetUserTeamsJoined).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/teams/created", controller.GetUserTeamsCreated).Methods("GET")

	r.HandleFunc("/teams", controller.GetTeams).Methods("GET")
	r.HandleFunc("/teams", controller.PostTeam).Methods("POST")
	r.HandleFunc("/teams", controller.DeleteTeams).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}", controller.GetTeam).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}", controller.PutTeam).Methods("PUT")
	r.HandleFunc("/teams/{id:[0-9]+}", controller.DeleteTeam).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/users", controller.GetTeamUsers).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/users", controller.DeleteTeamUsers).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/users/{uid:[0-9]+}", controller.PutTeamUser).Methods("PUT")
	r.HandleFunc("/teams/{id:[0-9]+}/users/{uid:[0-9]+}", controller.DeleteTeamUser).Methods("DELETE")

	r.HandleFunc("/teams/{id:[0-9]+}/bugs", controller.GetBugs).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs", controller.PostBug).Methods("POST")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs", controller.DeleteBugs).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}", controller.GetBug).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}", controller.PutBug).Methods("PUT")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}", controller.DeleteBug).Methods("DELETE")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}/logs", controller.GetBuglog).Methods("GET")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}/logs", controller.PostBuglog).Methods("POST")
	r.HandleFunc("/teams/{id:[0-9]+}/bugs/{bid:[0-9]+}/logs", controller.DeleteBuglog).Methods("DELETE")

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
