//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: login.go
//     description:
//         created: 2016-02-23 17:56:53
//          author: wystan
//
//===============================================================================

package controller

import "net/http"

func ShowLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/login.html")
}

//==================================== END ======================================
