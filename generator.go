package yesql

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	stdTmpl "text/template"

	"github.com/alimy/yesql/naming"
	"github.com/alimy/yesql/template"
)

const (
	_defaultSqlxPkgName = "github.com/jmoiron/sqlx"
)

var (
	_ Generator = (*sqlGenerator)(nil)
)

type tmplCtx struct {
	PkgName           string
	SqlxPkgName       string
	DefaultStructName string
	AllQuery          []*Query
	DefaultQueryMap   QueryMap
	ScopeQuery        ScopeQuery
	YesqlVer          string
}

type simplePrepareBuilder struct {
	p     PrepareContext
	hooks []func(string) string
}

type simplePreparexBuilder[T, S any] struct {
	p     PreparexContext[T, S]
	hooks []func(string) string
}

type sqlGenerator struct {
	tmpl *stdTmpl.Template
}

func (s *simplePrepareBuilder) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return s.p.PrepareContext(ctx, query)
}

func (s *simplePrepareBuilder) QueryHook(query string) string {
	for _, h := range s.hooks {
		query = h(query)
	}
	return query
}

func (s *simplePreparexBuilder[T, S]) PreparexContext(ctx context.Context, query string) (T, error) {
	return s.p.PreparexContext(ctx, query)
}

func (s *simplePreparexBuilder[T, S]) PrepareNamedContext(ctx context.Context, query string) (S, error) {
	return s.p.PrepareNamedContext(ctx, query)
}

func (s *simplePreparexBuilder[T, S]) Rebind(query string) string {
	return s.p.Rebind(query)
}

func (s *simplePreparexBuilder[T, S]) QueryHook(query string) string {
	for _, h := range s.hooks {
		query = h(query)
	}
	return query
}

func (t *tmplCtx) DefaultQueryMapNotEmpty() bool {
	return len(t.DefaultQueryMap) != 0
}

func (t *tmplCtx) ScopeQueryNotEmpty() bool {
	return len(t.ScopeQuery) != 0
}

func (s *sqlGenerator) Generate(dstPath string, pkgName string, query SQLQuery, opts ...option) (err error) {
	opt := &generateOption{
		goFileName:        "yesql.go",
		defaultStructName: "Yesql",
		sqlxPkgName:       _defaultSqlxPkgName,
	}
	for _, arg := range opts {
		arg.apply(opt)
	}
	data := &tmplCtx{
		PkgName:           pkgName,
		SqlxPkgName:       opt.sqlxPkgName,
		DefaultStructName: naming.Naming(opt.defaultStructName),
		AllQuery:          query.AllQuery(),
		ScopeQuery:        query.ListScope(),
		YesqlVer:          Version,
	}
	if len(data.PkgName) == 0 {
		data.PkgName = "yesql"
	}
	data.DefaultQueryMap, err = query.ListQuery()
	if err != nil {
		return err
	}

	if filepath.Ext(opt.goFileName) != ".go" {
		opt.goFileName += ".go"
	}
	filePath := filepath.Join(dstPath, opt.goFileName)
	dirPath := filepath.Dir(filePath)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return s.tmpl.Execute(file, data)
}

// NewPrepareBuilder create a simple prepare builder instance
func NewPrepareBuilder(p PrepareContext, hooks ...func(string) string) PrepareBuilder {
	obj := &simplePrepareBuilder{
		p: p,
	}
	for _, h := range hooks {
		if h != nil {
			obj.hooks = append(obj.hooks, h)
		}
	}
	return obj
}

// NewPreprarexBuilder[T, S] create a simple preparex builder instance
func NewPreparexBuilder[T, S any](p PreparexContext[T, S], hooks ...func(string) string) PreparexBuilder[T, S] {
	obj := &simplePreparexBuilder[T, S]{
		p: p,
	}
	for _, h := range hooks {
		if h != nil {
			obj.hooks = append(obj.hooks, h)
		}
	}
	return obj
}

// NewSqlGenerator create a sql generator use std sql
func NewSqlGenerator() Generator {
	return &sqlGenerator{
		tmpl: template.NewSqlTemplate(),
	}
}

// NewSqlxGenerator create a sqlx generator use sqlx
func NewSqlxGenerator() Generator {
	return &sqlGenerator{
		tmpl: template.NewSqlxTemplate(),
	}
}
