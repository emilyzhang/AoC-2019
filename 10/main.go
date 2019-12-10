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

	fmt.Println(m)

}
