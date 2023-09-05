-- name: ListAllSaves :many
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

-- name: GetSaveById :many
SELECT * FROM saves
WHERE id = ?;

-- name: GetSaveByNickname :many
SELECT * FROM saves
WHERE nickname = ?;
