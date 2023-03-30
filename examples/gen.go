package main

import (
	"log"

	"github.com/alimy/yesql"
)

//go:generate go run $GOFILE
func main() {
	log.Println("Yesql generate code start")
	g, err := yesql.NewSqlxGenerator()
	if err != nil {
		log.Fatalf("create a sqlx generator error: %s", err)
	}
	yesql.Generate(g, "yesql.sql", "auto", "yesql",
		yesql.DefaultStructNameOpt("Yesql"),
		yesql.GoFileNameOpt("yesql.go"))
	log.Println("Yesql generate code finish")
}
