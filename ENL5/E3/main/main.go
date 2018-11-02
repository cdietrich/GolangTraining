package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	s := sedan{
		luxury: false,
		vehicle: vehicle{
			color: "red",
			doors: 2,
		},
	}
	fmt.Println(s.color, s.doors, s.luxury)
	t := truck{
		fourWheel: true,
		vehicle: vehicle{
			color: "black",
			doors: 6,
		},
	}
	fmt.Println(t.color, t.doors, t.fourWheel)
}
