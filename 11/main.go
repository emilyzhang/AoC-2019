package main

import (
	"fmt"

	"github.com/emilyzhang/advent2019/11/pkg/intcode"
	"github.com/emilyzhang/advent2019/11/pkg/robot"
)

func main() {
	source, err := intcode.Read("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	p := intcode.New(source)
	r := robot.New(p)
	r.Run()

	go func() {
		for {
			err := <-r.Errors
			fmt.Println(err)
		}
	}()

	fmt.Println("panels painted:", r.PanelsPainted())
	r.Paint()
	r.PrintHull("solution.png")
	// test()
}
