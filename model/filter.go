//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: filter.go
//     description:
//         created: 2016-02-29 14:57:59
//          author: wystan
//
//===============================================================================

package model

type Filter struct {
	Offset    int      `json:"offset"`
	Limit     int      `json:"limit"`
	UserId    int      `json:"-"`
	TeamId    int      `json:"team_id"`
	Q         string   `json:"q"`
	Fields    []string `json:"fields"`
	Handler   []int    `json:"handler"`
	Priority  []int    `json:"priority"`
	CreatedBy []int    `json:"created_by"`
	Status    []int    `json:"status"`
	DateFrom  string   `json:"date_from"`
	DateTo    string   `json:"date_to"`
}

func NewFilter() *Filter {
	return &Filter{Offset: 0, Limit: 10}
}

//==================================== END ======================================
