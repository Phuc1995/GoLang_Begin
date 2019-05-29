package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)
const(
	host = "localhost:27017"
	database = "test-db"
	collection = "post"
)

type Post struct {
	Tile string
	Content string
	Test int
}

func main()  {
	info := &mgo.DialInfo{
		Addrs: []string{host},
		Timeout:  60 * time.Second,
		Database: database,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil{
		panic(err)
	}
	//create collection
	col := session.DB(database).C(collection)

	//write data
	err = col.Insert(&Post{"Hello World", "Đây là bài viết hướng dẫn thao tác với MongoDB trong Golang",1},
		&Post{"Thời tiết", "Hôm nay nắng đẹp quá!!!",2})
	if err != nil {
		log.Fatal(err)
	}

	// Đọc dữ liệu
	var posts []Post
	err = col.Find(bson.M{}).All(&posts)
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		log.Println(post)
	}
}