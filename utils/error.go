package utils

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gdb"
)

type Error interface {
	Append(err ...error)
	IsEmpty() bool
	HandleEmptyRecord(rec gdb.Record, tag string)
	Errs() []error
	ToString() string
}

type err struct {
	errs []error
}

func NewErr() Error {
	return &err{make([]error, 0)}
}

func (e *err) Append(err ...error) {
	for _, v := range err {
		if v != nil {
			e.errs = append(e.errs, v)
		}
	}
}

func (e *err) IsEmpty() bool {
	if len(e.errs) == 0 {
		return true
	}

	return false
}

func (e *err) HandleEmptyRecord(rec gdb.Record, tag string) {
	if rec.IsEmpty() {
		e.Append(errors.New(fmt.Sprintf("数据不存在：[%s]", tag)))
	}
}

func (e *err) Errs() []error {
	return e.errs
}

func (e *err) ToString() string {
	s := ""
	for _, v := range e.errs {
		s += fmt.Sprintf("%s\n", v.Error())
	}
	return s
}
