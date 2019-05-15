package export

type Vehicles interface {
	VehiclesInfo() string
	VehiclesRefuel(float64)
	VehiclesDriving(float64) int

}

type   DetailVehicles struct {
	VehiclesName string
	FuelAvaliable float64
	FuelConsumption float64
	FuelCapacity float64
	FuelAirConditioner float64
}
