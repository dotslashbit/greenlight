package main

import (
	"net/http"
	"time"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system-info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	// js := `{"status":"available","environment":%q,"version":%q}`
	// js = fmt.Sprintf(js, app.config.env, version)
	//
	time.Sleep(4 * time.Second)

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)

	}

}
