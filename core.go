package yesql

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/jmoiron/sqlx"
)

// Namespace just a placeholder type for indicate namespace of object
type Namespace struct{}

// Query is a parsed query along with tags.
type Query struct {
	Query string
	Tags  map[string]string
}

// QueryMap is a map associating a Tag to its Query
type QueryMap map[string]*Query

// SQLQuery sql query information interface
type SQLQuery interface {
	AddHooks(hooks ...func(query *Query) (*Query, error))
	ListQuery(namespace string) (QueryMap, error)
}

// PrepareScanner scan object interface
type PrepareScanner interface {
	SetPrepareHook(hook PrepareHook)
	Scan(obj any, query SQLQuery) error
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
