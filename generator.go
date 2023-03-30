package yesql

import (
	stdTmpl "html/template"
	"os"
	"path/filepath"

	"github.com/alimy/yesql/template"
)

var (
	_ Generator = (*sqlGenerator)(nil)
)

type tmplCtx struct {
	PkgName           string
	DefaultStructName string
	DefaultQueryMap   QueryMap
	ScopeQuery        ScopeQuery
	YesqlVer          string
}

type sqlGenerator struct {
	tmpl *stdTmpl.Template
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
		DefaultStructName: _defaultNamingStrategy.FiledNaming(opt.defaultStructName),
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

func NewSqlGenerator() (Generator, error) {
	tmpl, err := template.NewSqlTemplate()
	if err != nil {
		return nil, err
	}
	return &sqlGenerator{tmpl: tmpl}, nil
}

func NewSqlxGenerator() (Generator, error) {
	tmpl, err := template.NewSqlxTemplate()
	if err != nil {
		return nil, err
	}
	return &sqlGenerator{tmpl: tmpl}, nil
}
