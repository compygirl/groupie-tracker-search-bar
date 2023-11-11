package endpoints

import (
	"errors"
	"net/http"
	"strconv"

	data "groupieTrecker/internal/data"
)

func GetMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound, errors.New(" "))
		return
	}
	switch r.Method {
	case "GET":
		indexTemplate := "internal/templates/html_templates/index.html"

		RenderTemplate(w, indexTemplate, data.ArtistsList)
		return
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, errors.New(" "))
		return
	}
}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		ErrorHandler(w, http.StatusNotFound, errors.New(" "))
		return
	}
	switch r.Method {
	case "GET":
		id := r.FormValue("id") // we get id from <a href="/artist?id={{.Id}}"
		if !r.Form.Has("id") {
			ErrorHandler(w, http.StatusBadRequest, errors.New("you hould specify more details"))
			return
		}
		if len(id) > 2 || id[0] == '0' {
			ErrorHandler(w, http.StatusNotFound, errors.New("there is no artist with this ID"))
			return
		}
		indexTemplate := "internal/templates/html_templates/artistPage.html"
		index, err := strconv.Atoi(id)
		if err != nil {
			ErrorHandler(w, http.StatusNotFound, errors.New("there is no artist with this ID"))
			return
		}
		if index <= 0 || index > len(data.ArtistsList) {
			ErrorHandler(w, http.StatusBadRequest, errors.New("there is no artist with this ID"))
			return
		}
		artistInfo := data.ArtistsList[index-1]
		RenderTemplate(w, indexTemplate, artistInfo)
		return
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, errors.New(" "))
		return
	}
}
