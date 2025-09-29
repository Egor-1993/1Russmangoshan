package main

import "fmt"

type Car struct {
	Name    string
	Fuel    float64
	FuelMax float64
}

func (c *Car) Info() {
	fmt.Printf("Car:%s | Fuel:%.1f/%.1f L\n", c.Name, c.Fuel, c.FuelMax)
}

type GasStation struct {
	FuelStock float64
	Active    bool
}

func (gs *GasStation) Refuel(car *Car, liters float64) {
	if !gs.Active {
		fmt.Println("Station is not active! Cannot refuel.")
		return
	}
	if car.Fuel >= car.FuelMax {
		fmt.Printf("Car %s tank is full! Cannot go to the station.\n", car.Name)
		return
	}
	availableSpace := car.FuelMax - car.Fuel
	if liters > availableSpace {
		liters = availableSpace
	}
	if liters > gs.FuelStock {
		fmt.Printf("Station has only %.1f L left. Refueling with available fuel.\n ", gs.FuelStock)
		liters = gs.FuelStock
	}
	car.Fuel += liters
	gs.FuelStock -= liters
	fmt.Printf("Car %s refueled with %.1f L. Now: %.1f/%.1f L\n", car.Name, liters, car.Fuel, car.FuelMax)
}
func main() {
	car := Car{
		Name:    "Uaz",
		Fuel:    5,
		FuelMax: 50,
	}
	station := GasStation{
		FuelStock: 100,
		Active:    true,
	}
	car.Info()
	station.Refuel(&car, 40)
	car.Info()
	station.Refuel(&car, 10)
}
