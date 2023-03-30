package yesql

import (
	"context"
	"database/sql"
	"io"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

const (
	PrepareStyleStmt      = "stmt"
	PrepareStyleNamedStmt = "named_stmt"
	PrepareStyleRaw       = "raw"
	PrepareStyleUnknow    = "unknow"
)

// NamingStrategy naming strategy interface
type NamingStrategy interface {
	FiledNaming(string) string
}

// Namespace just a placeholder type for indicate namespace of object
type Namespace struct{}

type generateOption struct {
	goFileName        string
	defaultStructName string
}

type option interface {
	apply(opt *generateOption)
}

// Query is a parsed query along with tags.
type Query struct {
	Scope string
	Query string
	Tags  map[string]string
}

func (q *Query) PrepareStyle() string {
	prepareStyle := PrepareStyleStmt
	if style, exist := q.Tags["prepare"]; exist {
		style = strings.ToLower(strings.Trim(style, " "))
		switch style {
		case PrepareStyleStmt, PrepareStyleNamedStmt, PrepareStyleRaw, "string":
			prepareStyle = style
		default:
			prepareStyle = PrepareStyleUnknow
		}
	}
	return prepareStyle
}

// QueryMap is a map associating a Tag to its Query
type QueryMap map[string]*Query

func (q QueryMap) FilterByStyle(style string) QueryMap {
	qm := make(QueryMap, len(q))
	for name, query := range q {
		if query.PrepareStyle() == style {
			qm[name] = query
		}
	}
	return qm
}

// ScopeQuery is a namespace QueryMap
type ScopeQuery map[string]QueryMap

// SQLQuery sql query information interface
type SQLQuery interface {
	// SqlQuery get default QueryMap and namespace's QueryMap.
	// return default QueryMap if namespace is empty string
	SqlQuery(namespace string) (QueryMap, QueryMap, error)

	// ListQuery get QuryMap by namespace
	// get default QueryMap if namespace is not give or an empty name
	ListQuery(namespace ...string) (QueryMap, error)

	// ListScope get all namespace Querymap
	ListScope() ScopeQuery

	// AllQuery get all *Query list
	AllQuery() []*Query
}

// SQLParser sql file parser interface
type SQLParser interface {
	AddHooks(hooks ...func(query *Query) (*Query, error))
	ParseReader(reader io.Reader) (SQLQuery, error)
}

// PrepareScanner scan object interface
type PrepareScanner interface {
	SetPrepareHook(hook PrepareHook)
	ScanContext(ctx context.Context, obj any, query SQLQuery) error
}

// PrepareHook prepare hook for scan object
type PrepareHook interface {
	PrepareContext(ctx context.Context, field reflect.Type, query string) (any, error)
}

// PrepareContext enhances the Conn interface with context.
type PrepareContext interface {
	// PreparexContext returns a prepared statement, bound to this connection.
	// context is for the preparation of the statement,
	// it must not store the context within the statement itself.
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// PreparexContext enhances the Conn interface with context.
type PreparexContext interface {
	// PrepareContext prepares a statement.
	// The provided context is used for the preparation of the statement, not for
	// the execution of the statement.
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)

	// PrepareNamedContext returns an sqlx.NamedStmt
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
}

// Generator generate struct code automatic base SQLQuery
type Generator interface {
	Generate(dstPath string, pkgName string, query SQLQuery, opts ...option) error
}

type OptionFunc func(opt *generateOption)

func (f OptionFunc) apply(opt *generateOption) {
	f(opt)
}

func DefaultStructNameOpt(name string) OptionFunc {
	return func(opt *generateOption) {
		opt.defaultStructName = name
	}
}

func GoFileNameOpt(name string) OptionFunc {
	return func(opt *generateOption) {
		opt.goFileName = name
	}
}
