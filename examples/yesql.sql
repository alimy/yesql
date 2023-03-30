----------------------
-- default sql ddl --
----------------------

-- name: user_info_a
-- prepare: stmt
SELECT * FROM @user WHERE username=?

-- name: user_info_b
SELECT * FROM @user WHERE username=?

-- name: user_info_c
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_info_d
-- prepare: string
SELECT * FROM @user WHERE username=?

-- name: login_info_a
-- prepare: named_stmt
SELECT * FROM @user WHERE username=?

-- name: login_info_b
SELECT * FROM @user WHERE username=?

-- name: logout_info_c
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: login_info_d
-- prepare: string
SELECT * FROM @user WHERE username=?

-- name: login_info_e
-- prepare: named_stmt
SELECT * FROM @user WHERE username=?

----------------------
-- topic sql ddl --
----------------------

-- name: newest_tags@topic
-- prepare: named_stmt
-- get newest tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.id DESC 
LIMIT ? OFFSET ?;

-- name: hot_tags@topic
-- prepare: stmt
-- get get host tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.quote_num DESC 
LIMIT ? OFFSET ?;

-- name: tags_by_keyword_a@topic
-- get tags by keyword
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6;

-- name: tags_by_keyword_b@topic
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6;

-- name: insert_tag@topic
-- prepare: named_stmt
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1);

-- name: tags_by_id_a@topic
-- prepare: string
-- clause: in
SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;

-- name: tags_by_id_b@topic
-- prepare: string
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);

-- name: decr_tags_by_id@topic
-- prepare: string
-- clause: in
UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);

-- name: tags_for_incr@topic
-- prepare: string
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);

-- name: incr_tags_by_id@topic
-- prepare: string
-- clause: in
UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);

----------------------
-- shutter sql ddl ---
----------------------

-- name: newest_tags@shutter
-- prepare: named_stmt
-- get newest tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.id DESC 
LIMIT ? OFFSET ?;

-- name: hot_tags@shutter
-- prepare: stmt
-- get get host tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.quote_num DESC 
LIMIT ? OFFSET ?;

-- name: tags_by_keyword_a@shutter
-- get tags by keyword
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6;

-- name: tags_by_keyword_b@shutter
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6;

-- name: insert_tag@shutter
-- prepare: named_stmt
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1);

-- name: tags_by_id_a@shutter
-- prepare: string
-- clause: in
SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;

-- name: tags_by_id_b@shutter
-- prepare: string
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);

-- name: decr_tags_by_id@shutter
-- prepare: string
-- clause: in
UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);

-- name: tags_for_incr@shutter
-- prepare: string
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);

-- name: incr_tags_by_id@shutter
-- prepare: string
-- clause: in
UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);

----------------------
-- tags_info sql ddl ---
----------------------

-- name: newest_tags@tags_info
-- prepare: named_stmt
-- get newest tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.id DESC 
LIMIT ? OFFSET ?;

-- name: hot_tags@tags_info
-- prepare: stmt
-- get get host tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.quote_num DESC 
LIMIT ? OFFSET ?;

-- name: tags_by_keyword_a@tags_info
-- get tags by keyword
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6;

-- name: tags_by_keyword_b@tags_info
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6;

-- name: insert_tag@tags_info
-- prepare: named_stmt
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1);

-- name: tags_by_id_a@tags_info
-- prepare: string
-- clause: in
SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;

-- name: tags_by_id_b@tags_info
-- prepare: string
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);

-- name: decr_tags_by_id@tags_info
-- prepare: string
-- clause: in
UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);

-- name: tags_for_incr@tags_info
-- prepare: string
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);

-- name: incr_tags_by_id@tags_info
-- prepare: string
-- clause: in
UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);
