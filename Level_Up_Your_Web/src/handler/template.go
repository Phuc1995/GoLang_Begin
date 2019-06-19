package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var layoutFuncs = template.FuncMap{
	"yield": func() (string, error) {
		return "", fmt.Errorf("yield called inappropriately")
	},
}

var errorTemplate = `
<html>
	<body>
		<h1>Error rendering template %s</h1>
		<p>%s</p>
	</body>
</html>
`


func RenderTemplate(w http.ResponseWriter, request *http.Request, name string, data interface{})  {
	wd, err1 := os.Getwd()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("wd:***********",wd)
	//string1 := `\` + "templates" + `\` + "**" + `\` + "*.html"
	var templates = template.Must(template.New("t").ParseGlob(wd +"/templates/**/*.html"))
	//var templates = template.Must(template.New("t").ParseGlob(wd  + string1))
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err
		},
	}

	// Template
	//tpl, err1 := template.ParseFiles(wd + "/templates/index.html")

	var layout = template.Must(
		template.
			New("layout.html").
			Funcs(layoutFuncs).
			ParseFiles(wd+ "/templates/layout.html"),
	)

	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w, data)

	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}