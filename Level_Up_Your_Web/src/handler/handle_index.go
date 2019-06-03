package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"templates"
)

func HandleHome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Display Home Page
	templates.RenderTemplate(w, r, "index/home", nil)
}
