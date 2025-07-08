package main

import (
	"net/http"
)

// Health godoc
//
//	@Summary	Health status
//	@Description
//	@Tags		ops
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	string	"ok"
//	@Failure	500	{object}	error
//	@Security	ApiKeyAuth
//	@Router		/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
