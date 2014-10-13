package minutes

import (
	"appengine"
	"encoding/json"
	"net/http"

	"model/minutes"
)

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	message := r.FormValue("title")

	_, err := minutes.Saveas(c, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	minutes_list, err := minutes.AscList(c)

	js, err := json.Marshal(minutes_list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
