package web

import (
	"log"
	"nlp/tler"
	"nlp/web/views"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func Run() {
	router := fasthttprouter.New()
	router.GET("/", index)
	router.POST("/", indexPOST)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

func index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	ctx.SetContentType("html")
	views.WriteIndex(ctx, " ")
}

func indexPOST(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	var selects string = string(ctx.FormValue("kan"))
	var textarea string = string(ctx.FormValue("id"))

	if selects == "0" { // to english
		word, tler, gram, tag := tler.Indonesia(textarea)
		words := strings.Join(word, " ")
		tlers := strings.Join(tler, " ")
		grams := strings.Join(gram, " ")
		tags := strings.Join(tag, " ")
		ctx.SetContentType("html")
		views.WriteKam(ctx, words, tlers, grams, tags)
	} else {
		word, tler, gram, tag := tler.English(textarea)
		words := strings.Join(word, " ")
		tlers := strings.Join(tler, " ")
		grams := strings.Join(gram, " ")
		tags := strings.Join(tag, " ")
		ctx.SetContentType("html")
		views.WriteKam(ctx, words, tlers, grams, tags)
	}

}
