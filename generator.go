package yesql

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	stdTmpl "text/template"

	"github.com/alimy/yesql/naming"
	"github.com/alimy/yesql/template"
	"github.com/jmoiron/sqlx"
)

var (
	_ Generator = (*sqlGenerator)(nil)
)

type tmplCtx struct {
	PkgName           string
	DefaultStructName string
	AllQuery          []*Query
	DefaultQueryMap   QueryMap
	ScopeQuery        ScopeQuery
	YesqlVer          string
}

type simplePrepareBuilder struct {
	p    PrepareContext
	hook func(string) string
}

type simplePreparexBuilder struct {
	p    PreparexContext
	hook func(string) string
}

type sqlGenerator struct {
	tmpl *stdTmpl.Template
}

func (s *simplePrepareBuilder) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return s.p.PrepareContext(ctx, query)
}

func (s *simplePrepareBuilder) QueryHook(query string) string {
	if s.hook != nil {
		return s.hook(query)
	}
	return query
}

func (s *simplePreparexBuilder) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return s.p.PreparexContext(ctx, query)
}

func (s *simplePreparexBuilder) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	return s.p.PrepareNamedContext(ctx, query)
}

func (s *simplePreparexBuilder) Rebind(query string) string {
	return s.p.Rebind(query)
}

func (s *simplePreparexBuilder) QueryHook(query string) string {
	if s.hook != nil {
		return s.hook(query)
	}
	return query
}

func (t *tmplCtx) DefaultQueryMapNotEmpty() bool {
	return len(t.DefaultQueryMap) != 0
}

func (s *sqlGenerator) Generate(dstPath string, pkgName string, query SQLQuery, opts ...option) (err error) {
	opt := &generateOption{
		goFileName:        "yesql.go",
		defaultStructName: "Yesql",
	}
	for _, arg := range opts {
		arg.apply(opt)
	}
	data := &tmplCtx{
		PkgName:           pkgName,
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
func NewPrepareBuilder(p PrepareContext, hook ...func(string) string) PrepareBuilder {
	obj := &simplePrepareBuilder{
		p: p,
	}
	if len(hook) > 0 && hook[0] != nil {
		obj.hook = hook[0]
	}
	return obj
}

// NewPreprarexBuilder create a simple preparex builder instance
func NewPreparexBuilder(p PreparexContext, hook ...func(string) string) PreparexBuilder {
	obj := &simplePreparexBuilder{
		p: p,
	}
	if len(hook) > 0 && hook[0] != nil {
		obj.hook = hook[0]
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
