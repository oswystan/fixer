//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: unsupported.go
//     description:
//         created: 2016-02-28 17:53:41
//          author: wystan
//
//===============================================================================
package controller

import (
	"fmt"
	"log"
	"net/http"
)

type unsupportedMsg struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Url    string `json:"url"`
}

var unsupported = &unsupportedMsg{
	Code:   -1,
	Detail: "",
	Url:    "",
}

func ServeUnsupported(w http.ResponseWriter, r *http.Request) {
	log.Printf("Unsupported %s on %s", r.Method, r.URL.Path)
	unsupported.Url = r.URL.RawPath
	unsupported.Detail = fmt.Sprintf("[%s on %s]NOT supported now, coming soon.", r.Method, r.URL.Path)
	Json(w, unsupported, http.StatusForbidden)
}

//==================================== END ======================================
