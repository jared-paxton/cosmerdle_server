package app

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrapKey string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrapKey] = data

	var js []byte
	var err error
	if app.config.Env == "development" {
		js, err = json.MarshalIndent(wrapper, "", "\t")
	} else {
		js, err = json.Marshal(wrapper)
	}

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
