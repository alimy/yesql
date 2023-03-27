package yesql

import (
	"context"
	"reflect"
)

// Scope just a placeholder type for indicate namespace of object
type Scope any

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
	Prepare(field reflect.Type, query string) (any, error)
	PrepareContext(ctx context.Context, field reflect.Type, query string) (any, error)
}

type QueryFunc func(query *Query) (*Query, error)
