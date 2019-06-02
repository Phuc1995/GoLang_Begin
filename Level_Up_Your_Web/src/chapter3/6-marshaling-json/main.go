package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	example1()
	example2()
	example3()
}

func example1() {
	type Article struct {
		Name       string
		AuthorName string
		draft      bool
	}
	article := Article{
		Name:       "JSON in Go",
		AuthorName: "Mal Curtis",
		draft:      true,
	}
	data, err := json.Marshal(article)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(data))

}

func example2() {
	type Article struct {
		Name       string 
		AuthorName string `json:"author_name"`
		draft      bool
	}
	article := Article{
		Name:       "JSON in Go",
		AuthorName: "Mal Curtis",
		draft:      true,
	}
	data, err := json.MarshalIndent(article, " ", " . ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func example3() {
	type Article struct {
		Name string `json:"name"`
	}
	type ArticleCollection struct {
		Articles []Article `json:"articles"`
		Total    int       `json:"total"`
	}
	p1 := Article{Name: "JSON in go"}
	p2 := Article{Name: "Marshaling is easy"}

	articles := []Article{p1, p2}
	colection := ArticleCollection{
		Articles: articles,
		Total: len(articles),
	}
	data, err := json.MarshalIndent(colection, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}