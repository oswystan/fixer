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
	"time"

	"github.com/oswystan/fixer/model"
)

type StoreBugs interface {
	GetBugs() ([]model.Bug, error)
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
	TeamId    int32
	Priority  int32
	Handler   int32
	CreatedBy int32
	Status    int32
	DateFrom  time.Time
	DateTo    time.Time

	Offset uint32
	Count  uint32
}

type storebugs struct {
}

func (b *storebugs) GetBugs(f *BugFilter) ([]model.Bug, error) {
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
	return nil
}

func NewStoreBuglog() StoreBuglog {
	return nil
}

//==================================== END ======================================
