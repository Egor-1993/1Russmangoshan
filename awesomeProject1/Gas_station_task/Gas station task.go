package main

import "fmt"

type Car struct {
	Name    string
	Fuel    float64
	FuelMax float64
}

func (c *Car) info() {
	fmt.Printf("Car:%s | Fuel:%.1f/%.1f L\n", c.Name, c.Fuel, c.FuelMax)
}

type GasStation struct{}

func (gs *GasStation) Refuel(car *Car, liters float64) {
	if car.Fuel >= car.FuelMax {
		fmt.Printf("Car %s already has a full tank! \n ", car.Name)
		return
	}
	availableSpace := car.FuelMax - car.Fuel
	if liters > availableSpace {
		liters = availableSpace
	}
	car.Fuel += liters
	fmt.Printf("Car %s refueled with %.1f L. Now : %.1f/%.1f L\n", car.Name, liters, car.Fuel, car.FuelMax)
}
func main() {
	car := Car{
		Name:    "Lada",
		Fuel:    20,
		FuelMax: 50,
	}
	station := GasStation{}
	car.info()
	station.Refuel(&car, 40)
	car.info()
}
