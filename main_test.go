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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/oswystan/fixer/model"
)

type httperr struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	More        string `json:"more"`
}

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

func fmterr(r *request, a *response, res *http.Response) error {
	desc := &httperr{}
	_ = decodeBody(res, desc)
	return fmt.Errorf("[%s %s]status code %d != %d detail[%s]",
		r.method, r.url, res.StatusCode, a.code, desc.Description)
}

func get(r *request, a *response) error {
	res, err := http.Get(r.url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != a.code {
		return fmterr(r, a, res)
	}
	if a.actual != nil && a.target != nil {
		err = decodeBody(res, a.actual)
		if err != nil {
			return err
		}
		if !a.compare(a.actual, a.target) {
			return fmt.Errorf("[%s %s]%v <!=> %v", r.method, r.url, a.actual, a.target)
		}
	}
	return nil
}
func post(r *request, a *response) error {
	b, _ := json.Marshal(r.data)
	res, err := http.Post(r.url, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != a.code {
		return fmterr(r, a, res)
	}

	if a.actual != nil && a.target != nil {
		err = decodeBody(res, a.actual)
		if err != nil {
			return err
		}
		if !a.compare(a.actual, a.target) {
			return fmt.Errorf("[%s %s]%v <!=> %v", r.method, r.url, a.actual, a.target)
		}
	}

	return nil
}
func put(r *request, a *response) error {
	b, _ := json.Marshal(r.data)
	req, _ := http.NewRequest(r.method, r.url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != a.code {
		return fmterr(r, a, res)
	}

	if a.actual != nil && a.target != nil {
		err = decodeBody(res, a.actual)
		if err != nil {
			return err
		}
		if !a.compare(a.actual, a.target) {
			return fmt.Errorf("[%s %s]%v <!=> %v", r.method, r.url, a.actual, a.target)
		}
	}

	return nil
}
func del(r *request, a *response) error {
	req, _ := http.NewRequest(r.method, r.url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != a.code {
		return fmterr(r, a, res)
	}

	return nil
}
func option(r *request, a *response) error {
	return fmt.Errorf("unsupportted option operation.")
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

var teampair = []pair{
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
	{
		r: request{method: "GET", url: "http://localhost:8000/teams", data: nil},
		a: response{
			code: 200, actual: nil,
		},
	},
	{
		r: request{method: "GET", url: "http://localhost:8000/teams?q=john", data: nil},
		a: response{
			code: 200, actual: nil,
		},
	},
	{
		r: request{
			method: "POST",
			url:    "http://localhost:8000/teams",
			data: &model.Team{
				Name:     "john-funny",
				LeaderId: 1,
				Goal:     "make a funny team",
				Logo:     "static/images/1.jpg",
			},
		},
		a: response{
			code: 201, actual: nil,
		},
	},
	{
		r: request{
			method: "PUT",
			url:    "http://localhost:8000/teams/6",
			data: &model.Team{
				Name:     "john-funny-wy",
				LeaderId: 1,
				Goal:     "make a funny team",
				Logo:     "static/images/1.jpg",
			},
		},
		a: response{
			code: 200, actual: nil,
		},
	},
	{
		r: request{
			method: "DELETE",
			url:    "http://localhost:8000/teams/6",
		},
		a: response{
			code: 200, actual: nil,
		},
	},
}

var userpair = []pair{
	{
		r: request{method: "GET", url: "http://localhost:8000/users", data: nil},
		a: response{
			code: 200, actual: nil,
		},
	},

	{
		r: request{method: "GET", url: "http://localhost:8000/users/1", data: nil},
		a: response{
			code: 200, actual: nil,
		},
	},

	{
		r: request{
			method: "DELETE",
			url:    "http://localhost:8000/users",
		},
		a: response{
			code: 200, actual: nil,
		},
	},

	{
		r: request{
			method: "POST",
			url:    "http://localhost:8000/users",
			data: &model.User{
				Nicky:    "wystan",
				Pwd:      "123456",
				Portrait: "static/images/1.jpg",
				Email:    "wystan@11.com",
			},
		},
		a: response{
			code: 201, actual: nil,
		},
	},
}

func TestTeams(t *testing.T) {
	for i := 0; i < len(teampair); i++ {
		err := do(&teampair[i].r, &teampair[i].a)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUsers(t *testing.T) {
	for i := 0; i < len(userpair); i++ {
		err := do(&userpair[i].r, &userpair[i].a)
		if err != nil {
			t.Error(err)
		}
	}

}

//==================================== END ======================================
