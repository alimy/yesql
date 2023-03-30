package yesql

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := map[string]parsedLine{
		" ":                {lineBlank, "", ""},
		" SELECT * ":       {lineQuery, "", "SELECT *"},
		" -- name: tag ":   {lineTag, "name", "tag"},
		" -- some: param ": {lineTag, "some", "param"},
		" -- comment ":     {lineComment, "", "comment"},
		" --":              {lineComment, "", ""},
	}

	for line, expected := range tests {
		parsed := parseLine(line)

		if parsed != expected {
			t.Errorf("Invalid line parsing. Expected '%v', got '%v'", expected, parsed)
		}
	}
}

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

	expectedDefQueries := QueryMap{
		"simple":   &Query{Query: "SELECT * FROM simple;"},
		"comments": &Query{Query: "SELECT * FROM comments;"},
	}
	expectedNsQueries := QueryMap{
		"multiline": &Query{Query: "SELECT * FROM multiline WHERE line = 42;"},
	}

	queries, _ := sqlQuery.ListQuery()
	if len(queries) != len(expectedDefQueries) {
		t.Errorf(
			"%s should return %d requests, got %d",
			file, len(expectedDefQueries), len(queries),
		)
	}
	if len(queries["simple"].Tags) != 1 ||
		queries["simple"].Tags["raw"] != "1" {
		t.Errorf("Tag 'raw = 1' not found in 'simple' valid query")
	}
	for key, expectedQuery := range expectedDefQueries {
		if queries[key].Query != expectedQuery.Query {
			t.Errorf(
				"%s query should be '%s', got '%s'",
				key, expectedQuery, queries[key],
			)
		}
	}

	nsQueries, _ := sqlQuery.ListQuery("namespace")
	if len(nsQueries) != len(expectedNsQueries) {
		t.Errorf(
			"%s should return %d requests, got %d",
			file, len(expectedNsQueries), len(nsQueries),
		)
	}
	for key, expectedNsQuery := range expectedNsQueries {
		if nsQueries[key].Query != expectedNsQuery.Query {
			t.Errorf(
				"%s query should be '%s', got '%s'",
				key, expectedNsQuery, queries[key],
			)
		}
	}
}
