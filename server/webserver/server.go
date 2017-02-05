package webserver

import (
	"../config"
	"./handlers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func logger(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		logger := ctx.Logger()
		logger.Printf(ctx.RemoteAddr().String())
		next(ctx)
	}
}

func applyMiddleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	newHandler := logger(handler)

	return newHandler
}

func registerRoutes(router *fasthttprouter.Router) {

	router.GET("/soundlist", handlers.SoundList)
	router.PUT("/upload", handlers.FileUpload)

	router.ServeFiles("/static/*filepath", "./static")
	router.ServeFiles("/sounds/*filepath", "./sounds")

	router.NotFound = func(ctx *fasthttp.RequestCtx) {
		fasthttp.ServeFile(ctx, "./index.html")
	}
}

func Start() {
	router := fasthttprouter.New()

	registerRoutes(router)

	// apply our middleware
	handlers := applyMiddleware(router.Handler)

	// start web server
	log.Fatal(fasthttp.ListenAndServe(config.Config.ServerAddr, handlers))
}
