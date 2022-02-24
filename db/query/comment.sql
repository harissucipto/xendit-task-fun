-- name: CreateComment :one
INSERT INTO comments (
  content,
  org_name
) VALUES (
  $1, $2
) RETURNING *;

-- name: ListComments :many
SELECT * FROM comments
WHERE org_name = $1
ORDER BY created_at ASC;


-- name: DeleteComment :exec
DELETE FROM comments
WHERE org_name = $1;
