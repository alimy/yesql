package yesql

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
)

var (
	_ PrepareHook = (*stdPrepareHook)(nil)
	_ PrepareHook = (*sqlxPrepareHook)(nil)
)

type stdPrepareHook struct {
	db *sql.DB
}

type sqlxPrepareHook struct {
	db *sqlx.DB
}

func (s *stdPrepareHook) Prepare(field reflect.Type, query string) (any, error) {
	switch field.String() {
	case "string":
		// Unprepared SQL query.
		return query, nil
	case "*sql.Stmt":
		// Prepared query.
		stmt, err := s.db.Prepare(query)
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
		stmt, err := s.db.Preparex(query)
		if err != nil {
			return nil, fmt.Errorf("Error preparing query '%s': %v", query, err)
		}
		return stmt, nil
	case "*sqlx.NamedStmt":
		// Prepared query.
		stmt, err := s.db.PrepareNamed(query)
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
func NewPrepareHook(db *sql.DB) PrepareHook {
	return &stdPrepareHook{
		db: db,
	}
}

// NewSqlxPrepareHook create a prepare hook with *sqlx.DB
func NewSqlxPrepareHook(db *sqlx.DB) PrepareHook {
	return &sqlxPrepareHook{
		db: db,
	}
}
