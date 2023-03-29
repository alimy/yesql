package yesql

import (
	"testing"
)

func TestMustParseFilePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseFile should panic if an error occurs, got '%s'", r)
		}
	}()
	MustParseFile("testdata/missing.sql")
}

func TestMustParseFileNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustParseFile should not panic if no error occurs, got '%s'", r)
		}
	}()
	MustParseFile("testdata/valid.sql")
}

func TestMustParseBytesPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseBytes should panic if an error occurs, got '%s'", r)
		}
	}()
	MustParseBytes([]byte("I won't work"))
}

func TestMustParseBytesNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustParseBytes should not panic if an error occurs, got '%s'", r)
		}
	}()
	MustParseBytes([]byte("-- name: byte-me\nSELECT * FROM bytes;"))
}

func TestScan(t *testing.T) {
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

	err := Scan(&q, queries, NewPrepareHook(nil))
	if err != nil {
		t.Errorf("[q] failed to scan raw query to struct: %s", err)
	}
	if q.SimpleQuery != `SELECT * FROM simple;` {
		t.Errorf("[q] want simple query but got %s", q.SimpleQuery)
	}
	if q.RawQuery != `SELECT * FROM multiline WHERE line = 42;` {
		t.Errorf("[q] want raw query but got %s", q.RawQuery)
	}

	if err = Scan(&q2, queries, NewPrepareHook(nil)); err == nil {
		t.Error("[q2] expected to fail at non-existent query 'does-not-exist' but didn't")
	}

	SetDefaultPrepareHook(NewSqlxPrepareHook(nil))
	if err = Scan(&q3, queries); err != nil {
		t.Errorf("[q3] failed to scan raw query to struct: %s", err)
	}
	if q3.SimpleQuery != `SELECT * FROM simple;` {
		t.Errorf("[q3] want simple query but got %s", q3.SimpleQuery)
	}
	if q3.CommentsQuery != `SELECT * FROM comments;` {
		t.Errorf("[q3] want simple query but got %s", q3.CommentsQuery)
	}
}

func BenchmarkMustParseFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustParseFile("testdata/valid.sql")
	}
}
