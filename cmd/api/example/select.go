package example

import (
	"context"

	"github.com/lhbelfanti/ditto/database"
	"github.com/lhbelfanti/ditto/log"
)

type SelectAll func(ctx context.Context) ([]DTO, error)

func MakeSelectAll(db database.Connection, collect database.CollectRows[DAO]) SelectAll {
	const query = `
		SELECT id, name, data 
		FROM example_table
	`

	return func(ctx context.Context) ([]DTO, error) {
		rows, err := db.Query(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveExampleData
		}
		defer rows.Close()

		daos, err := collect(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveExampleData
		}

		dtos := make([]DTO, 0, len(daos))
		for _, dao := range daos {
			dtos = append(dtos, DTO{
				ID:   dao.ID,
				Name: dao.Name,
			})
		}

		return dtos, nil
	}
}
