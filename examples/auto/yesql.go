// Code generated by go-mir. DO NOT EDIT.
// versions:
// - Yesql v1.0.0

package yesql

import (
	"context"

	"github.com/alimy/yesql"
	"github.com/jmoiron/sqlx"
)

const (
	_LoginInfoD              = `SELECT * FROM @user WHERE username=?`
	_LoginInfoE              = `SELECT * FROM @user WHERE username=?`
	_UserInfoA               = `SELECT ` + "`" + `username` + "`" + ` FROM @user WHERE username=?`
	_UserInfoB               = `SELECT ` + "`" + `username` + "`" + `, ` + "`" + `nickname` + "`" + ` FROM @user WHERE username=?`
	_UserInfoC               = `SELECT "username" FROM @user WHERE username=?`
	_LogoutInfoC             = `SELECT * FROM @user WHERE username=?`
	_UserInfoD               = `SELECT "username", "nickname" FROM @user WHERE username=?`
	_LoginInfoA              = `SELECT * FROM @user WHERE username=?`
	_LoginInfoB              = `SELECT * FROM @user WHERE username=?`
	_NewestTags_TagsInfo     = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.id DESC LIMIT ? OFFSET ?;`
	_InsertTag_TagsInfo      = `INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1);`
	_TagsByIdA_TagsInfo      = `SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;`
	_TagsByIdB_TagsInfo      = `SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);`
	_IncrTagsById_TagsInfo   = `UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);`
	_HotTags_TagsInfo        = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.quote_num DESC LIMIT ? OFFSET ?;`
	_TagsByKeywordA_TagsInfo = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6;`
	_TagsByKeywordB_TagsInfo = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6;`
	_DecrTagsById_TagsInfo   = `UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);`
	_TagsForIncr_TagsInfo    = `SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);`
	_TagsByKeywordB_Topic    = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6;`
	_InsertTag_Topic         = `INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1);`
	_TagsByIdB_Topic         = `SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);`
	_HotTags_Topic           = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.quote_num DESC LIMIT ? OFFSET ?;`
	_TagsByKeywordA_Topic    = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6;`
	_TagsByIdA_Topic         = `SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;`
	_DecrTagsById_Topic      = `UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);`
	_TagsForIncr_Topic       = `SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);`
	_IncrTagsById_Topic      = `UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);`
	_NewestTags_Topic        = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.id DESC LIMIT ? OFFSET ?;`
	_IncrTagsById_Shutter    = `UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);`
	_HotTags_Shutter         = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.quote_num DESC LIMIT ? OFFSET ?;`
	_TagsByIdB_Shutter       = `SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);`
	_DecrTagsById_Shutter    = `UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);`
	_InsertTag_Shutter       = `INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1);`
	_TagsByIdA_Shutter       = `SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;`
	_TagsForIncr_Shutter     = `SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);`
	_NewestTags_Shutter      = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.id DESC LIMIT ? OFFSET ?;`
	_TagsByKeywordA_Shutter  = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6;`
	_TagsByKeywordB_Shutter  = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6;`
)

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

func BuildYesql(p yesql.PreparexContext, ctx ...context.Context) (obj *Yesql, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] == nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Yesql{
		LoginInfoD:  _LoginInfoD,
		LogoutInfoC: _LogoutInfoC,
		UserInfoC:   _UserInfoC,
		UserInfoD:   _UserInfoD,
	}
	if obj.LoginInfoB, err = p.PreparexContext(c, p.Rebind(_LoginInfoB)); err != nil {
		return
	}
	if obj.UserInfoA, err = p.PreparexContext(c, p.Rebind(_UserInfoA)); err != nil {
		return
	}
	if obj.UserInfoB, err = p.PreparexContext(c, p.Rebind(_UserInfoB)); err != nil {
		return
	}
	if obj.LoginInfoA, err = p.PrepareNamedContext(c, p.Rebind(_LoginInfoA)); err != nil {
		return
	}
	if obj.LoginInfoE, err = p.PrepareNamedContext(c, p.Rebind(_LoginInfoE)); err != nil {
		return
	}
	return
}

func BuildShutter(p yesql.PreparexContext, ctx ...context.Context) (obj *Shutter, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] == nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Shutter{
		DecrTagsById: _DecrTagsById_Shutter,
		IncrTagsById: _IncrTagsById_Shutter,
		TagsByIdA:    _TagsByIdA_Shutter,
		TagsByIdB:    _TagsByIdB_Shutter,
		TagsForIncr:  _TagsForIncr_Shutter,
	}
	if obj.HotTags, err = p.PreparexContext(c, p.Rebind(_HotTags_Shutter)); err != nil {
		return
	}
	if obj.TagsByKeywordA, err = p.PreparexContext(c, p.Rebind(_TagsByKeywordA_Shutter)); err != nil {
		return
	}
	if obj.TagsByKeywordB, err = p.PreparexContext(c, p.Rebind(_TagsByKeywordB_Shutter)); err != nil {
		return
	}
	if obj.InsertTag, err = p.PrepareNamedContext(c, p.Rebind(_InsertTag_Shutter)); err != nil {
		return
	}
	if obj.NewestTags, err = p.PrepareNamedContext(c, p.Rebind(_NewestTags_Shutter)); err != nil {
		return
	}
	return
}

func BuildTagsInfo(p yesql.PreparexContext, ctx ...context.Context) (obj *TagsInfo, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] == nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &TagsInfo{
		DecrTagsById: _DecrTagsById_TagsInfo,
		IncrTagsById: _IncrTagsById_TagsInfo,
		TagsByIdA:    _TagsByIdA_TagsInfo,
		TagsByIdB:    _TagsByIdB_TagsInfo,
		TagsForIncr:  _TagsForIncr_TagsInfo,
	}
	if obj.HotTags, err = p.PreparexContext(c, p.Rebind(_HotTags_TagsInfo)); err != nil {
		return
	}
	if obj.TagsByKeywordA, err = p.PreparexContext(c, p.Rebind(_TagsByKeywordA_TagsInfo)); err != nil {
		return
	}
	if obj.TagsByKeywordB, err = p.PreparexContext(c, p.Rebind(_TagsByKeywordB_TagsInfo)); err != nil {
		return
	}
	if obj.InsertTag, err = p.PrepareNamedContext(c, p.Rebind(_InsertTag_TagsInfo)); err != nil {
		return
	}
	if obj.NewestTags, err = p.PrepareNamedContext(c, p.Rebind(_NewestTags_TagsInfo)); err != nil {
		return
	}
	return
}

func BuildTopic(p yesql.PreparexContext, ctx ...context.Context) (obj *Topic, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] == nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Topic{
		DecrTagsById: _DecrTagsById_Topic,
		IncrTagsById: _IncrTagsById_Topic,
		TagsByIdA:    _TagsByIdA_Topic,
		TagsByIdB:    _TagsByIdB_Topic,
		TagsForIncr:  _TagsForIncr_Topic,
	}
	if obj.HotTags, err = p.PreparexContext(c, p.Rebind(_HotTags_Topic)); err != nil {
		return
	}
	if obj.TagsByKeywordA, err = p.PreparexContext(c, p.Rebind(_TagsByKeywordA_Topic)); err != nil {
		return
	}
	if obj.TagsByKeywordB, err = p.PreparexContext(c, p.Rebind(_TagsByKeywordB_Topic)); err != nil {
		return
	}
	if obj.InsertTag, err = p.PrepareNamedContext(c, p.Rebind(_InsertTag_Topic)); err != nil {
		return
	}
	if obj.NewestTags, err = p.PrepareNamedContext(c, p.Rebind(_NewestTags_Topic)); err != nil {
		return
	}
	return
}
