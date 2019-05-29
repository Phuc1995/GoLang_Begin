package main

import (
	"fmt"
	"net/http"
)

/*type liteHandler struct {
	message string
}

func (m *liteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, m.message)
}

func main()  {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("html"))
	mux.Handle("/",fs)
	hdl := &liteHandler{"Welcome Golang"}
	mux.Handle("/welcome", hdl)
	http.ListenAndServe(":8080",mux)
}*/

func liteHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Chào mừng đến lập trình Go cho web!")
}

func main()  {
/*	mux := http.NewServeMux()
	//Ep kieu liteHandler ve http.HandleFunc
	hdl := http.HandlerFunc(liteHandler)
	mux.Handle("/welcome", hdl)
	http.ListenAndServe(":8080",mux)*/

	http.HandleFunc("/welcome",liteHandler)
	http.ListenAndServe(":8080",nil)

}