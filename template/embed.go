package template

import (
	_ "embed"
	"html/template"

	"github.com/alimy/yesql/naming"
)

//go:embed sql.tmpl
var sqlTmplStr string

//go:embed sqlx.tmpl
var sqlxTmplStr string

func NewSqlTemplate() (*template.Template, error) {
	return newTemplate("sql").Parse(sqlTmplStr)
}

func NewSqlxTemplate() (*template.Template, error) {
	return newTemplate("sqlx").Parse(sqlxTmplStr)
}

func newTemplate(name string) *template.Template {
	return template.New(name).Funcs(template.FuncMap{
		"naming":      naming.Naming,
		"notEmptyStr": notEmptyStr,
	})
}

func notEmptyStr(s string) bool {
	return len(s) > 0
}
