package session2

import (
	"fmt"
)

func main() {

	vehicles := []session2.Vehicles{
		session2.NewCar("Santafe 2019", 30, 0.09, 50, 0.009),
		session2.NewCar("C300 AMG 2019", 40, 0.08, 50, 0.009),
		session2.NewTruck("ISUZU", 60, 0.12, 90, 0.016, 0.95),
		session2.NewTruck("ISUZU", 60, 0.12, 90, 0.016, 0.95),
	}

	fmt.Printf("\t\t\t\t========== PROGRAM MANAGAMENT VEHICLES ============\n")
	for _, v := range vehicles {
		fmt.Println(v.ShowInfo())
	}

	var n int
	var d float64
	for {
		fmt.Print("Which are you choose type vehicle ? (0: Exit)")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		v := vehicles[n-1]
		fmt.Println(v.ShowInfo())
		fmt.Print("Enter distance =  ")
		fmt.Scan(&d)
		status := v.Driving(d)
		if status == 1 {
			fmt.Println("You just add 50 liter fuel")
			v.Refuel(50)
		} else if status == 2 {
			fmt.Println("You just add 20 liter fuel")
			v.Refuel(20)
		} else {
			fmt.Println("You have want add fuel (1: add/ or 0: No) ?")
			var l int
			fmt.Scan(&l)
			if n == 1 {
				fmt.Println("You just add 20 liter fuel")
				v.Refuel(20)
			}
		}
	}
}
