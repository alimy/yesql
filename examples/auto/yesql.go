// Code generated by Yesql. DO NOT EDIT.
// versions:
// - Yesql v1.9.0

package yesql

import (
	"context"
	"fmt"

	"github.com/alimy/yesql"
	"github.com/bitbus/sqlx"
)

var (
	_ = fmt.Errorf("error for placeholder")
)

const (
	_LoginInfoA              = `SELECT * FROM @user WHERE username=?`
	_LoginInfoB              = `SELECT * FROM @user WHERE username=?`
	_LoginInfoD              = `SELECT * FROM @user WHERE username=?`
	_LoginInfoE              = `SELECT * FROM @user WHERE username=?`
	_LogoutInfoC             = `SELECT * FROM @user WHERE username=?`
	_UserInfoA               = `SELECT ` + "`" + `username` + "`" + ` FROM @user WHERE username=?`
	_UserInfoB               = `SELECT ` + "`" + `username` + "`" + `, ` + "`" + `nickname` + "`" + ` FROM @user WHERE username=?`
	_UserInfoC               = `SELECT "username" FROM @user WHERE username=?`
	_UserInfoD               = `SELECT "username", "nickname" FROM @user WHERE username=?`
	_Shutter_DecrTagsById    = `UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?)`
	_Shutter_HotTags         = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.quote_num DESC LIMIT ? OFFSET ?`
	_Shutter_IncrTagsById    = `UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?)`
	_Shutter_InsertTag       = `INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1)`
	_Shutter_NewestTags      = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.id DESC LIMIT ? OFFSET ?`
	_Shutter_TagsByIdA       = `SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0`
	_Shutter_TagsByIdB       = `SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?)`
	_Shutter_TagsByKeywordA  = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6`
	_Shutter_TagsByKeywordB  = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6`
	_Shutter_TagsForIncr     = `SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?)`
	_TagsInfo_DecrTagsById   = `UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?)`
	_TagsInfo_HotTags        = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.quote_num DESC LIMIT ? OFFSET ?`
	_TagsInfo_IncrTagsById   = `UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?)`
	_TagsInfo_InsertTag      = `INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1)`
	_TagsInfo_NewestTags     = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.id DESC LIMIT ? OFFSET ?`
	_TagsInfo_TagsByIdA      = `SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0`
	_TagsInfo_TagsByIdB      = `SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?)`
	_TagsInfo_TagsByKeywordA = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6`
	_TagsInfo_TagsByKeywordB = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6`
	_TagsInfo_TagsForIncr    = `SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?)`
	_Topic_DecrTagsById      = `UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?)`
	_Topic_HotTags           = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.quote_num DESC LIMIT ? OFFSET ?`
	_Topic_IncrTagsById      = `UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?)`
	_Topic_InsertTag         = `INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1)`
	_Topic_NewestTags        = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.id DESC LIMIT ? OFFSET ?`
	_Topic_TagsByIdA         = `SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0`
	_Topic_TagsByIdB         = `SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?)`
	_Topic_TagsByKeywordA    = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6`
	_Topic_TagsByKeywordB    = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6`
	_Topic_TagsForIncr       = `SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?)`
)

// PreparexContext enhances the Conn interface with context.
type PreparexContext interface {
	// PrepareContext prepares a statement.
	// The provided context is used for the preparation of the statement, not for
	// the execution of the statement.
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)

	// PrepareNamedContext returns an sqlx.NamedStmt
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)

	// Rebind rebind query to adapte SQL Driver
	Rebind(query string) string
}

// PreparexBuilder preparex builder interface for sqlx
type PreparexBuilder interface {
	PreparexContext
	QueryHook(query string) string
}

type Yesql struct {
	LoginInfoD  string          `yesql:"login_info_d"`
	LogoutInfoC string          `yesql:"logout_info_c"`
	UserInfoC   string          `yesql:"user_info_c"`
	UserInfoD   string          `yesql:"user_info_d"`
	LoginInfoB  *sqlx.Stmt      `yesql:"login_info_b"`
	UserInfoA   *sqlx.Stmt      `yesql:"user_info_a"`
	UserInfoB   *sqlx.Stmt      `yesql:"user_info_b"`
	LoginInfoA  *sqlx.NamedStmt `yesql:"login_info_a"`
	LoginInfoE  *sqlx.NamedStmt `yesql:"login_info_e"`
}

type Shutter struct {
	yesql.Namespace `yesql:"shutter"`
	DecrTagsById    string          `yesql:"decr_tags_by_id"`
	IncrTagsById    string          `yesql:"incr_tags_by_id"`
	TagsByIdA       string          `yesql:"tags_by_id_a"`
	TagsByIdB       string          `yesql:"tags_by_id_b"`
	TagsForIncr     string          `yesql:"tags_for_incr"`
	HotTags         *sqlx.Stmt      `yesql:"hot_tags"`
	TagsByKeywordA  *sqlx.Stmt      `yesql:"tags_by_keyword_a"`
	TagsByKeywordB  *sqlx.Stmt      `yesql:"tags_by_keyword_b"`
	InsertTag       *sqlx.NamedStmt `yesql:"insert_tag"`
	NewestTags      *sqlx.NamedStmt `yesql:"newest_tags"`
}

type TagsInfo struct {
	yesql.Namespace `yesql:"tags_info"`
	DecrTagsById    string          `yesql:"decr_tags_by_id"`
	IncrTagsById    string          `yesql:"incr_tags_by_id"`
	TagsByIdA       string          `yesql:"tags_by_id_a"`
	TagsByIdB       string          `yesql:"tags_by_id_b"`
	TagsForIncr     string          `yesql:"tags_for_incr"`
	HotTags         *sqlx.Stmt      `yesql:"hot_tags"`
	TagsByKeywordA  *sqlx.Stmt      `yesql:"tags_by_keyword_a"`
	TagsByKeywordB  *sqlx.Stmt      `yesql:"tags_by_keyword_b"`
	InsertTag       *sqlx.NamedStmt `yesql:"insert_tag"`
	NewestTags      *sqlx.NamedStmt `yesql:"newest_tags"`
}

type Topic struct {
	yesql.Namespace `yesql:"topic"`
	DecrTagsById    string          `yesql:"decr_tags_by_id"`
	IncrTagsById    string          `yesql:"incr_tags_by_id"`
	TagsByIdA       string          `yesql:"tags_by_id_a"`
	TagsByIdB       string          `yesql:"tags_by_id_b"`
	TagsForIncr     string          `yesql:"tags_for_incr"`
	HotTags         *sqlx.Stmt      `yesql:"hot_tags"`
	TagsByKeywordA  *sqlx.Stmt      `yesql:"tags_by_keyword_a"`
	TagsByKeywordB  *sqlx.Stmt      `yesql:"tags_by_keyword_b"`
	InsertTag       *sqlx.NamedStmt `yesql:"insert_tag"`
	NewestTags      *sqlx.NamedStmt `yesql:"newest_tags"`
}

func BuildYesql(p PreparexBuilder, ctx ...context.Context) (obj *Yesql, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Yesql{
		LoginInfoD:  p.QueryHook(_LoginInfoD),
		LogoutInfoC: p.QueryHook(_LogoutInfoC),
		UserInfoC:   p.QueryHook(_UserInfoC),
		UserInfoD:   p.QueryHook(_UserInfoD),
	}
	if obj.LoginInfoB, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_LoginInfoB))); err != nil {
		return nil, fmt.Errorf("prepare _LoginInfoB error: %w", err)
	}
	if obj.UserInfoA, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_UserInfoA))); err != nil {
		return nil, fmt.Errorf("prepare _UserInfoA error: %w", err)
	}
	if obj.UserInfoB, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_UserInfoB))); err != nil {
		return nil, fmt.Errorf("prepare _UserInfoB error: %w", err)
	}
	if obj.LoginInfoA, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_LoginInfoA))); err != nil {
		return nil, fmt.Errorf("prepare _LoginInfoA error: %w", err)
	}
	if obj.LoginInfoE, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_LoginInfoE))); err != nil {
		return nil, fmt.Errorf("prepare _LoginInfoE error: %w", err)
	}
	return
}

func BuildShutter(p PreparexBuilder, ctx ...context.Context) (obj *Shutter, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Shutter{
		DecrTagsById: p.QueryHook(_Shutter_DecrTagsById),
		IncrTagsById: p.QueryHook(_Shutter_IncrTagsById),
		TagsByIdA:    p.QueryHook(_Shutter_TagsByIdA),
		TagsByIdB:    p.QueryHook(_Shutter_TagsByIdB),
		TagsForIncr:  p.QueryHook(_Shutter_TagsForIncr),
	}
	if obj.HotTags, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_Shutter_HotTags))); err != nil {
		return nil, fmt.Errorf("prepare _Shutter_HotTags error: %w", err)
	}
	if obj.TagsByKeywordA, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_Shutter_TagsByKeywordA))); err != nil {
		return nil, fmt.Errorf("prepare _Shutter_TagsByKeywordA error: %w", err)
	}
	if obj.TagsByKeywordB, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_Shutter_TagsByKeywordB))); err != nil {
		return nil, fmt.Errorf("prepare _Shutter_TagsByKeywordB error: %w", err)
	}
	if obj.InsertTag, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_Shutter_InsertTag))); err != nil {
		return nil, fmt.Errorf("prepare _Shutter_InsertTag error: %w", err)
	}
	if obj.NewestTags, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_Shutter_NewestTags))); err != nil {
		return nil, fmt.Errorf("prepare _Shutter_NewestTags error: %w", err)
	}
	return
}

func BuildTagsInfo(p PreparexBuilder, ctx ...context.Context) (obj *TagsInfo, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &TagsInfo{
		DecrTagsById: p.QueryHook(_TagsInfo_DecrTagsById),
		IncrTagsById: p.QueryHook(_TagsInfo_IncrTagsById),
		TagsByIdA:    p.QueryHook(_TagsInfo_TagsByIdA),
		TagsByIdB:    p.QueryHook(_TagsInfo_TagsByIdB),
		TagsForIncr:  p.QueryHook(_TagsInfo_TagsForIncr),
	}
	if obj.HotTags, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TagsInfo_HotTags))); err != nil {
		return nil, fmt.Errorf("prepare _TagsInfo_HotTags error: %w", err)
	}
	if obj.TagsByKeywordA, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TagsInfo_TagsByKeywordA))); err != nil {
		return nil, fmt.Errorf("prepare _TagsInfo_TagsByKeywordA error: %w", err)
	}
	if obj.TagsByKeywordB, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TagsInfo_TagsByKeywordB))); err != nil {
		return nil, fmt.Errorf("prepare _TagsInfo_TagsByKeywordB error: %w", err)
	}
	if obj.InsertTag, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_TagsInfo_InsertTag))); err != nil {
		return nil, fmt.Errorf("prepare _TagsInfo_InsertTag error: %w", err)
	}
	if obj.NewestTags, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_TagsInfo_NewestTags))); err != nil {
		return nil, fmt.Errorf("prepare _TagsInfo_NewestTags error: %w", err)
	}
	return
}

func BuildTopic(p PreparexBuilder, ctx ...context.Context) (obj *Topic, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Topic{
		DecrTagsById: p.QueryHook(_Topic_DecrTagsById),
		IncrTagsById: p.QueryHook(_Topic_IncrTagsById),
		TagsByIdA:    p.QueryHook(_Topic_TagsByIdA),
		TagsByIdB:    p.QueryHook(_Topic_TagsByIdB),
		TagsForIncr:  p.QueryHook(_Topic_TagsForIncr),
	}
	if obj.HotTags, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_Topic_HotTags))); err != nil {
		return nil, fmt.Errorf("prepare _Topic_HotTags error: %w", err)
	}
	if obj.TagsByKeywordA, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_Topic_TagsByKeywordA))); err != nil {
		return nil, fmt.Errorf("prepare _Topic_TagsByKeywordA error: %w", err)
	}
	if obj.TagsByKeywordB, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_Topic_TagsByKeywordB))); err != nil {
		return nil, fmt.Errorf("prepare _Topic_TagsByKeywordB error: %w", err)
	}
	if obj.InsertTag, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_Topic_InsertTag))); err != nil {
		return nil, fmt.Errorf("prepare _Topic_InsertTag error: %w", err)
	}
	if obj.NewestTags, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_Topic_NewestTags))); err != nil {
		return nil, fmt.Errorf("prepare _Topic_NewestTags error: %w", err)
	}
	return
}
