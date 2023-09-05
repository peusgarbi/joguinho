// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: saves.sql

package sqlite

import (
	"context"
)

const createSave = `-- name: CreateSave :one
INSERT INTO saves (
    nickname, currency
)
VALUES (?, ?)
RETURNING id, nickname, currency
`

type CreateSaveParams struct {
	Nickname string  `db:"nickname" json:"nickname"`
	Currency float64 `db:"currency" json:"currency"`
}

func (q *Queries) CreateSave(ctx context.Context, arg CreateSaveParams) (Safe, error) {
	row := q.db.QueryRowContext(ctx, createSave, arg.Nickname, arg.Currency)
	var i Safe
	err := row.Scan(&i.ID, &i.Nickname, &i.Currency)
	return i, err
}

const getSaveById = `-- name: GetSaveById :many
SELECT id, nickname, currency FROM saves
WHERE id = ?
`

func (q *Queries) GetSaveById(ctx context.Context, id int64) ([]Safe, error) {
	rows, err := q.db.QueryContext(ctx, getSaveById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Safe{}
	for rows.Next() {
		var i Safe
		if err := rows.Scan(&i.ID, &i.Nickname, &i.Currency); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSaveByNickname = `-- name: GetSaveByNickname :many
SELECT id, nickname, currency FROM saves
WHERE nickname = ?
`

func (q *Queries) GetSaveByNickname(ctx context.Context, nickname string) ([]Safe, error) {
	rows, err := q.db.QueryContext(ctx, getSaveByNickname, nickname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Safe{}
	for rows.Next() {
		var i Safe
		if err := rows.Scan(&i.ID, &i.Nickname, &i.Currency); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllSaves = `-- name: ListAllSaves :many
SELECT id, nickname, currency FROM saves
ORDER BY id ASC
`

func (q *Queries) ListAllSaves(ctx context.Context) ([]Safe, error) {
	rows, err := q.db.QueryContext(ctx, listAllSaves)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Safe{}
	for rows.Next() {
		var i Safe
		if err := rows.Scan(&i.ID, &i.Nickname, &i.Currency); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSave = `-- name: UpdateSave :one
UPDATE saves SET
nickname = ?,
currency = ?
WHERE id = ?
RETURNING id, nickname, currency
`

type UpdateSaveParams struct {
	Nickname string  `db:"nickname" json:"nickname"`
	Currency float64 `db:"currency" json:"currency"`
	ID       int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateSave(ctx context.Context, arg UpdateSaveParams) (Safe, error) {
	row := q.db.QueryRowContext(ctx, updateSave, arg.Nickname, arg.Currency, arg.ID)
	var i Safe
	err := row.Scan(&i.ID, &i.Nickname, &i.Currency)
	return i, err
}