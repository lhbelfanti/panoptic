package example

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lhbelfanti/ditto/http/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectAllHandlerV1(t *testing.T) {
	t.Run("success returns 200 OK", func(t *testing.T) {
		mockSelectAll := func(ctx context.Context) ([]DTO, error) {
			return []DTO{
				{ID: 1, Name: "Test Output"},
			}, nil
		}

		req, err := http.NewRequest(http.MethodGet, "/example/v1", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := SelectAllHandlerV1(mockSelectAll)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var res response.DTO
		err = json.Unmarshal(rr.Body.Bytes(), &res)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "OK", res.Message)
		// Usually requires casting struct logic for assert.Equal on inner data payload but validates standard format.
	})

	t.Run("error returns 500 Internal Server Error", func(t *testing.T) {
		mockSelectAll := func(ctx context.Context) ([]DTO, error) {
			return nil, FailedToRetrieveExampleData
		}

		req, err := http.NewRequest(http.MethodGet, "/example/v1", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := SelectAllHandlerV1(mockSelectAll)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var res response.DTO
		err = json.Unmarshal(rr.Body.Bytes(), &res)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, InternalServerErrorMessage, res.Message)
	})
}
