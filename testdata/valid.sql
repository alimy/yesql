-- name: simple
-- raw: 1
SELECT * FROM simple;

-- name: multiline@namespace
SELECT *
FROM multiline
WHERE line = 42;


-- name: comments
-- yoyo

SELECT *
-- inline
FROM comments;
