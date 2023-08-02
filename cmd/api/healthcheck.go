package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	if err := app.writeJSON(w, 200, data, nil); err != nil {
		app.logger.Print(err)
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}
}
