//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: utils.go
//     description:
//         created: 2016-02-29 20:05:28
//          author: wystan
//
//===============================================================================

package model

type FieldsFilter map[string]uint8

func JsonFilter(v interface{}, fields FieldsFilter) ([]byte, error) {
	//if fields == nil {
	//    return json.Marshal(v)
	//}

	//typ := reflect.TypeOf(v)
	//if typ.Kind() == reflect.Ptr {
	//    typ := reflect.TypeOf(v).Elem()
	//}

	//for i := 0; i < typ.NumField(); i++ {
	//    _, exists := fields[typ.Field(i).Tag.Get("json")]
	//}

	return nil, nil
}

//==================================== END ======================================
