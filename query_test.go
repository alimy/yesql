package yesql

import (
	"testing"
)

func TestScannerErrTags(t *testing.T) {
	for _, key := range []string{"missing", "doubloon"} {
		_, err := ParseFile("testdata/tag_" + key + ".sql")
		if err == nil {
			t.Errorf("Expected error, but got nil.")
		}
	}
}

func TestScannerValid(t *testing.T) {
	file := "testdata/valid.sql"

	sqlQuery, err := ParseFile(file)
	if err != nil {
		t.Fatal(err)
	}

	expectedQueries := QueryMap{
		"simple":    &Query{Query: "SELECT * FROM simple;"},
		"multiline": &Query{Query: "SELECT * FROM multiline WHERE line = 42;"},
		"comments":  &Query{Query: "SELECT * FROM comments;"},
	}

	queries, _ := sqlQuery.ListQuery("")
	if len(queries) != len(expectedQueries) {
		t.Errorf(
			"%s should return %d requests, got %d",
			file, len(expectedQueries), len(queries),
		)
	}

	if len(queries["simple"].Tags) != 1 ||
		queries["simple"].Tags["raw"] != "1" {
		t.Errorf("Tag 'raw = 1' not found in 'simple' valid query")
	}

	for key, expectedQuery := range expectedQueries {
		if queries[key].Query != expectedQuery.Query {
			t.Errorf(
				"%s query should be '%s', got '%s'",
				key, expectedQuery, queries[key],
			)
		}
	}
}
