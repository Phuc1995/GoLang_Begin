package main

import (
	"log"
	"net/http"
	 //"templates/template.go"

)
import "Level"
//var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))
//var errorTemplate = `
//<html>
//	<body>
//		<h1>Error rendering template %s</h1>
//		<p>%s</p>
//	</body>
//</html>
//`
//func RenderTemplate(w http.ResponseWriter, request *http.Request, name string, data interface{})  {
//	err := templates.ExecuteTemplate(w, name, data)
//	if err != nil{
//		http.Error(
//			w,
//			fmt.Sprintf(errorTemplate, name, err),
//			http.StatusInternalServerError,
//		)
//	}
//}

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, r, "index/home", nil)
	})

	mux.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))),
		)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
