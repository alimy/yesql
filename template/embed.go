package template

import (
	_ "embed"
	"strings"
	"text/template"

	"github.com/alimy/yesql/naming"
)

//go:embed sql.tmpl
var sqlTmplStr string

//go:embed sqlx.tmpl
var sqlxTmplStr string

func NewSqlTemplate() *template.Template {
	return template.Must(newTemplate("sql").Parse(sqlTmplStr))
}

func NewSqlxTemplate() *template.Template {
	return template.Must(newTemplate("sqlx").Parse(sqlxTmplStr))
}

func newTemplate(name string) *template.Template {
	return template.New(name).Funcs(template.FuncMap{
		"naming":      naming.Naming,
		"notEmptyStr": notEmptyStr,
		"escape":      escapeBacktick,
	})
}

func notEmptyStr(s string) bool {
	return len(s) > 0
}

// Go string literals cannot contain backtick. If a string contains
// a backtick, replace it the following way:
//
// input:
//
//	SELECT `group` FROM foo
//
// output:
//
//	SELECT ` + "`" + `group` + "`" + ` FROM foo
//
// # The escaped string must be rendered inside an existing string literal
//
// A string cannot be escaped twice
func escapeBacktick(s string) string {
	return strings.Replace(s, "`", "`+\"`\"+`", -1)
}
