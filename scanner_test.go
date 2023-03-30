package yesql

import (
	"testing"
)

func TestPrepareScan(t *testing.T) {
	type Q struct {
		Namespace   `yesql:"namespace"`
		Ignore      string
		SimpleQuery string `yesql:"$simple"`
		RawQuery    string `yesql:"multiline"`
	}
	type Q2 struct {
		RawQuery string `yesql:"does-not-exist"`
	}
	type Q3 struct {
		SimpleQuery   string `yesql:"$simple"`
		CommentsQuery string `yesql:"comments"`
	}

	var (
		q  Q
		q2 Q2
		q3 Q3
	)
	queries := MustParseFile("testdata/valid.sql")
	scanner := NewPrepareScanner(NewPrepareHook(nil))
	err := scanner.Scan(&q, queries)
	if err != nil {
		t.Errorf("[q] failed to scan raw query to struct: %s", err)
	}
	if q.SimpleQuery != `SELECT * FROM simple;` {
		t.Errorf("[q] want simple query but got %s", q.SimpleQuery)
	}
	if q.RawQuery != `SELECT * FROM multiline WHERE line = 42;` {
		t.Errorf("[q] want raw query but got %s", q.RawQuery)
	}

	if err = scanner.Scan(&q2, queries); err == nil {
		t.Error("[q2] expected to fail at non-existent query 'does-not-exist' but didn't")
	}

	if err = scanner.Scan(&q3, queries); err != nil {
		t.Errorf("[q3] failed to scan raw query to struct: %s", err)
	}
	if q3.SimpleQuery != `SELECT * FROM simple;` {
		t.Errorf("[q3] want simple query but got %s", q3.SimpleQuery)
	}
	if q3.CommentsQuery != `SELECT * FROM comments;` {
		t.Errorf("[q3] want simple query but got %s", q3.CommentsQuery)
	}
}
