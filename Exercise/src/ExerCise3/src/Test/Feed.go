package main

import "fmt"

type Post struct {
	body string
	publishDate int64
	next *Post
}

type Feed struct {
	length int
	start *Post
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0{
		f.start = newPost
	} else {
		currentPost := f.start
		for currentPost.next != nil{
			currentPost = currentPost.next
		}
		currentPost.next = newPost
	}
	f.length++
}

func main()  {
	f := &Feed{}
	p1 := Post{
		body:"Lorem ipsum",
	}
	f.Append(&p1)
	fmt.Println(&p1)
	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)

	p2 := Post{
		body: "Dolor sit amet",
	}
	f.Append(&p2)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)
	fmt.Printf("Second: %v\n", f.start.next)

}