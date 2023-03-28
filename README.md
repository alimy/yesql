## Yesql
[![Go](https://github.com/alimy/yesql/actions/workflows/go.yml/badge.svg)](https://github.com/alimy/yesql/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/alimy/yesql?status.svg)](https://pkg.go.dev/github.com/alimy/yesql)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?logo=sourcegraph)](https://sourcegraph.com/github.com/alimy/yesql)

Yesql parses a file and associate SQL queries to a map. Useful for separating SQL from code logic.

> This package is based on [knadh/goyesql](https://github.com/knadh/goyesql) but is not compatible with it any more. This package introduces support for arbitrary tag types and changes structs and error types.

### Installation

```
$ go get github.com/alimy/yesql
```

### Usage

Create a file containing your SQL queries

```sql
-- sql file yesql.sql

-- name: newest_tags@topic
-- get newest tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t JOIN @user u ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.id DESC 
LIMIT ? OFFSET ?;

-- name: hot_tags@topic
-- get get host tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t JOIN @user u ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.quote_num DESC 
LIMIT ? OFFSET ?;

-- name: tags_by_keyword_a@topic
-- get tags by keyword
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6;

-- name: tags_by_keyword_b@topic
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6;

-- name: insert_tag@topic
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1);

-- name: tags_by_id_a@topic
-- clause: in
SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;

-- name: tags_by_id_b@topic
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);

-- name: decr_tags_by_id@topic
-- clause: in
UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);

-- name: tags_for_incr@topic
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);

-- name: incr_tags_by_id@topic
-- clause: in
UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);
```

And just call them in your code!

```go
// file: sqlx.go

package sakila

import (
	"context"
	"database/sql"
	_ "embed"
	"strings"

	"github.com/alimy/yesql"
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

//go:embed yesql.sql
var yesqlBytes []byte

type sqlxSrv struct {
	db *sqlx.DB
}

func (s *sqlxSrv) with(handle func(tx *sqlx.Tx) error) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err = handle(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *sqlxSrv) withTx(ctx context.Context, opts *sql.TxOptions, handle func(*sqlx.Tx) error) error {
	tx, err := s.db.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err = handle(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *sqlxSrv) in(query string, args ...any) (string, []any, error) {
	q, params, err := sqlx.In(query, args...)
	if err != nil {
		return "", nil, err
	}
	return s.db.Rebind(q), params, nil
}

func (s *sqlxSrv) inExec(execer sqlx.Execer, query string, args ...any) (sql.Result, error) {
	q, params, err := sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	return execer.Exec(s.db.Rebind(q), params...)
}

func (s *sqlxSrv) inSelect(queryer sqlx.Queryer, dest any, query string, args ...any) error {
	q, params, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}
	return sqlx.Select(queryer, dest, s.db.Rebind(q), params...)
}

func (s *sqlxSrv) inGet(queryer sqlx.Queryer, dest any, query string, args ...any) error {
	q, params, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}
	return sqlx.Get(queryer, dest, s.db.Rebind(q), params...)
}

func newSqlxSrv(db *sqlx.DB) *sqlxSrv {
	return &sqlxSrv{
		db: db,
	}
}

func initSqlxDB() {
	_db = conf.MustSqlxDB()
	yesql.UseSqlx(_db)
	yesql.SetDefaultQueryHooks(func(query *yesql.Query) (*yesql.Query, error) {
		qstr := strings.TrimRight(query.Query, ";")
		// table name fixed
		qstr = strings.Replace(qstr, "@", conf.DatabaseSetting.TablePrefix, -1)
		// rebind query
		if clause, exist := query.Tags["clause"]; !exist || clause != "in" {
			qstr = _db.Rebind(qstr)
		}
		query.Query = qstr
		return query, nil
	})
}
```

### Scanning

Often, it's necessary to scan multiple queries from a SQL file, prepare them into \*sql.Stmt and use them throught the application. yesql comes with a helper function that helps with this. Given a yesql map of queries, it can turn the queries into prepared statements and scan them into a struct that can be passed around.

```go
// file: sakila.go

package sakila

import (
	"strings"

	"github.com/alimy/yesql"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/sirupsen/logrus"
)

type topicSrv struct {
	*sqlxSrv
	yesql.Namespace    `yesql:"topic"`
	StmtNewestTags     *sqlx.Stmt `yesql:"newest_tags"`
	StmtHotTags        *sqlx.Stmt `yesql:"hot_tags"`
	StmtTagsByKeywordA *sqlx.Stmt `yesql:"tags_by_keyword_a"`
	StmtTagsByKeywordB *sqlx.Stmt `yesql:"tags_by_keyword_b"`
	StmtInsertTag      *sqlx.Stmt `yesql:"insert_tag"`
	SqlTagsByIdA       string     `yesql:"tags_by_id_a"`
	SqlTagsByIdB       string     `yesql:"tags_by_id_b"`
	SqlDecrTagsById    string     `yesql:"decr_tags_by_id"`
	SqlTagsForIncr     string     `yesql:"tags_for_incr"`
	SqlIncrTagsById    string     `yesql:"incr_tags_by_id"`
}

func (s *topicSrv) ListTags(typ cs.TagType, limit int, offset int) (res cs.TagList, err error) {
	switch typ {
	case cs.TagTypeHot:
		err = s.StmtHotTags.Select(&res, limit, offset)
	case cs.TagTypeNew:
		err = s.StmtNewestTags.Select(&res, limit, offset)
	}
	return
}

func newTopicService(db *sqlx.DB) core.TopicService {
	initSqlxDB()
	
	obj := &topicSrv{
		sqlxSrv: newSqlxSrv(db),
	}
	query := yesql.MustParseBytes(yesqlBytes)
	if err := yesql.Scan(obj, query); err != nil {
		logrus.Fatal(err)
	}
	return obj
}
```
> Source code from [github.com/rocboss/paopao-ce](https://github.com/rocboss/paopao-ce/tree/r/paopao-ce-plus/internal/dao/sakila).

### Projects that used [Yesql](https://github.com/alimy/yesql) 
* [paopao-ce](https://github.com/rocboss/paopao-ce/tree/r/paopao-ce-plus) - A artistic "twitter like" community built on gin+zinc+vue+ts.      