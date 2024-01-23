package main

import (
	"net/http"

	"github.com/rostis232/micro-logger/data"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requrstPayload JSONPayload

	_ = app.readJSON(w, r, &requrstPayload)

	event := data.LogEntry{
		Name: requrstPayload.Name,
		Data: requrstPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	
	resp := jsonResp{
		Error: false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}