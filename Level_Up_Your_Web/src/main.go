package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"handler"
	"log"
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

	middleware := handler.Middleware{}
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