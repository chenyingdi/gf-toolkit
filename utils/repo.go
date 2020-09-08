package utils

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type REPO interface {
	Find(q Query) (Page, error)
	Insert(p g.Map) (int64, error)
	Update(where g.Map, p g.Map) error
	Delete(where g.Map) error
}

type repo struct {
	table string
}

type Query struct {
	Where     g.Map   `json:"where"`
	And       g.Map   `json:"and"`
	Or        g.Map   `json:"or"`
	LeftJoin  []g.Map `json:"left_join"`
	RightJoin []g.Map `json:"right_join"`
	InnerJoin []g.Map `json:"inner_join"`
	Fields    string  `json:"fields"`
	FieldsEx  string  `json:"fields_ex"`
	Group     string  `json:"group"`
	Order     string  `json:"order"`
	Having    string  `json:"having"`
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
}

type Page struct {
	Results gdb.Result `json:"results"`
	Total   int        `json:"total"`
}

func NewRepo(table string) REPO {
	return &repo{
		table: table,
	}
}

func (r *repo) Find(q Query) (Page, error) {
	var (
		err   error
		count int
		res   gdb.Result
	)

	query := g.DB().Table(r.table)

	// left join
	for _, v := range q.LeftJoin {
		query = query.LeftJoin(gconv.String(v["table"]), gconv.String(v["on"]))
	}

	//  right join
	for _, v := range q.RightJoin {
		query = query.RightJoin(gconv.String(v["table"]), gconv.String(v["on"]))
	}

	//  inner join
	for _, v := range q.InnerJoin {
		query = query.InnerJoin(gconv.String(v["table"]), gconv.String(v["on"]))
	}

	// where and or
	query = query.Where(q.Where).And(q.And).Or(q.Or)

	if q.Fields != "" {
		query = query.Fields(q.Fields)
	}

	if q.FieldsEx != "" {
		query = query.FieldsEx(q.FieldsEx)
	}

	if q.Having != "" {
		query = query.Having(q.Having)
	}

	if q.Group != "" {
		query = query.Group(q.Group)
	}

	// 统计总数
	count, err = query.Clone().FindCount()

	// order page limit
	query = query.Order(q.Order).Page(q.Page, q.Limit)

	res, err = query.FindAll()

	return Page{
		Results: res,
		Total:   count,
	}, err
}

func (r *repo) Insert(p g.Map) (int64, error) {
	res, err := g.DB().Table(r.table).Insert(p)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *repo) Update(where g.Map, p g.Map) error {
	_, err := g.DB().Table(r.table).Where(where).Data(p).Update()
	return err
}

func (r *repo) Delete(where g.Map) error {
	_, err := g.DB().Table(r.table).Delete(where)
	return err
}
