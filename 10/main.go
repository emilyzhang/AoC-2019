package main

import (
	"fmt"

	"github.com/emilyzhang/advent2019/10/pkg/asteroid"
)

func main() {
	m, err := asteroid.Read("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(m.Asteroids)
	// fmt.Println(m.MonitorMax())
	// fmt.Println(m.MonitoringStation.X, m.MonitoringStation.Y)
	x := m.Destroy(200)
	fmt.Println(x.X*100 + x.Y)

}
