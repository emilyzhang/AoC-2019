package main

import (
	"fmt"

	"github.com/emilyzhang/advent2019/12/pkg/moon"
)

func main() {
	moons, err := moon.Read("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(moon.StepUntilMatch(moons))

	// for i := 0; i < 1000; i++ {
	// 	moon.Step(moons)
	// }
	// fmt.Println(moon.TotalEnergy(moons))
	// fmt.Println(moons)

	// moons, err := moon.Read("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(moon.StepUntilMatch(moons))
}
