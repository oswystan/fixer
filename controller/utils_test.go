//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: utils_test.go
//     description:
//         created: 2016-02-29 11:57:37
//          author: wystan
//
//===============================================================================

package controller

import "testing"

func TestDecodeQuery(t *testing.T) {
	qstr := "offset=1&limit=10&q=john"
	type Filter struct {
		Offset int    `json:"offset"`
		Limit  int    `json:"limit"`
		Q      string `json:"q"`
	}

	f := &Filter{}
	err := decodeQuery(qstr, f)
	if err != nil {
		t.Errorf("fail to decode %s", qstr)
		return
	}
}

//==================================== END ======================================
