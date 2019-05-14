package main

import (
	"fmt"
)

//Interface
type Vehicles interface {
	VehiclesInfo() string
	VehiclesRefuel(float64)
	VehiclesDriving(float64) int
}

type  DetailVehicles struct {
	VehiclesName string
	FuelAvaliable float64
	FuelConsumption float64
	FuelCapacity float64
	FuelAirConditioner float64
}

//Car
type Car struct {
	DetailVehicles
}

func NewCar(VehiclesName string,FuelAvaliable float64, FuelConsumption float64, FuelCapacity float64, FuelAirConditioner float64)  *Car{
	car := &Car{DetailVehicles{VehiclesName,FuelAvaliable, FuelConsumption,FuelCapacity, FuelAirConditioner}}
	return car
}

func (m Car) VehiclesInfo() string {
	return fmt.Sprintf("This is Car have name: %s  FuelQuantity: %0.0f  FuelConsumption: %0.0f liter/km FuelCapacity: %0.0f FuelAirConditioner: %0.0f ",m.VehiclesName,m.FuelAvaliable,m.FuelConsumption,m.FuelCapacity,m.FuelAirConditioner)

}

func (m *Car) VehiclesRefuel(fuel float64)  {
	m.FuelAvaliable += fuel
}


func (m *Car) VehiclesDriving(distance float64)(status int) {
	fuelNeed := (m.FuelConsumption + m.FuelAirConditioner)*distance
	if fuelNeed > m.FuelAvaliable{
		status =1
		fmt.Println("You cann't go to..............")
		n := fuelNeed -m.FuelAvaliable
		fmt.Printf("You need to add miximum %0.0f(l)",n)
		if(n>m.FuelCapacity){
			status = 3
		}

	}else {
		m.FuelAvaliable -= fuelNeed
		status =2
		fmt.Println("You can go to..............")
		fmt.Println("You used fuel is: ",fuelNeed)
	}

	return
}

func main()  {
	listVehicles := []Vehicles{
		NewCar("ToYoTa",30, 1, 60, 0.9 ),
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
		if(choice == 1 ){
			fmt.Print("Input distance you want to go: ")
			fmt.Scan(&distance)
			status := vehicle.VehiclesDriving(distance)
			if(status==1){
				fmt.Println("You need to input fuel: ")
				var fuelAdd float64
				fmt.Scan(&fuelAdd)
				vehicle.VehiclesRefuel(fuelAdd)
			}
			if(status==2){
				var yes string
				fmt.Println(vehicle.VehiclesInfo())
				fmt.Print("Do do you want to go continue? (y : yes, n: out)")
				fmt.Scan(&yes)
				if(yes=="n"){
					break
				}else if(yes=="y") {
					fmt.Printf("Iput distance you want to go to far more: ")
					var DistanceGo float64
					fmt.Scan(&DistanceGo)
					vehicle.VehiclesDriving(DistanceGo)
				}
			}else if status==2{
				fmt.Printf("You went to over capality of Car")
			}
		}
	}




}
