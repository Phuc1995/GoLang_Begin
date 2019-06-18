package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"handler"
	"log"
<<<<<<< HEAD
	"middlewaree"
=======

	"middlerwaree"
>>>>>>> a8c6af148516d9a83e29d78934d9f34a3a1501ae
	"net/http"
)

type NotFound struct{}

func main()  {
	router := NewRouter()

	router.Handle("GET", "/", handler.HandleHome)
	router.Handle("GET","/register", handler.HandlerUserNew)
	router.Handle("POST", "/register", handler.HandleUserCreate)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)

<<<<<<< HEAD
	middleware := middlewaree.Middleware{}
=======
	middleware := middlerwaree.Middleware{}
>>>>>>> a8c6af148516d9a83e29d78934d9f34a3a1501ae
	middleware.Add(router)


	fmt.Println("Server running......")
	log.Fatal(http.ListenAndServe(":3000", middleware))
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	return router
}

/*var layout = template.Must(
	template.
		New("layout.html").
		Funcs(layoutFuncs).
		Parse(),
)*/