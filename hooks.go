package yesql

import (
	"context"
	"fmt"
	"reflect"
)

var (
	_ PrepareHook = (*stdPrepareHook)(nil)
	_ PrepareHook = (*sqlxPrepareHook[any, any])(nil)
)

type stdPrepareHook struct {
	prepare PrepareContext
}

type sqlxPrepareHook[T, S any] struct {
	prepare PreparexContext[T, S]
}

func (s *stdPrepareHook) Prepare(field reflect.Type, query string) (any, error) {
	switch field.String() {
	case "string":
		// Unprepared SQL query.
		return query, nil
	case "*sql.Stmt":
		// Prepared query.
		stmt, err := s.prepare.PrepareContext(context.Background(), query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	default:
		return nil, fmt.Errorf("not support filed type: %v", field)
	}
}

func (s *stdPrepareHook) PrepareContext(ctx context.Context, field reflect.Type, query string) (any, error) {
	switch field.String() {
	case "string":
		// Unprepared SQL query.
		return query, nil
	case "*sql.Stmt":
		// Prepared query.
		stmt, err := s.prepare.PrepareContext(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	default:
		return nil, fmt.Errorf("not support filed type '%s': %v", query, field)
	}
}

func (s *sqlxPrepareHook[T, S]) Prepare(field reflect.Type, query string) (any, error) {
	switch field.String() {
	case "string":
		// Unprepared SQL query.
		return query, nil
	case "*sqlx.Stmt":
		// Prepared query.
		stmt, err := s.prepare.PreparexContext(context.Background(), query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	case "*sqlx.NamedStmt":
		// Prepared query.
		stmt, err := s.prepare.PrepareNamedContext(context.Background(), query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	default:
		return nil, fmt.Errorf("not support filed type '%s': %v", query, field)
	}
}

func (s *sqlxPrepareHook[T, S]) PrepareContext(ctx context.Context, field reflect.Type, query string) (any, error) {
	switch field.String() {
	case "string":
		// Unprepared SQL query.
		return query, nil
	case "*sqlx.Stmt":
		// Prepared query.
		stmt, err := s.prepare.PreparexContext(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	case "*sqlx.NamedStmt":
		// Prepared query.
		stmt, err := s.prepare.PrepareNamedContext(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	default:
		return nil, fmt.Errorf("not support filed type '%s': %v", query, field)
	}
}

// NewPrepareHook create a prepare hook with prepare that implement PrepareContext
func NewPrepareHook(p PrepareContext) PrepareHook {
	return &stdPrepareHook{
		prepare: p,
	}
}

// NewSqlxPrepareHook[T] create a prepare hook prepare that implement PreparexContext
func NewSqlxPrepareHook[T, S any](p PreparexContext[T, S]) PrepareHook {
	return &sqlxPrepareHook[T, S]{
		prepare: p,
	}
}
