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

var sqlBuglist = "select b.*, get_nicky(b.current_handler) as handler_nicky, get_nicky(b.created_by) as created_nicky from %s as b"

var sqlDelBugs = `delete from %s`
var sqlDelBug = `delete from %s where id=?`
var sqlGetBug = `select * from %s where id=?`
var sqlCreateBug = `insert into %s(created_by, current_handler,	priority, status, 
                   created_time, last_update, title, attachments, detail) values
				   (?created_by, ?current_handler, ?priority, ?status, now(),
			       now(), ?title, ?attachments, ?detail) returning *`
var sqlUpdateBug = `update %s set current_handler=?current_handler, priority=?priority, status=?status, 
			       last_update=?last_update, title=?title, attachments=?attachments, detail=?detail 
				   where id=?id returning *`

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

type BugStore interface {
	SetTeam(tid int) error
	GetBugs(*model.Filter) ([]model.Bug, error)
	DeleteBugs() error
	GetBug(id int) (*model.Bug, error)
	Create(*model.Bug) (*model.Bug, error)
	Update(*model.Bug) (*model.Bug, error)
	Delete(int) error

	GetLogs(bugId int, f *model.Filter) ([]model.Buglog, error)
	CreateLog(bl *model.Buglog) (*model.Buglog, error)
	DeleteLog(bl *model.Buglog) error
}

type bugstore struct {
	team   *model.Team
	bugtab string
}

func (ds *bugstore) buildSql(f *model.Filter) (string, error) {
	sql := fmt.Sprintf(sqlBuglist, ds.bugtab)
	str := ""
	if len(f.Handler) != 0 {
	}
	if len(f.CreatedBy) != 0 {
		//str = fmt.Sprintf("%s and created_by = %d", str, f.CreatedBy)
	}
	if len(f.Status) != 0 {
		//str = fmt.Sprintf("%s and status = %d", str, f.Status)
	}
	if len(f.DateFrom) != 0 {
		//str = fmt.Sprintf("%s and created_time >= '%s'", str, f.DateFrom.Format("2006-01-02 15:04:05"))
	}
	if len(f.DateTo) != 0 {
		//str = fmt.Sprintf("%s and created_time <= '%s'", str, f.DateTo.Format("2006-01-02 15:04:05"))
	}
	if f.Limit != 0 {
		str = fmt.Sprintf("%s limit %d", str, f.Limit)
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

func (ds *bugstore) SetTeam(tid int) error {
	ts := NewTeamStore()
	t, err := ts.GetTeam(tid)
	if err != nil {
		return err
	}
	if t.Status == 0 || len(t.BugTable) == 0 || t.BugTableStatus != 1 {
		return fmt.Errorf("invalid status(%d,%s,%d)", t.Status, t.BugTable, t.BugTableStatus)
	}

	ds.bugtab = fmt.Sprintf("bugs_%s", strings.TrimSpace(t.BugTable))
	ds.team = t
	return nil
}
func (ds *bugstore) GetBugs(f *model.Filter) ([]model.Bug, error) {
	if ds.team == nil {
		return nil, fmt.Errorf("should set tid first")
	}
	sql, err := ds.buildSql(f)
	if err != nil {
		return nil, err
	}

	var bugs []model.Bug
	_, err = GetDB().pg.Query(&bugs, sql)
	return bugs, err
}

func (ds *bugstore) DeleteBugs() error {
	if ds.team == nil {
		return fmt.Errorf("should set tid first")
	}
	sql := fmt.Sprintf(sqlDelBugs, ds.bugtab)
	_, err := GetDB().pg.Exec(sql)
	return err
}

func (ds *bugstore) GetBug(id int) (*model.Bug, error) {
	if ds.team == nil {
		return nil, fmt.Errorf("should set tid first")
	}
	bug := &model.Bug{}
	sql := fmt.Sprintf(sqlGetBug, ds.bugtab)
	_, err := GetDB().pg.QueryOne(bug, sql, id)
	return bug, err
}
func (ds *bugstore) Create(b *model.Bug) (*model.Bug, error) {
	if ds.team == nil {
		return nil, fmt.Errorf("should set tid first")
	}
	sql := fmt.Sprintf(sqlCreateBug, ds.bugtab)
	_, err := GetDB().pg.QueryOne(b, sql, b)
	return b, err
}
func (ds *bugstore) Update(b *model.Bug) (*model.Bug, error) {
	if ds.team == nil {
		return nil, fmt.Errorf("should set tid first")
	}
	sql := fmt.Sprintf(sqlUpdateBug, ds.bugtab)
	_, err := GetDB().pg.QueryOne(b, sql, b)
	return b, err
}
func (ds *bugstore) Delete(id int) error {
	if ds.team == nil {
		return fmt.Errorf("should set tid first")
	}
	sql := fmt.Sprintf(sqlDelBug, ds.bugtab)
	_, err := GetDB().pg.Exec(sql, id)
	return err
}

func (ds *bugstore) GetLogs(bugId int, f *model.Filter) ([]model.Buglog, error) {
	return nil, nil
}
func (ds *bugstore) CreateLog(bl *model.Buglog) (*model.Buglog, error) {
	return nil, nil
}
func (ds *bugstore) DeleteLog(bl *model.Buglog) error {
	return nil
}

func NewBugStore() BugStore {
	return &bugstore{team: nil}
}

//==================================== END ======================================
