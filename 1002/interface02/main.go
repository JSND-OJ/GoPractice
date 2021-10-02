package main

import "fmt"

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interfaces {
	Accelerate()
	Brake()
	Steer(string)
}

func TryVehicle(Vehicle Vehicle) {
	vehicle.Accelerate()
	Vehicle.Steer("left")
	vehicle.Steer("right")
	Vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		Truck.LoadCargo("test cargo")
	}
}

func main() {
	TryVehicle(Truck("Fnord F180"))
}
