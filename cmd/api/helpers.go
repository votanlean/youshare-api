package main

import (
	"encoding/json"
	"net/http"
)

type envelop map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelop, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
