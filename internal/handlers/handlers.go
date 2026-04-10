// Package handlers contains the handlers logic
package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type Handling struct {
	Tmpls     *template.Template
	Fragments *template.Template
}

func NewHandling(tmpls, fragments *template.Template) *Handling {
	newHandling := &Handling{Tmpls: tmpls, Fragments: fragments}
	return newHandling
}

func (h Handling) RootHandle(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpls.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		notice := fmt.Sprintf("could not find template: %e", err)
		http.Error(w, notice, http.StatusNotFound)
		return
	}
}
