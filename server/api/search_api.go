package api

import (
	"net/http"
	"strings"

	"serpgo/search"
)

func makeSearchHandler[T any](searchFn func(string) ([]T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		query := strings.TrimSpace(r.URL.Query().Get("q"))
		if query == "" {
			writeError(w, http.StatusBadRequest, "Missing query parameter: q")
			return
		}

		results, err := searchFn(query)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeJSON(w, http.StatusOK, results)
	}
}

var (
	SearchHandlerDDG   = makeSearchHandler(search.GetDDGResults)
	SearchHandlerBrave = makeSearchHandler(search.GetBraveResults)
)