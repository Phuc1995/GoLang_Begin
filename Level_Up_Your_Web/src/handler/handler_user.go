package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandlerUserNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	RenderTemplate(w, r, "users/new",nil)
}

func HandleUserCreate(w http.ResponseWriter, r * http.Request,_ httprouter.Params){
	user, err := NewUser(
		r.FormValue("username"),
		r.FormValue("email"),
		r.FormValue("password"),
		)
	//fmt.Print(err)
	if err != nil {
		if IsValidationError(err){
			RenderTemplate(w, r,"users/new", map[string]interface{}{
				"Error" : err.Error(),
				"User" : user,
			} )

			return
		}
		panic(err)
		return
	}
	err = globalUserStore.Save(user)
	if err != nil {
		panic(err)
		return
	}



	http.Redirect(w, r, "/?flash=User+created", http.StatusFound)
}