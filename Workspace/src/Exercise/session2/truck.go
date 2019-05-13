package session2

import "fmt"

type Truck struct {
	Parameters
	TankConsumption float64
}

func NewTruck(name string, fuelAmount, fuelConsumption, fuelTankCapacity, fuelAirConditioning, tankConsumption float64) *Truck {
	truck := &Truck{Parameters{name, fuelAmount, fuelConsumption, fuelTankCapacity, fuelAirConditioning}, tankConsumption}
	return truck
}

func (m Truck) ShowInfo() string {
	return fmt.Sprintf("I'm : %s \nFuelAmount : %0.2f \tFuelConsumption : %0.2f \tFuelTankCapacity : %0.2f \tFuelAirConditioning : %0.2f"+
		"\tTankConsumption : %0.2f \n", m.Name, m.FuelAmount, m.FuelConsumption, m.FuelTankCapacity, m.FuelAirConditioning, m.TankConsumption)
}

func (m *Truck) Refuel(fuel float64) {
	m.FuelAmount += (fuel * m.TankConsumption)
}

func (m *Truck) Driving(distance float64) (status int) {
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
		fmt.Println("OK")
	}
	return
}
