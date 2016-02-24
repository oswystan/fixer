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
	"log"
	"net/http"

	"github.com/oswystan/fixer/datastore"
	"github.com/oswystan/fixer/router"
)

func main() {
	log.Printf("start server ...")

	db := datastore.GetDB()
	err := db.Open("pgtest", "123456", "fixer")
	if err != nil {
		log.Printf("ERROR: %s", err)
		return
	}
	log.Printf("database connected.")

	r := router.NewRouter()
	http.ListenAndServe(":8000", r)

	db.Close()
	return
}

//==================================== END ======================================
