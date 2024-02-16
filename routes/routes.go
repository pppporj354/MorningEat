package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"morningEat/pkg/config"
	"morningEat/pkg/handlers"
	"morningEat/pkg/middlewares"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurfWrapper(app))
	mux.Use(SessionLoadWrapper(app))

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

func NoSurfWrapper(app *config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return middlewares.NoSurf(next, app)
	}
}

func SessionLoadWrapper(app *config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return middlewares.SessionLoad(next, app)
	}
}
