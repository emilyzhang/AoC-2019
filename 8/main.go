package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	shittyimage "./pkg/image"
)

func Read(filename string) ([]int, error) {
	var data []int

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}
	d := strings.TrimSpace(string(f))
	for _, n := range d {
		num, err := strconv.Atoi(string(n))
		if err != nil {
			return data, err
		}
		data = append(data, num)
	}
	return data, nil

}

func main() {
	data, err := Read("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	i := shittyimage.New(data, 25, 6)

	var fewestzeroLayer shittyimage.Layer
	fewest := 1000000000
	for _, layer := range i.Layers {
		if layer.ZeroDigits < fewest {
			fewestzeroLayer = layer
			fewest = layer.ZeroDigits
		}
	}
	fmt.Println(fewestzeroLayer.OneDigits * fewestzeroLayer.TwoDigits)

	i.Decode()
	i.Write("solution.png")

	// data, err := Read("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// i := shittyimage.Decode(data, 3, 2)
	// fmt.Println(i.Layers)

	// var maxzeroLayer shittyimage.Layer
	// maxzeros := 0
	// for _, layer := range i.Layers {
	// 	// println(maxzeros)
	// 	if layer.ZeroDigits > maxzeros {
	// 		maxzeroLayer = layer
	// 		maxzeros = layer.ZeroDigits
	// 	}
	// }
	// fmt.Println(maxzeroLayer.OneDigits * maxzeroLayer.TwoDigits)
}
