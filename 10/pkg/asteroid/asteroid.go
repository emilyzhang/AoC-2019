package asteroid

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Map struct {
	Asteroids         []Asteroid
	MonitoringStation Asteroid
}

type Asteroid struct {
	X              int
	Y              int
	angles         map[float64]bool
	otherAsteroids map[float64][]Asteroid
	seen           int
}

func (a Asteroid) Angle(b Asteroid) float64 {
	angle := math.Atan2(float64(b.Y-a.Y), float64(b.X-a.X))
	return angle
}

func (a Asteroid) Distance(b Asteroid) float64 {
	distance := math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
	return distance
}

func (m *Map) MonitorMax() int {
	maxAsteroids := 0
	for _, a := range m.Asteroids {
		for _, b := range m.Asteroids {
			if !(a.X == b.X && a.Y == b.Y) {
				angle := a.Angle(b)
				if angle < -math.Pi/2 {
					angle += 2 * math.Pi
				}
				if !a.angles[angle] {
					a.seen++
					fmt.Printf("(%v,%v) able to view asteroid (%v,%v) ", a.X, a.Y, b.X, b.Y)
				}
				a.angles[angle] = true
				fmt.Println(angle, "a", a.X, a.Y, "b", b.X, b.Y)
			}
		}
		if a.seen > maxAsteroids {
			m.MonitoringStation = a
			maxAsteroids = a.seen
		}
		fmt.Println()
	}
	return maxAsteroids
}

func (m *Map) Destroy(vaporize int) Asteroid {
	maxAsteroids := 0
	for _, a := range m.Asteroids {
		for _, b := range m.Asteroids {
			if !(a.X == b.X && a.Y == b.Y) {
				angle := a.Angle(b)
				if angle < -math.Pi/2 {
					angle += 2 * math.Pi
				}
				if _, ok := a.angles[angle]; ok {
					a.otherAsteroids[angle] = append(a.otherAsteroids[angle], b)
				} else {
					a.otherAsteroids[angle] = []Asteroid{b}
				}
				// fmt.Printf("(%v,%v) added asteroid (%v,%v) ", a.X, a.Y, b.X, b.Y)
				// fmt.Println(angle, "a", a.X, a.Y, "b", b.X, b.Y)
				if !a.angles[angle] {
					a.seen++
				}
				a.angles[angle] = true
			}
		}
		if a.seen > maxAsteroids {
			m.MonitoringStation = a
			maxAsteroids = a.seen
		}
	}

	keys := make([]float64, len(m.MonitoringStation.otherAsteroids))
	i := 0
	for k := range m.MonitoringStation.otherAsteroids {
		keys[i] = k
		i++
	}
	sort.Float64s(keys)
	for _, k := range keys {
		asteroids := m.MonitoringStation.otherAsteroids[k]
		sort.Slice(asteroids,
			func(i, j int) bool {
				return m.MonitoringStation.Distance(asteroids[i]) < m.MonitoringStation.Distance(asteroids[j])
			})
	}

	return m.MonitoringStation.otherAsteroids[keys[vaporize-1]][0]

	// destroyed := 0
	// var x Asteroid
	// for true {
	// 	for _, k := range keys {
	// 		asteroids, ok := m.MonitoringStation.otherAsteroids[k]
	// 		if !ok {
	// 			continue
	// 		}
	// 		println(k, asteroids, destroyed, vaporize)
	// 		if len(asteroids) == 0 {
	// 			delete(m.MonitoringStation.otherAsteroids, k)
	// 			continue
	// 		}
	// 		x, m.MonitoringStation.otherAsteroids[k] = asteroids[len(asteroids)-1], asteroids[:len(asteroids)-1]
	// 		destroyed++
	// 		if destroyed >= vaporize {
	// 			return x
	// 		}
	// 	}
	// }

	// return x
}

// Read takes in a filename and returns the asteroid map
// specified in the file.
func Read(filename string) (*Map, error) {
	m := make([]Asteroid, 0)
	var asteroidMap Map
	file, err := os.Open(filename)
	if err != nil {
		return &asteroidMap, err
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
				m = append(m, Asteroid{X: x, Y: y, angles: make(map[float64]bool), otherAsteroids: make(map[float64][]Asteroid)})
			}
			x++
		}
		y++
	}
	asteroidMap = Map{Asteroids: m}
	return &asteroidMap, nil
}
