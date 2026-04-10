// Package handlers contains the handlers logic
package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/adrr-dev/portfolio/internal/repository"
)

type Handling struct {
	Tmpls     *template.Template
	Fragments *template.Template
}

func NewHandling(tmpls, fragments *template.Template) *Handling {
	newHandling := &Handling{Tmpls: tmpls, Fragments: fragments}
	return newHandling
}

func (h Handling) RootHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpls.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		notice := fmt.Sprintf("could not find template: %e", err)
		http.Error(w, notice, http.StatusNotFound)
		return
	}
}

func (h Handling) HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Fragments.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		notice := fmt.Sprintf("could not find template: %e", err)
		http.Error(w, notice, http.StatusNotFound)
		return
	}
}

func (h Handling) AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Fragments.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		notice := fmt.Sprintf("could not find template: %e", err)
		http.Error(w, notice, http.StatusNotFound)
		return
	}
}

func (h Handling) ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	data := repository.Projects
	err := h.Fragments.ExecuteTemplate(w, "projects.html", data)
	if err != nil {
		notice := fmt.Sprintf("could not find template: %e", err)
		http.Error(w, notice, http.StatusNotFound)
		return
	}
}

func (h Handling) ContactHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Fragments.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		notice := fmt.Sprintf("could not find template: %e", err)
		http.Error(w, notice, http.StatusNotFound)
		return
	}
}

func (h Handling) SkillsHandler(w http.ResponseWriter, r *http.Request) {
	data := repository.Skills
	err := h.Fragments.ExecuteTemplate(w, "skills.html", data)
	if err != nil {
		notice := fmt.Sprintf("could not find template: %e", err)
		http.Error(w, notice, http.StatusNotFound)
		return
	}
}
