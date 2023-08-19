----------------------
-- default sql ddl --
----------------------

-- name: user_info_a
-- prepare: stmt
SELECT `username` FROM @user WHERE username=?

-- name: user_info_b
SELECT `username`, `nickname` FROM @user WHERE username=?

-- name: user_info_c
-- prepare: raw
SELECT "username" FROM @user WHERE username=?

-- name: user_info_d
-- prepare: string
SELECT "username", "nickname" FROM @user WHERE username=?

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
