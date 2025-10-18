package handler

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/valyala/fasthttp"

	DB_version2 "github.com/IvanKalug-QA/Go-URL-shorter/internal/config/db"
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/forms"
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/model"
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/response"
	"github.com/thanhpk/randstr"
)

var db = model.UrlShortDataBase{
	Name:     "First Database",
	UrlDict:  make(map[string]string),
	FilePath: "shorturl.json",
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
		resp, err := json.MarshalIndent(subj, "", " ")
		resp = append(resp, []byte(",")...)
		file, errr := os.OpenFile(db.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if errr != nil {
			ctx.Error(errr.Error(), fasthttp.StatusInternalServerError)
			return
		}
		defer file.Close()
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
		_, file_err := file.Write(resp)
		if file_err != nil {
			ctx.Error(file_err.Error(), fasthttp.StatusInternalServerError)
			return
		}
		resp = resp[:len(resp)-1]
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

func Ping(ct *fasthttp.RequestCtx) {
	Db := DB_version2.GetDB()
	defer Db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := Db.PingContext(ctx); err != nil {
		ct.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	ct.SetStatusCode(fasthttp.StatusOK)
}
