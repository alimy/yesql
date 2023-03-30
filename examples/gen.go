package main

import (
	"log"

	"github.com/alimy/yesql"
)

//go:generate go run $GOFILE
func main() {
	log.Println("Yesql generate code start")
	g, err := yesql.NewSqlGenerator()
	if err != nil {
		log.Fatalf("create a sqlx generator error: %s", err)
	}
	if err := yesql.Generate(g, "yesql.sql", "auto", "yesql"); err != nil {
		log.Fatalf("generate code occurs error: %s", err)
	}
	log.Println("Yesql generate code finish")
}
