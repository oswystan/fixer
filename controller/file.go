//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: file.go
//     description:
//         created: 2016-03-24 16:37:36
//          author: wystan
//
//===============================================================================

package controller

import "net/http"

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/login.html")
}

func ServeUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/user.html")
}

func ServeFixer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/fixer.html")
}

//==================================== END ======================================
