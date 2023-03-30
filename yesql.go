// Package yesql is a Go port of Yesql
//
// It allows you to write SQL queries in separate files.
// See rationale at https://github.com/krisajenkins/yesql#rationale
package yesql

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	_structTag             = "yesql"
	_defaultNamingStrategy = NewFieldUpNamingStrategy()
	_defaultPrepareScanner PrepareScanner
	_defaultQueryHooks     []func(query *Query) (*Query, error)
)

// Use use default prepare scanner with prepare that implement PrepareContext
func Use(p PrepareContext) {
	prepareHook := NewPrepareHook(p)
	_defaultPrepareScanner = NewPrepareScanner(prepareHook)
}

// UseSqlx use default prepare scanner withprepare that implement PreparexContext
func UseSqlx(p PreparexContext) {
	prepareHook := NewSqlxPrepareHook(p)
	_defaultPrepareScanner = NewPrepareScanner(prepareHook)
}

// SetDeafultTag set default struct tag
func SetDefaultTag(tag string) {
	tag = strings.Trim(tag, " ")
	if len(tag) > 0 {
		_structTag = tag
	}
}

// SetDefaultPrepareHook set default prepare hook
// Reset default prepare hook if hook is nil
func SetDefaultPrepareHook(hook PrepareHook) {
	_defaultPrepareScanner = nil
	if hook != nil {
		_defaultPrepareScanner = NewPrepareScanner(hook)
	}
}

// SetDefaultQueryHook set default query hooks
func SetDefaultQueryHook(hooks ...func(query *Query) (*Query, error)) {
	_defaultQueryHooks = nil
	for _, hook := range hooks {
		if hook != nil {
			_defaultQueryHooks = append(_defaultQueryHooks, hook)
		}
	}
}

func SetNamingStrategy(ns NamingStrategy) {
	if ns != nil {
		_defaultNamingStrategy = ns
	}
}

// ParseFile reads a file and return Queries or an error
func ParseFile(path string, hooks ...func(query *Query) (*Query, error)) (SQLQuery, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ParseReader(file, hooks...)
}

// MustParseFile calls ParseFile but panic if an error occurs
func MustParseFile(path string, hooks ...func(query *Query) (*Query, error)) SQLQuery {
	queries, err := ParseFile(path, hooks...)
	if err != nil {
		panic(err)
	}

	return queries
}

// ParseBytes parses bytes and returns Queries or an error.
func ParseBytes(b []byte, hooks ...func(query *Query) (*Query, error)) (SQLQuery, error) {
	return ParseReader(bytes.NewReader(b), hooks...)
}

// MustParseBytes parses bytes but panics if an error occurs.
func MustParseBytes(b []byte, hooks ...func(query *Query) (*Query, error)) SQLQuery {
	queries, err := ParseBytes(b)
	if err != nil {
		panic(err)
	}
	return queries
}

// ParseReader takes an io.Reader and returns Queries or an error.
func ParseReader(reader io.Reader, hooks ...func(query *Query) (*Query, error)) (SQLQuery, error) {
	parser := newSQLParser(_defaultQueryHooks...)
	parser.AddHooks(hooks...)
	query, err := parser.ParseReader(reader)
	if err != nil {
		return nil, err
	}
	return query, nil
}

// Scan scan object from a SQLQuery
func Scan(obj any, query SQLQuery, hook ...PrepareHook) error {
	return ScanContext(context.Background(), obj, query, hook...)
}

// ScanContext scan object from a SQLQuery with context.Context
func ScanContext(ctx context.Context, obj any, query SQLQuery, hook ...PrepareHook) error {
	scanner := _defaultPrepareScanner
	if len(hook) > 0 && hook[0] != nil {
		scanner = NewPrepareScanner(hook[0])
	}
	if scanner == nil {
		return fmt.Errorf("prepare hook must set or set a default prepare hook")
	}
	return scanner.ScanContext(ctx, obj, query)
}

// Generate generate struct type autumatic by sql file
func Generate(g Generator, sqlFilePath string, dstPath string, pkgName string, opts ...option) error {
	query, err := ParseFile(sqlFilePath)
	if err != nil {
		return err
	}
	return g.Generate(dstPath, pkgName, query, opts...)
}
