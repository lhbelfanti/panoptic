package example

import (
	"encoding/json"
	"net/http"

	"github.com/lhbelfanti/ditto/http/response"
)

func SelectAllHandlerV1(selectAll SelectAll) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		w.Header().Set("Content-Type", "application/json")

		examples, err := selectAll(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(response.DTO{
				Code:    http.StatusInternalServerError,
				Message: InternalServerErrorMessage,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(response.DTO{
			Code:    http.StatusOK,
			Message: "OK",
			Data:    examples,
		})
	}
}
