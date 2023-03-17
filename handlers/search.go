package handlers

import (
	"fmt"
	"lib_man/models"
	"lib_man/utility"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	status := utility.AuthToken(w, r)
	if !status {
		return
	}
	t := r.URL.Query().Get("type")
	if t == "" {
		utility.Respond(http.StatusBadRequest, "missing query parameter", &w, false)
		return
	}
	if t == "name" {
		searchBookName(w, r)
	}
	if t == "cat" {
		searchBookCat(w, r)
	}
}

func searchBookCat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("cat")
	if query == "" {
		utility.Respond(http.StatusBadRequest, "missing query parameter", &w, false)
		return
	}
	books, err := models.GetBookByCategory(query)
	if err != nil {
		utility.Respond(http.StatusNotFound, err.Error(), &w, false)
		return
	}
	utility.RespondBooks(books, &w, true)
	if err != nil {
		utility.Respond(500, "some internal error", &w, false)
		return
	}
}

func searchBookName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")
	if query == "" {
		utility.Respond(http.StatusBadRequest, "missing query parameter", &w, false)
		return
	}
	books, err := models.GetBookByName(query)
	if err != nil {
		utility.Respond(http.StatusNotFound, err.Error(), &w, false)
		return
	}
	utility.RespondBooks(books, &w, true)
	if err != nil {
		fmt.Println(err.Error())
		utility.Respond(500, "some internal error", &w, false)
	}
}
