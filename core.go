package yesql

import (
	"context"
	"database/sql"
	"io"
	"reflect"
	"strings"
)

const (
	PrepareStyleStmt      = "stmt"
	PrepareStyleNamedStmt = "named_stmt"
	PrepareStyleRaw       = "raw"
	PrepareStyleUnknow    = "unknow"
)

// Namespace just a placeholder type for indicate namespace of object
type Namespace struct{}

type generateOption struct {
	goFileName        string
	defaultStructName string
	sqlxPkgName       string
}

type option interface {
	apply(opt *generateOption)
}

// Query is a parsed query along with tags.
type Query struct {
	Scope string
	Name  string
	Query string
	Tags  map[string]string
}

// QueryList query list
type QueryList []*Query

// QueryMap is a map associating a Tag to its Query
type QueryMap map[string]*Query

func (q *Query) PrepareStyle() string {
	prepareStyle := PrepareStyleStmt
	if style, exist := q.Tags["prepare"]; exist {
		style = strings.ToLower(strings.Trim(style, " "))
		switch style {
		case PrepareStyleStmt, PrepareStyleNamedStmt, PrepareStyleRaw:
			prepareStyle = style
		default:
			prepareStyle = PrepareStyleRaw
		}
	}
	return prepareStyle
}

func (q QueryList) Len() int {
	return len(q)
}

func (q QueryList) Less(i, j int) bool {
	return q[i].Scope+"_"+q[i].Name < q[j].Scope+"_"+q[j].Name
}

func (q QueryList) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q QueryMap) FilterByStyle(style string) QueryMap {
	qm := make(QueryMap, len(q))
	for name, query := range q {
		if query.PrepareStyle() == style {
			qm[name] = query
		}
	}
	return qm
}

func (q QueryMap) IsNotEmpty() bool {
	return len(q) > 0
}

func (q QueryMap) IsRawQueryNotEmpty() bool {
	for _, query := range q {
		if query.PrepareStyle() == PrepareStyleRaw {
			return true
		}
	}
	return false
}

func (q QueryMap) IsStmtQueryNotEmpty() bool {
	for _, query := range q {
		switch query.PrepareStyle() {
		case PrepareStyleStmt, PrepareStyleNamedStmt:
			return true
		default:
			// nothing
		}
	}
	return false
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
	AllQuery() QueryList
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

// PreparexContext[T, S] enhances the Conn interface with context.
type PreparexContext[T, S any] interface {
	// PrepareContext prepares a statement.
	// The provided context is used for the preparation of the statement, not for
	// the execution of the statement.
	PreparexContext(ctx context.Context, query string) (T, error)

	// PrepareNamedContext returns an sqlx.NamedStmt
	PrepareNamedContext(ctx context.Context, query string) (S, error)

	// Rebind rebind query to adapte SQL Driver
	Rebind(query string) string
}

// PrepareBuilder prepare builder interface sql
type PrepareBuilder interface {
	PrepareContext
	QueryHook(query string) string
}

// PreparexBuilder[T, S] preparex builder interface for sqlx
type PreparexBuilder[T, S any] interface {
	PreparexContext[T, S]
	QueryHook(query string) string
}

// Generator generate struct code automatic base SQLQuery
type Generator interface {
	Generate(dstPath string, pkgName string, query SQLQuery, opts ...option) error
}

// OptFn option function
type OptFn func(opt *generateOption)

func (f OptFn) apply(opt *generateOption) {
	f(opt)
}

// DefaultStructNameOpt set custom default global sql query struct name
func DefaultStructNameOpt(name string) OptFn {
	return func(opt *generateOption) {
		opt.defaultStructName = name
	}
}

// GoFileNameOpt set custom go file name to generate
func GoFileNameOpt(name string) OptFn {
	return func(opt *generateOption) {
		opt.goFileName = name
	}
}

// SqlxPkgNameOpt set sqlx pakcage name
func SqlxPkgNameOpt(name string) OptFn {
	return func(opt *generateOption) {
		opt.sqlxPkgName = name
	}
}
