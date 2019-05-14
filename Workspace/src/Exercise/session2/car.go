package session2

import "fmt"

type Car struct {
	Parameters
}

func NewCar(name string, fuelAmount, fuelConsumption, fuelTankCapacity, fuelAirConditioning float64) *Car {
	car := &Car{Parameters{name, fuelAmount, fuelConsumption, fuelTankCapacity, fuelAirConditioning}}
	return car
}

func (m Car) ShowInfo() string {
	return fmt.Sprintf("I'm : %s \nFuelAmount : %0.2f \tFuelConsumption : %0.2f \tFuelTankCapacity : %0.2f \tFuelAirConditioning : %0.2f \n",
		m.Name, m.FuelAmount, m.FuelConsumption, m.FuelTankCapacity, m.FuelAirConditioning)
}

func (m *Car) Refuel(fuel float64) {
	m.FuelAmount += fuel
}

func (m *Car) DistnaceNew (float64) {
	return
}

func (m *Car) Driving(distance float64) (status int) {
	neededFuel := (m.FuelConsumption + m.FuelAirConditioning) * distance

	if neededFuel > m.FuelAmount {
		status = 1
		fmt.Println("Warning ...")
		fmt.Println("Fuel is not enough to drive this distance ...")
	} else if neededFuel > (m.FuelAmount - float64(3)) {
		m.FuelAmount -= neededFuel
		status = 2
		fmt.Println("OK ...")
		fmt.Println("Fuel is running out. You should to add feul ...")
	} else {
		m.FuelAmount -= neededFuel
		status = 3
		fmt.Println("OK ...")
	}
	return
}
