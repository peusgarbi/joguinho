-- name: GetAllSaves :many
SELECT * FROM saves
ORDER BY id ASC;

-- name: CreateSave :one
INSERT INTO saves (
    nickname, currency
)
VALUES (?, ?)
RETURNING *;

-- name: UpdateSave :one
UPDATE saves SET
nickname = ?,
currency = ?
WHERE id = ?
RETURNING *;

-- name: GetSave :one
SELECT * FROM saves
WHERE id = ?
LIMIT 1;
