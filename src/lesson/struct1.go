package main

import (
	"fmt"
)

type Part struct {
	Name        string
	Description string
	NeedsSpare  bool
}
type Parts []Part

type Bicycle struct {
	Size string
	Parts
}

func main() {
	var (
		RoadBikeParts = Parts{
			{"chain", "10-speed", true},
			{"tire_size", "23", true},
			{"tape_color", "red", true},
		}

		MountainBikeParts = Parts{
			{"chain", "10-speed", true},
			{"tire_size", "2.1", true},
			{"front_shock", "Manitou", false},
			{"rear_shock", "Fox", true},
		}

		RecumbentBikeParts = Parts{
			{"chain", "9-speed", true},
			{"tire_size", "28", true},
			{"flag", "tall and orange", true},
		}
	)
	roadBike := Bicycle{Size: "L", Parts: RoadBikeParts}
	montainBike := Bicycle{Size: "L", Parts: MountainBikeParts}
	RecumbentBike := Bicycle{Size: "L", Parts: RecumbentBikeParts}

	fmt.Println(roadBike)
	fmt.Println(montainBike)
	fmt.Println(RecumbentBike)

}
