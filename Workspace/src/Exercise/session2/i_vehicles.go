package session2

type Vehicles interface {
	ShowInfo() string
	Refuel(float64)
	Driving(float64) int
}

type Parameters struct {
	Name                string
	FuelAmount          float64
	FuelConsumption     float64
	FuelTankCapacity    float64
	FuelAirConditioning float64
}
