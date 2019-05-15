package main

import (
	"Exercise2/export"
	"fmt"
)



func main()  {
	listVehicles := []export.Vehicles{
		export.NewCar("ToYoTa",30, 1, 60, 0.9 ),
		export.NewTruck("Hyndai", 50,2, 100,1.6, 0.95),
	}

	for _, v := range listVehicles {
		fmt.Println(v.VehiclesInfo())
	}

	var choice int
	var distance float64
	fmt.Println("=================================================================")
	fmt.Print("Enter number that choice type vehicle(1: car, 2: truct, 0: exit )")
	fmt.Scan(&choice)
	for{
		vehicle := listVehicles[choice - 1]
		fmt.Println(vehicle.VehiclesInfo())
		if(choice ==0 || choice <0 || choice >2 ){
			break
		}
			fmt.Print("Input distance you want to go: ")
			fmt.Scan(&distance)
			status := vehicle.VehiclesDriving(distance)
			if(status==1){
				fmt.Println("You need to input fuel: ")
				var fuelAdd float64
				fmt.Scan(&fuelAdd)
				vehicle.VehiclesRefuel(fuelAdd)
				vehicle.VehiclesDriving(distance)
			}else if(status==2){
				var yes string
				fmt.Println(vehicle.VehiclesInfo())
				fmt.Print("Do do you want to go continue? (y : yes, n: out)")
				fmt.Scan(&yes)
				if(yes=="n"){
					break
				}else if(yes=="y") {
					fmt.Printf("Input distance you want to go to far more: f")
					var DistanceGo float64
					fmt.Scan(&DistanceGo)
					vehicle.VehiclesDriving(DistanceGo)
				}
			}else if status==3{
				fmt.Printf("You went to over capality of Car")
		}
	}
}
