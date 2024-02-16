package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"morningEat/pkg/config"
	"morningEat/pkg/handlers"
	"morningEat/pkg/render"
	"morningEat/routes"
	"net/http"
	"time"
)

const portNumber = ":9000"

var session *scs.SessionManager

func main() {
	var app config.AppConfig
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	render.NewTemplates(&app)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println(fmt.Sprintln("Server is running on port: ", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes.Routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
