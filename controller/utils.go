//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: utils.go
//     description:
//         created: 2016-02-29 00:10:05
//          author: wystan
//
//===============================================================================

package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

type FixerError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	More        string `json:"more"`
}

func Json(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if v == nil {
		w.WriteHeader(code)
		return
	}

	bs, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bs = append(bs, '\n')
	w.WriteHeader(code)
	w.Write(bs)
}

func JsonErr(w http.ResponseWriter, r *http.Request, code int, desc string) {
	fe := &FixerError{code, desc, r.URL.String()}
	Json(w, fe, code)
}

func decodeQuery(qstr string, out interface{}) error {
	qmap, err := url.ParseQuery(qstr)
	if err != nil {
		return err
	}

	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return fmt.Errorf("Not settable out value.")
	}

	typ := reflect.TypeOf(out).Elem()
	v := reflect.ValueOf(out).Elem()
	key := ""

LOOP:
	for i := 0; i < typ.NumField(); i++ {
		key = typ.Field(i).Tag.Get("json")
		if len(key) == 0 || key == "-" {
			continue
		}
		if vm, exist := qmap[key]; exist {
			if err = decode(vm, v.Field(i)); err != nil {
				break LOOP
			}
		}
	}

	return err
}

func decode(s []string, v reflect.Value) error {
	var err error = nil

OUT:
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			break
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ui, err := strconv.ParseUint(s[0], 10, 64)
		if err != nil {
			break
		}
		v.SetUint(ui)
	case reflect.String:
		v.SetString(s[0])
	case reflect.Bool:
		b, err := strconv.ParseBool(s[0])
		if err != nil {
			break
		}
		v.SetBool(b)
	case reflect.Slice:
		typ := v.Type()
		v.Set(reflect.MakeSlice(typ, len(s), len(s)))
		switch typ.Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			for i := 0; i < len(s); i++ {
				vi, err := strconv.ParseInt(s[i], 10, 64)
				if err != nil {
					break OUT
				}
				v.Index(i).SetInt(vi)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			for i := 0; i < len(s); i++ {
				vi, err := strconv.ParseUint(s[0], 10, 64)
				if err != nil {
					break OUT
				}
				v.Index(i).SetUint(vi)
			}
		case reflect.String:
			for i := 0; i < len(s); i++ {
				v.Index(i).SetString(s[i])
			}
		default:
			err = fmt.Errorf("unsupported slice type: %d", typ.Elem().Kind())
		}

	default:
		return fmt.Errorf("unsupported type :%d", v.Kind())
	}

	return err
}

func decodeBody(r *http.Request, v ...interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 8192))
	if err != nil {
		return err
	}
	r.Body.Close()

	for i := 0; i < len(v); i++ {
		if err = json.Unmarshal(body, v[i]); err != nil {
			return err
		}
	}

	return nil
}

//==================================== END ======================================
