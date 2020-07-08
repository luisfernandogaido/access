package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/luisfernandogaido/access/model"
)

const token = "yjZFwp3d5ww1h4ja6Uya"

var apps = map[string]bool{"profinanc": true, "consorcio": true}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":4013", nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("token") != "yjZFwp3d5ww1h4ja6Uya" {
		http.Error(w, "não autorizado", http.StatusUnauthorized)
		return
	}
	fim := time.Now()
	ini := fim.Add(-time.Hour)
	var err error
	if r.URL.Query().Get("ini") != "" {
		ini, err = time.Parse("2006-01-02T15:04:05.000-07:00", r.URL.Query().Get("ini")+".000-03:00")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.URL.Query().Get("fim") != "" {
		fim, err = time.Parse("2006-01-02T15:04:05.000-07:00", r.URL.Query().Get("fim")+".000-03:00")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if _, ok := apps[r.URL.Query().Get("app")]; !ok {
		http.Error(w, "app não encontrado", http.StatusUnauthorized)
		return
	}
	access, err := model.List(r.URL.Query().Get("app"), ini, fim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	printJson(w, access)
}

func printJson(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	dec := json.NewEncoder(w)
	dec.SetIndent("", " ")
	return dec.Encode(v)
}
