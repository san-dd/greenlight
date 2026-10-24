package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/san-dd/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanc",
		//Runtime:   102,
		Genres:  []string{"drama", "romance", "war"},
		Version: 1,
	}
	if err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
