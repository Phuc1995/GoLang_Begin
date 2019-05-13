package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Student map[int]string

func (s Student) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if len(s[id]) >0{
		fmt.Fprintf(w,  "Tên sinh viên có mã %d là: %s\n", id, s[id])
		}else {
		fmt.Fprintf(w, "Không có sinh viên có mã %d!\n", id)
	}
}

func main()  {
	sv := Student{1: "Mai", 2: "Lan", 3: "Cúc", 4: "Trúc"}
	http.ListenAndServe("localhost:8080", sv)
}