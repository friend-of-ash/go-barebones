package router

import (
	"pkg/api"
	"pkg/static"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var router *chi.Mux

func GetInitialisedRouter() *chi.Mux {
	initRouter()
	initRoutes()

	return router
}

func initRouter() {
	router = chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
}

func initRoutes() {
	api.RegisterEndpoints(router)
	static.RegisterEndpoints(router)
}
