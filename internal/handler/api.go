package handler

import (
	"encoding/json"

	"github.com/valyala/fasthttp"

	"github.com/IvanKalug-QA/Go-URL-shorter/internal/forms"
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/model"
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/response"
	"github.com/thanhpk/randstr"
)

var db = model.UrlShortDataBase{
	Name:    "First Database",
	UrlDict: make(map[string]string),
}

func MainPage(ctx *fasthttp.RequestCtx) {
	id, ok := ctx.UserValue("id").(string)
	if ok {
		GetURLById(id, ctx)
		return
	}
	if ctx.IsPost() {
		url := ctx.FormValue("url")
		short := randstr.String(10)
		db.Add(string(url), short)
		subj := response.ShortUrl{
			Status: fasthttp.StatusCreated,
			Url:    short,
		}
		resp, err := json.Marshal(subj)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
		ctx.Response.Header.Set("content-type", "application/json")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Write(resp)

	} else {
		ctx.Response.Header.Set("Content-Type", "text/html")
		ctx.WriteString(forms.UrlForm)
	}
}

func GetURLById(id string, ctx *fasthttp.RequestCtx) {
	original, has := db.GetUrl(id)
	if !has {
		ctx.Error("Такого Url нет!", fasthttp.StatusNotFound)
		return
	}
	subj := response.OriginalUrl{
		Status: fasthttp.StatusOK,
		Url:    original,
	}
	resp, err := json.Marshal(subj)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	ctx.Response.Header.Set("content-type", "application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(resp)
}
