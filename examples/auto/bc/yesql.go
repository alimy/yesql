// Code generated by Yesql. DO NOT EDIT.
// versions:
// - Yesql v1.8.4

package bc

import (
	"context"

	"github.com/bitbus/sqlx"
)

const (
	_LoginInfoA  = `SELECT * FROM @user WHERE username=?`
	_LoginInfoB  = `SELECT * FROM @user WHERE username=?`
	_LoginInfoD  = `SELECT * FROM @user WHERE username=?`
	_LoginInfoE  = `SELECT * FROM @user WHERE username=?`
	_LogoutInfoC = `SELECT * FROM @user WHERE username=?`
	_UserInfoA   = `SELECT ` + "`" + `username` + "`" + ` FROM @user WHERE username=?`
	_UserInfoB   = `SELECT ` + "`" + `username` + "`" + `, ` + "`" + `nickname` + "`" + ` FROM @user WHERE username=?`
	_UserInfoC   = `SELECT "username" FROM @user WHERE username=?`
	_UserInfoD   = `SELECT "username", "nickname" FROM @user WHERE username=?`
)

// PreparexContext enhances the Conn interface with context.
type PreparexContext interface {
	// PrepareContext prepares a statement.
	// The provided context is used for the preparation of the statement, not for
	// the execution of the statement.
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)

	// PrepareNamedContext returns an sqlx.NamedStmt
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)

	// Rebind rebind query to adapte SQL Driver
	Rebind(query string) string
}

// PreparexBuilder preparex builder interface for sqlx
type PreparexBuilder interface {
	PreparexContext
	QueryHook(query string) string
}

type Yesql struct {
	LoginInfoD  string          `yesql:"login_info_d"`
	LogoutInfoC string          `yesql:"logout_info_c"`
	UserInfoC   string          `yesql:"user_info_c"`
	UserInfoD   string          `yesql:"user_info_d"`
	LoginInfoB  *sqlx.Stmt      `yesql:"login_info_b"`
	UserInfoA   *sqlx.Stmt      `yesql:"user_info_a"`
	UserInfoB   *sqlx.Stmt      `yesql:"user_info_b"`
	LoginInfoA  *sqlx.NamedStmt `yesql:"login_info_a"`
	LoginInfoE  *sqlx.NamedStmt `yesql:"login_info_e"`
}

func BuildYesql(p PreparexBuilder, ctx ...context.Context) (obj *Yesql, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Yesql{
		LoginInfoD:  p.QueryHook(_LoginInfoD),
		LogoutInfoC: p.QueryHook(_LogoutInfoC),
		UserInfoC:   p.QueryHook(_UserInfoC),
		UserInfoD:   p.QueryHook(_UserInfoD),
	}
	if obj.LoginInfoB, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_LoginInfoB))); err != nil {
		return
	}
	if obj.UserInfoA, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_UserInfoA))); err != nil {
		return
	}
	if obj.UserInfoB, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_UserInfoB))); err != nil {
		return
	}
	if obj.LoginInfoA, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_LoginInfoA))); err != nil {
		return
	}
	if obj.LoginInfoE, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_LoginInfoE))); err != nil {
		return
	}
	return
}