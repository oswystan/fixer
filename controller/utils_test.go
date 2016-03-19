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

import (
	"testing"

	"github.com/oswystan/fixer/model"
)

func TestDecodeQuery(t *testing.T) {
	qstr := "offset=1&limit=10&q=john&handler=1&handler=2&handler=3&fields=a&fields=b&date_from=20110101"
	f := &model.Filter{}
	err := decodeQuery(qstr, f)
	if err != nil {
		t.Fatalf("fail to decode %s, error=%s", qstr, err)
	}
}

//==================================== END ======================================
