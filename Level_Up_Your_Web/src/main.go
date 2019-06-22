package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"handler"
	"log"
	"middlewaree"
	"session"

	"net/http"
)

type NotFound struct{}

func main()  {
	router := NewRouter()

	router.Handle("GET", "/", handler.HandleHome)
	router.Handle("GET","/register", handler.HandlerUserNew)
	router.Handle("POST", "/register", handler.HandleUserCreate)
	router.Handle("GET", "/login", handler.HandleSessionNew)
	router.Handle("POST", "/login", handler.HandleSessionCreate)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)

	secureRouter := NewRouter()
	secureRouter.Handle("GET", "/sign-out",handler.HandleSessionDestroy)

	middleware := middlewaree.Middleware{}
	middleware.Add(router)
	middleware.Add(http.HandlerFunc(session.RequireLogin))
	middleware.Add(secureRouter)

	fmt.Println("Server running......")
	log.Fatal(http.ListenAndServe(":3000", middleware))
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	return router
}

