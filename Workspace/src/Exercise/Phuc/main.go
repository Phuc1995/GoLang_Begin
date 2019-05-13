package main

import "fmt"

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

func NewCar(VehiclesName string,FuelQuantity float64, FuelConsumption float64, FuelCapacity float64, FuelAirConditioner float64)  *Car{
	car := &Car{DetailVehicles{VehiclesName,FuelQuantity, FuelConsumption,FuelCapacity, FuelAirConditioner}}
	return car
}

func (m Car) ShowInfo() string {
	return fmt.Sprintf("This is: %s" +
		"\n FuelQuantity: %f" +
		"\n FuelConsumption: %f" +
		"\n FuelCapacity: %f" +
		"\n FuelAirConditioner: %f",m.VehiclesName,m.FuelAvaliable,m.FuelConsumption,m.FuelCapacity,m.FuelAirConditioner)
}

func (m *Car) AddRefuel(fuel float64)  {
	m.FuelAvaliable += fuel
}

func (m *Car) Driving(distance float64)(status int) {
	neededFuel := (m.FuelConsumption + m.FuelAirConditioner) * distance

	if neededFuel > m.FuelAvaliable{
		status = 1
		fmt.Println("Warning ...")
		fmt.Println("Fuel is not enough to drive this distance ...")
	}
}

func main()  {
	m := []Vehicles{
		NewCar("ToYoTa",30, 0.9, 40, 0.09 ),
		NewCar("C300 AMG 2019", 40, 0.08, 50, 0.009),

	}
	for _, v := range Vehicles {
		fmt.Println(v.ShowInfo())
	}
}
