package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//panic("paniced")
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	// Add a 4 second delay.
	//time.Sleep(15 * time.Second)
	if err := app.writeJSON(w, http.StatusOK, env, nil); err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}
