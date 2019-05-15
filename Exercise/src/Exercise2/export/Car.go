package export

import "fmt"

//Car
type Car struct {
	DetailVehicles
}

func (m *Car) abc(a int) int {
	return a
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
		fmt.Println()
		if(n>m.FuelCapacity){
			status = 3
		}

	}else {
		m.FuelAvaliable -= fuelNeed
		status =2
		fmt.Println("You can go to..............")
		fmt.Println("You used fuel is: ",fuelNeed)
	}

	return status
}

