package main

import "fmt"

type vehicle struct {
	doors string
	colour string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck := truck{
		vehicle: vehicle{
			doors: "four",
			colour: "red",
		},
		fourWheels: true,
	}

	sedan := sedan{
		vehicle: vehicle{
			doors: "four",
			colour: "blue",
		},
		luxury: false,
	}

	car := vehicle{
		doors: "nine",
		colour: "rain",
	}

	fmt.Println(truck.colour, truck.doors, truck.fourWheels)
	fmt.Println(sedan.colour, sedan.doors, sedan.luxury)
	fmt.Println(car.colour, car.doors)

	fmt.Println(truck)
	fmt.Println(sedan)
	fmt.Println(car)

}
