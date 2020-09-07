package utils

import (
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// 密码登录
func SignInWithPassword(tag, val, password, table, key string) (gdb.Record, Error) {
	e := NewErr()

	user, err := g.DB().Table(table).FindOne(tag, val)
	e.Append(err)

	if user.IsEmpty() {
		e.Append(errors.New("用户名或密码错误！"))
		return nil, e
	}

	if !PBKDF2Decode(password, user["password"].String(), 1024, func() string {
		return key
	}) {
		e.Append(errors.New("用户名或密码错误！"))
		return nil, e
	}

	return user, e
}

