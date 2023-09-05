package utils

import (
	"context"
	"joguinho/database/sqlite"
	"joguinho/schema"
	"time"
)

func SaveGame(g *schema.GameContext) error {
	if g.SaveId == 0 {
		return nil
	}

	db, err := sqlite.LocalSqlite.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := sqlite.New(db)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = queries.UpdateSave(ctx, sqlite.UpdateSaveParams{
		ID:       g.SaveId,
		Nickname: g.Nickname,
		Currency: g.Currency,
	})
	if err != nil {
		return err
	}

	return nil
}

func CreateSave(g *schema.GameContext) error {
	db, err := sqlite.LocalSqlite.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := sqlite.New(db)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = queries.CreateSave(ctx, sqlite.CreateSaveParams{
		Nickname: g.Nickname,
		Currency: g.Currency,
	})
	if err != nil {
		return err
	}

	return nil
}
