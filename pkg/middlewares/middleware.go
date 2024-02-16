package middlewares

import (
	"fmt"
	"github.com/justinas/nosurf"
	"morningEat/pkg/config"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Hit the page")
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler, app *config.AppConfig) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler, s *config.AppConfig) http.Handler {
	return s.Session.LoadAndSave(next)
}
