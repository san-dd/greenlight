package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//1st appoch
	// fmt.Fprintln(w, "Status:available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)
	// //2nd approch
	// js := `{"status": "available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)
	//3 rd approch
	// data := map[string]any{
	// 	"status":      "available",
	// 	"environment": app.config.env,
	// 	"version":     version,
	// 	"helo":        []int{2, 3, 4},
	// 	"time":        time.Now(),
	// 	"string":      []byte(`{"status": "available", "environment": 2, "version": 2}`),
	// }
	// js, err := json.Marshal(data)
	// if err != nil {
	// 	app.logger.Error((err.Error()))
	// 	http.Error(w, "the server encountera problem and couldnot process your request", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// //2nd approc
	// //w.Write([]byte(js))
	// js = append(js, '\n')
	// w.Write(js)
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	if err := app.writeJSON(w, http.StatusOK, env, nil); err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}
