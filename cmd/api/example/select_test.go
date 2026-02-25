package example

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/lhbelfanti/ditto/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMakeSelectAll(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockDB := new(database.MockPostgresConnection)
		mockRows := new(database.MockPgxRows)

		mockDB.On("Query", ctx, mock.Anything, mock.Anything).Return(mockRows, nil)
		mockRows.On("Close").Return()

		collectRows := func(rows pgx.Rows) ([]DAO, error) {
			return []DAO{
				{ID: 1, Name: "Test Name", Data: "Test Data"},
			}, nil
		}

		selectAll := MakeSelectAll(mockDB, collectRows)
		res, err := selectAll(ctx)

		require.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, "Test Name", res[0].Name)
		assert.Equal(t, 1, res[0].ID)

		mockDB.AssertExpectations(t)
		mockRows.AssertExpectations(t)
	})

	t.Run("query failure", func(t *testing.T) {
		mockDB := new(database.MockPostgresConnection)

		mockDB.On("Query", ctx, mock.Anything, mock.Anything).Return(new(database.MockPgxRows), errors.New("db connection error"))

		selectAll := MakeSelectAll(mockDB, nil)
		res, err := selectAll(ctx)

		require.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, FailedToRetrieveExampleData, err)

		mockDB.AssertExpectations(t)
	})
}
