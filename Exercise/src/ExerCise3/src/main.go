package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// type AnimalFarm interface{
// 	Shout()
// 	GiveBirth()
// 	GiveMilk()
// 	Statictis
// }
// type Animal struct{
// 	Chicken
// 	Sheep
// 	Cow
// }
type Chicken struct {
	quantity int
	birds    int
}
type Sheep struct {
	quantity int
	birds    int
	milk     float32
}
type Cow struct {
	quantity int
	birds    int
	milk     float32
}

func (C *Chicken) GiveBirth() {
	for i := 1; i <= C.quantity; i++ {
		C.birds += rand.Intn(5) + 1
	}
}
func (S *Sheep) GiveBirth() {
	for i := 1; i <= S.quantity; i++ {
		S.birds += rand.Intn(2) + 1
	}
}
func (C *Cow) GiveBirth() {
	for i := 1; i <= C.quantity; i++ {
		C.birds += rand.Intn(2) + 1
	}
}
func (S *Sheep) GiveMilk() {
	for i := 1; i <= S.quantity; i++ {
		S.milk += float32(rand.Intn(2) + 1)
	}
}
func (C *Cow) GiveMilk() {
	for i := 1; i <= C.quantity; i++ {
		C.milk += float32(rand.Intn(5) + 1)
	}
}
func (C *Chicken) Shout() {
	for i := 1; i <= C.quantity; i++ {
		fmt.Println("Chiken ", i, " shouted.")
	}
}
func (S *Sheep) Shout() {
	for i := 1; i <= S.quantity; i++ {
		fmt.Println("Sheep ", i, " shouted.")
	}
}
func (C *Cow) Shout() {
	for i := 1; i <= C.quantity; i++ {
		fmt.Println("Cow ", i, " shouted.")
	}
}
func main() {
	arg := os.Args[1:2]
	var intt int
	for _, v := range arg {
		intt, _ = strconv.Atoi(v)
	}
	scanner := bufio.NewScanner(os.Stdin)
	Chicken := &Chicken{}
	Sheep := &Sheep{}
	Cow := &Cow{}
	rand.Seed(time.Now().UnixNano())
	Chicken.quantity = rand.Intn(intt)
	Sheep.quantity = rand.Intn(intt - Chicken.quantity)
	Cow.quantity = intt - Chicken.quantity - Sheep.quantity
	fmt.Println("chicken ", Chicken.quantity, " Sheep ", Sheep.quantity, " Cow ", Cow.quantity)
	for {
		scanner.Scan()
		if scanner.Text() == "1" {
			Chicken.GiveBirth()
			Cow.GiveBirth()
			Sheep.GiveBirth()
		} else if scanner.Text() == "2" {
			Cow.GiveMilk()
			Sheep.GiveMilk()
		} else if scanner.Text() == "3" {
			Chicken.Shout()
			Cow.Shout()
			Sheep.Shout()
		} else if scanner.Text() == "4" {
			fmt.Println("Chicken birds: ", Chicken.birds, "Sheep birds: ",
				Sheep.birds, "Cow birds: ", Cow.birds)
			fmt.Println("Sheep milk: ", Sheep.milk, "Cow milk: ",
				Cow.milk)
		} else if scanner.Text() == "5" {
			break
		} else {
			fmt.Println("loai")
		}
	}

}
