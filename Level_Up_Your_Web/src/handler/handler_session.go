package handler

import (
	error "error"
	"github.com/julienschmidt/httprouter"
	"net/http"
	s "session"
	u "user"
)

func HandleSessionDestroy(w http.ResponseWriter, r *http.Request) {
	session := s.RequestSession(r)
	if session != nil {
		err := s.GlobalSessionStore.Delete(session)
		if err != nil {
			panic(err)
		}
	}
}

func HandleSessionNew(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	RenderTemplate(w, r, "session/new",
		map[string]interface{}{
			"Next": next,
		})
}

func HandleSessionCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	next := r.FormValue("next")

	user, err := u.FindUser(username, password)

	if err != nil {
		if error.IsValidationError(err){
			RenderTemplate(w, r, "session/new.html",
				map[string]interface{}{
					"Error": err,
					"User": user,
					"Next": next,
				})
			return
		}
		panic(err)
	}

	session := FindOrCreateSession(w,r)

}
