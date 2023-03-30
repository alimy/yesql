package main

import (
	"log"

	"github.com/alimy/yesql"
)

//go:generate go run $GOFILE
func main() {
	log.Println("Yesql generate code start")
	if err := yesql.Generate("yesql.sql", "auto", "yesql"); err != nil {
		log.Fatalf("generate code occurs error: %s", err)
	}
	log.Println("Yesql generate code finish")
}
