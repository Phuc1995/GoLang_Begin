package handler

import (
	error2 "error"
	"github.com/julienschmidt/httprouter"
	"net/http"
	s "session"
	user2 "user"
)

func HandlerUserNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	RenderTemplate(w, r, "users/new",nil)
}

func HandleUserCreate(w http.ResponseWriter, r * http.Request,_ httprouter.Params){
	user, err := user2.NewUser(
		r.FormValue("username"),
		r.FormValue("email"),
		r.FormValue("password"),
		)
	//fmt.Print(err)
	if err != nil {
		if error2.IsValidationError(err){
			RenderTemplate(w, r,"users/new", map[string]interface{}{
				"Error" : err.Error(),
				"User" : user,
			} )

			return
		}
		panic(err)
		return
	}
	err = user2.GlobalUserStore.Save(user)
	if err != nil {
		panic(err)
		return
	}

	//create a new sesssion
	session := s.NewSession(w)
	session.UserID = user.ID
	err = s.GlobalSessionStore.Save(session)
	if err != nil {
		panic(err)
		return
	}
	http.Redirect(w, r, "/?flash=User+created", http.StatusFound)
}