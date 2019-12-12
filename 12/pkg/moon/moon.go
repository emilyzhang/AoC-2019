package moon

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Moon struct {
	Velocity Velocity
	Position Position
	xCycle   int
	yCycle   int
	zCycle   int
}

type ThreeVector struct {
	X int
	Y int
	Z int
}

type Velocity ThreeVector
type Position ThreeVector

func New() *Moon {
	m := Moon{Velocity: Velocity{0, 0, 0}}
	return &m
}

func (m *Moon) StepGravity(other *Moon) {
	if other.Position.X > m.Position.X {
		other.Velocity.X--
		m.Velocity.X++
	} else if other.Position.X < m.Position.X {
		other.Velocity.X++
		m.Velocity.X--
	}

	if other.Position.Y > m.Position.Y {
		other.Velocity.Y--
		m.Velocity.Y++
	} else if other.Position.Y < m.Position.Y {
		other.Velocity.Y++
		m.Velocity.Y--
	}

	if other.Position.Z > m.Position.Z {
		other.Velocity.Z--
		m.Velocity.Z++
	} else if other.Position.Z < m.Position.Z {
		other.Velocity.Z++
		m.Velocity.Z--
	}
}

func (m *Moon) StepVelocity() {
	m.Position.X += m.Velocity.X
	m.Position.Y += m.Velocity.Y
	m.Position.Z += m.Velocity.Z
}

func (m *Moon) TotalEnergy() float64 {
	return m.PotentialEnergy() * m.KineticEnergy()
}

func (m *Moon) PotentialEnergy() float64 {
	return math.Abs(float64(m.Position.X)) + math.Abs(float64(m.Position.Y)) + math.Abs(float64(m.Position.Z))
}

func (m *Moon) KineticEnergy() float64 {
	return math.Abs(float64(m.Velocity.X)) + math.Abs(float64(m.Velocity.Y)) + math.Abs(float64(m.Velocity.Z))
}

func (m *Moon) FoundCycle() bool {
	return m.xCycle != 0 && m.yCycle != 0 && m.zCycle != 0
}

func Step(moons []*Moon) {
	for i, moon := range moons {
		for k := i; k < len(moons); k++ {
			other := moons[k]
			moon.StepGravity(other)
		}
	}

	for _, moon := range moons {
		moon.StepVelocity()
	}
}

func StepUntilMatch(moons []*Moon) int {
	prevX := make(map[string]bool)
	prevY := make(map[string]bool)
	prevZ := make(map[string]bool)
	steps := 0
	var cycleX, cycleY, cycleZ int
	for {
		hashX := Hash(moons, "x")
		hashY := Hash(moons, "y")
		hashZ := Hash(moons, "z")

		if prevX[hashX] && cycleX == 0 {
			cycleX = steps
		}
		if prevY[hashY] && cycleY == 0 {
			cycleY = steps
		}
		if prevZ[hashZ] && cycleZ == 0 {
			cycleZ = steps
		}

		prevX[hashX] = true
		prevY[hashY] = true
		prevZ[hashZ] = true

		if cycleX != 0 && cycleY != 0 && cycleZ != 0 {
			break
		}

		Step(moons)
		steps++
	}
	return LCM(cycleX, cycleY, cycleZ)
}

func (m *Moon) Hash(axis string) string {
	var pos, vel int
	switch axis {
	case "x":
		pos = m.Position.X
		vel = m.Velocity.X
	case "y":
		pos = m.Position.Y
		vel = m.Velocity.Y
	case "z":
		pos = m.Position.Z
		vel = m.Velocity.Z
	}
	return strconv.Itoa(pos) + strconv.Itoa(vel)
}

func Hash(moons []*Moon, axis string) string {
	var totalHash string
	for _, m := range moons {
		totalHash += m.Hash(axis)
	}
	return totalHash
}

func TotalEnergy(moons []*Moon) float64 {
	var energy float64
	for _, moon := range moons {
		energy += moon.TotalEnergy()
	}
	return energy
}

// Read takes in a filename and creates moons.
func Read(filename string) ([]*Moon, error) {
	moons := make([]*Moon, 0)
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return moons, err
	}

	data := strings.Split(strings.TrimSpace(string(f)), ">")
	for _, n := range data {
		if len(n) > 1 {
			m := strings.Split(strings.Trim(strings.TrimSpace(n), "<"), ",")
			x, err := strconv.Atoi(strings.Split(m[0], "=")[1])
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(strings.Split(m[1], "=")[1])
			if err != nil {
				return nil, err
			}
			z, err := strconv.Atoi(strings.Split(m[2], "=")[1])
			if err != nil {
				return nil, err
			}
			moon := Moon{Position: Position{X: x, Y: y, Z: z}}
			fmt.Println(moon)
			moons = append(moons, &moon)
		}
	}

	return moons, nil
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
