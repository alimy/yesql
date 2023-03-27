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
		Ignore   string
		RawQuery string `yesql:"multiline"`
	}
	type Q2 struct {
		RawQuery string `yesql:"does-not-exist"`
	}
	var (
		q  Q
		q2 Q2
	)

	queries := MustParseFile("testdata/valid.sql")
	err := Scan(&q, queries, NewPrepareHook(nil))
	if err != nil {
		t.Errorf("failed to scan raw query to struct: %v", err)
	}

	err = Scan(&q2, queries, NewPrepareHook(nil))
	if err == nil {
		t.Error("expected to fail at non-existent query 'does-not-exist' but didn't")
	}
}

func BenchmarkMustParseFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustParseFile("testdata/valid.sql")
	}
}
