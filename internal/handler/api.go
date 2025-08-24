package handler

import (
	"encoding/json"
	"net/http"

	"github.com/IvanKalug-QA/Go-URL-shorter/internal/forms"
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/model"
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/response"
	"github.com/thanhpk/randstr"
)

var db = model.UrlShortDataBase{
	Name:    "First Database",
	UrlDict: make(map[string]string),
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	if id != "" {
		GetUrlById(id, w, r)
		return
	}
	if r.Method == http.MethodPost {
		url := r.FormValue("url")
		short := randstr.String(10)
		db.Add(url, short)
		subj := response.ShortUrl{
			Status: http.StatusCreated,
			Url:    short,
		}
		resp, err := json.Marshal(subj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

	} else {
		w.Write([]byte(forms.UrlForm))
	}
}

func GetUrlById(id string, w http.ResponseWriter, r *http.Request) {
	original, has := db.GetUrl(id)
	if !has {
		http.Error(w, "Такого Url нет!", http.StatusNotFound)
		return
	}
	subj := response.OriginalUrl{
		Status: http.StatusOK,
		Url:    original,
	}
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
