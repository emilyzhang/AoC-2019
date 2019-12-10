package asteroid

import (
	"bufio"
	"os"
	"strings"
)

type Map struct {
	Asteroids []Asteroid
}

type Asteroid struct {
	X int
	Y int
}

// Read takes in a filename and returns the asteroid map
// specified in the file.
func Read(filename string) ([]Asteroid, error) {
	m := make([]Asteroid, 0)
	file, err := os.Open(filename)
	if err != nil {
		return m, err
	}
	defer file.Close()

	var data string
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = strings.TrimSpace(scanner.Text())
		x := 0
		for _, d := range data {
			if (string(d)) == "#" {
				m = append(m, Asteroid{X: x, Y: y})
			}
			x++
		}
		y++
	}
	return m, nil
}
