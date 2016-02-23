//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: user.go
//     description:
//         created: 2016-02-23 21:49:11
//          author: wystan
//
//===============================================================================

package controller

import "net/http"

func ServeUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/user.html")
}

//==================================== END ======================================
