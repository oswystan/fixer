//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: main_test.go
//     description:
//         created: 2016-03-19 18:37:26
//          author: wystan
//
//===============================================================================

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/oswystan/fixer/model"
)

type request struct {
	method string
	url    string
	data   interface{}
}
type response struct {
	code    int
	actual  interface{}
	target  interface{}
	compare func(a, t interface{}) bool
}

type pair struct {
	r request
	a response
}

func get(r *request, a *response) error {
	res, err := http.Get(r.url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != a.code {
		return fmt.Errorf("[%s %s]status code %d != %d", r.method, r.url, res.StatusCode, a.code)
	}
	if a.actual != nil && a.target != nil {
		err = decodeBody(res, a.actual)
		if err != nil {
			return err
		}
		if !a.compare(a.actual, a.target) {
			return fmt.Errorf("%v <!=> %v", a.actual, a.target)
		}
	}
	return nil
}
func post(r *request, a *response) error {
	return nil
}
func put(r *request, a *response) error {
	return nil
}
func del(r *request, a *response) error {
	return nil
}
func option(r *request, a *response) error {
	return nil
}
func do(r *request, a *response) error {
	switch r.method {
	case "GET":
		return get(r, a)
	case "POST":
		return post(r, a)
	case "PUT":
		return put(r, a)
	case "DELETE":
		return del(r, a)
	case "OPTION":
		return option(r, a)
	default:
		return fmt.Errorf("invalid method %s", r.method)
	}
	return nil
}
func decodeBody(res *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	return err
}

var ps = []pair{
	{
		r: request{method: "GET", url: "http://localhost:8000/teams/1", data: nil},
		a: response{
			code: 200, actual: &model.Team{},
			target: &model.Team{
				Id:         1,
				LeaderId:   1,
				Name:       "john-frog",
				LeaderName: "john",
			},
			compare: func(a, b interface{}) bool {
				l := a.(*model.Team)
				r := b.(*model.Team)
				return l.Id == r.Id && l.Name == r.Name && l.LeaderId == r.LeaderId && l.LeaderName == r.LeaderName
			},
		},
	},
}

func TestTeams(t *testing.T) {
	for i := 0; i < len(ps); i++ {
		err := do(&ps[i].r, &ps[i].a)
		if err != nil {
			t.Fatal(err)
		}
	}
}

//==================================== END ======================================
