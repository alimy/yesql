package template

import (
	_ "embed"
	"html/template"
)

//go:embed sql.tmpl
var sqlTmplStr string

//go:embed sqlx.tmpl
var sqlxTmplStr string

func NewSqlTemplate() (*template.Template, error) {
	return template.New("sql").Parse(sqlTmplStr)
}

func NewSqlxTemplate() (*template.Template, error) {
	return template.New("sqlx").Parse(sqlxTmplStr)
}
