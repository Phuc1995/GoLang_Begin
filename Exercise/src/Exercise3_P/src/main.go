package main

import (
	"Exercise3_P/src/session2"
	"fmt"
	"math/rand"
	"time"
)

func Random(n int) (ck, cr, sh int) {
	ck = rand.Intn(n)
	cr = rand.Intn(n - ck)
	sh = n - ck - cr
	return ck, cr, sh
}

func main() {
	var qty int
	var animal []session2.Animal
	var lc int
	fmt.Println("Enter quantity animal")
	fmt.Scan(&qty)

	ck, cr, sh := Random(qty)
	fmt.Printf("Created Animal Famer : Ga %d \tBo %d \tCuu %d \n", ck, cr, sh)

	for ck > 0 {
		animal = append(animal, session2.NewChicken("GA 500K", 1))
		ck--
	}

	for cr > 0 {
		animal = append(animal, session2.NewCow("Bo To Cu Chi", 2, 5))
		cr--
	}

	for sh > 0 {
		animal = append(animal, session2.NewSheep("Dooly", 3, 4))
		sh--
	}

	fmt.Println("================= List animal of Famer ====================")
	for _, ani := range animal {
		ani.Statistic()
	}

	for {
		go func() { // After 5 seconds animal will give birth
			for {
				time.Sleep(5 * time.Second)
				for _, ani := range animal {
					ani.GiveBirth(5)
				}
				for _, ani := range animal {
					ani.Statistic()
				}
			}
		}()

		fmt.Println("Choolse 1 func (0 : Exit) = ")
		fmt.Scan(&lc)

		if lc == 0 {
			break
		}
		switch lc {
		case 1:
			for _, ani := range animal {
				ani.Shout()
			}
		case 2:
			for _, ani := range animal {
				ani.GiveBirth(5)

			}
		case 3:
			for _, ani := range animal {
				ani.GiveMilk(3)
			}
		case 4:
			for _, ani := range animal {
				ani.Statistic()
			}
		}

	}

}
