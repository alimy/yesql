package main

import (
	"log"
	"strings"

	"github.com/alimy/yesql"
)

//go:generate go run $GOFILE
func main() {
	log.Println("[Yesql] generate code start")
	yesql.SetDefaultQueryHook(func(query *yesql.Query) (*yesql.Query, error) {
		query.Query = strings.TrimRight(query.Query, ";")
		return query, nil
	})
	opt := yesql.SqlxPkgName("github.com/bitbus/sqlx")
	if err := yesql.Generate("yesql.sql", "auto", "yesql", opt); err != nil {
		log.Fatalf("generate yesql code occurs error: %s", err)
	}
	if err := yesql.Generate("yesql_ac.sql", "auto/ac", "ac", opt); err != nil {
		log.Fatalf("generate ac code occurs error: %s", err)
	}
	if err := yesql.Generate("yesql_cc.sql", "auto/cc", "cc", opt); err != nil {
		log.Fatalf("generate cc code occurs error: %s", err)
	}
	log.Println("[Yesql] generate code finish")
}
