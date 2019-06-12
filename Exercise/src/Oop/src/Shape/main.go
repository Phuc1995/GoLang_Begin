package main

func main()  {
	circle := Circle{radius:10,location:Location{3,6}}
	rectangle := Rectangle{side1:10, side2:15, location:Location{x:9, y:12}}
	//shape := []Shape{&circle,&rectangle}
	circle.ToString()
	rectangle.ToString()
}