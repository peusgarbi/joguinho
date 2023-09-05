package sqlite

import (
	"context"
	"fmt"
)

func GetSaveById(ctx context.Context, id int64) (*Safe, error) {
	db, err := LocalSqlite.Connect()
	if err != nil {
		return nil, fmt.Errorf("erro em 'GetSaveById': não foi possível se conectar ao sqlite:\n%v", err)
	}
	defer db.Close()

	queries := New(db)
	save, err := queries.GetSaveById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("erro em 'GetSaveById': não foi possível realizar a query:\n%v", err)
	}
	if len(save) == 0 {
		return nil, nil
	}

	return &save[0], nil
}

func ListAllSaves(ctx context.Context) (*[]Safe, error) {
	db, err := LocalSqlite.Connect()
	if err != nil {
		return nil, fmt.Errorf("erro em 'ListAllSaves': não foi possível se conectar ao sqlite:\n%v", err)
	}
	defer db.Close()

	queries := New(db)
	saves, err := queries.ListAllSaves(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro em 'ListAllSaves': não foi possível realizar a query:\n%v", err)
	}

	return &saves, nil
}
