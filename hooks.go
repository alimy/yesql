package yesql

import (
	"context"
	"fmt"
	"reflect"
)

var (
	_ PrepareHook = (*stdPrepareHook)(nil)
	_ PrepareHook = (*sqlxPrepareHook)(nil)
)

type stdPrepareHook struct {
	db PrepareContext
}

type sqlxPrepareHook struct {
	db PreparexContext
}

func (s *stdPrepareHook) Prepare(field reflect.Type, query string) (any, error) {
	switch field.String() {
	case "string":
		// Unprepared SQL query.
		return query, nil
	case "*sql.Stmt":
		// Prepared query.
		stmt, err := s.db.PrepareContext(context.Background(), query)
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
		stmt, err := s.db.PrepareContext(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	default:
		return nil, fmt.Errorf("not support filed type '%s': %v", query, field)
	}
}

func (s *sqlxPrepareHook) Prepare(field reflect.Type, query string) (any, error) {
	switch field.String() {
	case "string":
		// Unprepared SQL query.
		return query, nil
	case "*sqlx.Stmt":
		// Prepared query.
		stmt, err := s.db.PreparexContext(context.Background(), query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	case "*sqlx.NamedStmt":
		// Prepared query.
		stmt, err := s.db.PrepareNamedContext(context.Background(), query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	default:
		return nil, fmt.Errorf("not support filed type '%s': %v", query, field)
	}
}

func (s *sqlxPrepareHook) PrepareContext(ctx context.Context, field reflect.Type, query string) (any, error) {
	switch field.String() {
	case "string":
		// Unprepared SQL query.
		return query, nil
	case "*sqlx.Stmt":
		// Prepared query.
		stmt, err := s.db.PreparexContext(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	case "*sqlx.NamedStmt":
		// Prepared query.
		stmt, err := s.db.PrepareNamedContext(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	default:
		return nil, fmt.Errorf("not support filed type '%s': %v", query, field)
	}
}

// NewPrepareHook create a prepare hook with *sql.DB
func NewPrepareHook(db PrepareContext) PrepareHook {
	return &stdPrepareHook{
		db: db,
	}
}

// NewSqlxPrepareHook create a prepare hook with *sqlx.DB
func NewSqlxPrepareHook(db PreparexContext) PrepareHook {
	return &sqlxPrepareHook{
		db: db,
	}
}
