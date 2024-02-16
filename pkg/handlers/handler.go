package handlers

import (
	"morningEat/pkg/config"
	"morningEat/pkg/models"
	"morningEat/pkg/render"
	"net/http"
)

var Repo *Repository // Global variable for Repository

// Repository struct
type Repository struct {
	App *config.AppConfig // AppConfig instance
}

// NewRepo function creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers function sets the global Repo variable
func NewHandlers(r *Repository) {
	Repo = r
}

// Home function handles the home route
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "/home.page.tmpl", &models.TemplateData{})
}

// About function handles the about route
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// Remote IP
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// Render the about page template
	render.RenderTemplate(w, "/about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
