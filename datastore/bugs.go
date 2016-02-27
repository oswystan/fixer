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
	"time"

	"github.com/oswystan/fixer/model"
)

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

func (b *storebugs) GetBugs(f *BugFilter) ([]model.Bug, error) {
	if f.TeamId == 0 {
		return nil, fmt.Errorf("need to set a team id")
	}

	// get bug_table and status of the team
	return nil, nil
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
