package main

import (
	"fmt"

	"github.com/emilyzhang/advent2019/9/pkg/intcode"
)

func main() {
	// source, err := intcode.Read("input.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// p := intcode.New(source)

	testjump()
}

func testjump() {
	source, err := intcode.Read("testjump.txt")
	if err != nil {
		fmt.Println(err)
	}
	p := intcode.New(source)

	for !p.Halted() {
		err = p.Run()
		if err != nil {
			fmt.Println(err)
			break
		}
		if p.RequiresInput() {
			p.Input(7)
		}
		if p.HasOutput() {
			fmt.Println(p.Output())
		}
	}
}
