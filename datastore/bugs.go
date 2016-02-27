//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: bugs.go
//     description:
//         created: 2016-02-26 22:13:49
//          author: wystan
//
//===============================================================================
package datastore

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oswystan/fixer/model"
)

var sqlBuglist = "select b.*, get_nicky(b.current_handler) as handler_nicky, get_nicky(b.created_by) as created_nicky from bugs_%s as b"

type StoreBugs interface {
	GetBugs(*BugFilter) ([]model.Bug, error)
	GetBugById(id int32) (*model.Bug, error)
	Create(*model.Bug) error
	Update(*model.Bug) error
	Delete(*model.Bug) error
}

type StoreBuglog interface {
	GetLogs(id int32) ([]model.Buglog, error)
	Create(*model.Buglog) error
	Delete(*model.Buglog) error
}

type BugFilter struct {
	TeamId    int
	Priority  int
	Handler   int
	CreatedBy int
	Status    int
	DateFrom  time.Time
	DateTo    time.Time

	Offset int
	Count  int
}

type storebugs struct {
	t StoreTeamList
}

func (b *storebugs) buildSql(f *BugFilter) (string, error) {
	// get bug_table and status of the team
	t, err := b.t.GetTeamById(f.TeamId)
	if err != nil {
		return "", err
	}
	if t.Status != 1 || t.BugTable == "" {
		return "", fmt.Errorf("team status invalid [%d]", t.Status)
	}

	sql := fmt.Sprintf(sqlBuglist, strings.TrimSpace(t.BugTable))
	str := ""
	if f.Handler != 0 {
		str = fmt.Sprintf("%s and current_handler = %d", str, f.Handler)
	}
	if f.CreatedBy != 0 {
		str = fmt.Sprintf("%s and created_by = %d", str, f.CreatedBy)
	}
	if f.Status != 0 {
		str = fmt.Sprintf("%s and status = %d", str, f.Status)
	}
	if !f.DateFrom.IsZero() {
		str = fmt.Sprintf("%s and created_time >= '%s'", str, f.DateFrom.Format("2006-01-02 15:04:05"))
	}
	if !f.DateTo.IsZero() {
		str = fmt.Sprintf("%s and created_time <= '%s'", str, f.DateTo.Format("2006-01-02 15:04:05"))
	}
	if f.Count != 0 {
		str = fmt.Sprintf("%s limit %d", str, f.Count)
	}
	if f.Offset != 0 {
		str = fmt.Sprintf("%s offset %d", str, f.Offset)
	}

	if len(str) != 0 {
		str = strings.Replace(str, "and", "where", 1)
		sql += str
	}
	log.Println(sql)

	return sql, nil
}

func (b *storebugs) GetBugs(f *BugFilter) ([]model.Bug, error) {
	if f.TeamId == 0 {
		return nil, fmt.Errorf("need to set a team id")
	}
	sql, err := b.buildSql(f)
	if err != nil {
		return nil, err
	}

	var bl []model.Bug
	_, err = GetDB().pg.Query(&bl, sql)
	if err != nil {
		return nil, err
	}

	return bl, nil
}

func (b *storebugs) GetBugById(id int32) (*model.Bug, error) {
	return nil, nil
}

func (b *storebugs) Create(*model.Bug) error {
	return nil
}

func (b *storebugs) Update(*model.Bug) error {
	return nil
}

func (b *storebugs) Delete(*model.Bug) error {
	return nil
}

func NewStoreBugs() StoreBugs {
	return &storebugs{t: NewStoreTeamList()}
}

func NewStoreBuglog() StoreBuglog {
	return nil
}

//==================================== END ======================================
