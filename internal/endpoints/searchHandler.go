package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"

	funcs "groupieTrecker/internal/functions"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		ErrorHandler(w, http.StatusNotFound, errors.New(" "))
		return
	}
	switch r.Method {
	case "GET":
		searchTerm := r.FormValue("searchTerm")
		suggestions, set := funcs.SearchSuggestions(searchTerm)
		responseMap := map[string]interface{}{
			"suggestions": suggestions,
			"set":         set,
		}

		response, err := json.Marshal(responseMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, errors.New(" "))
		return
	}
}
