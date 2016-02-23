//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: static.go
//     description:
//         created: 2016-02-23 19:47:57
//          author: wystan
//
//===============================================================================

package controller

import (
	"log"
	"net/http"
	"os"
)

func ServeStaticFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get static file %s", r.URL.Path)
	fname := r.URL.Path[1:]
	f, err := os.Stat(fname)
	if err != nil || f.IsDir() {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, fname)
}

//==================================== END ======================================
