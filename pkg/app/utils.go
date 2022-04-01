package app

import (
	"encoding/json"
	"net/http"
	"time"
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

func addCookie(w http.ResponseWriter, name string, value string, time_ time.Duration) {
	expire := time.Now().Add(time_)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
		MaxAge:  300,
	}
	http.SetCookie(w, &cookie)
}
