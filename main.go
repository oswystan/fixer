//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: main.go
//     description:
//         created: 2016-02-23 16:57:19
//          author: wystan
//
//===============================================================================
package main

import (
	"net/http"

	"github.com/oswystan/fixer/router"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":8000", r)
}

//==================================== END ======================================
