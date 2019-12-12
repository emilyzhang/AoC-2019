package robot

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/emilyzhang/advent2019/11/pkg/intcode"
)

const (
	NORTH Direction = 0
	EAST  Direction = 1
	SOUTH Direction = 2
	WEST  Direction = 3
	BLACK Color     = 0
	WHITE Color     = 1
)

type Robot struct {
	panels          map[Location]Color
	painted         map[Location]bool
	program         *intcode.Program
	outputstate     int
	Errors          chan error
	direction       Direction
	currentLocation Location
	hull            *image.RGBA
}

type Location struct {
	X int
	Y int
}

func (l Location) Move(d Direction) Location {
	newLocation := Location{X: l.X, Y: l.Y}
	switch d {
	case NORTH:
		newLocation.Y++
	case EAST:
		newLocation.X++
	case SOUTH:
		newLocation.Y--
	case WEST:
		newLocation.X--
	}
	return newLocation
}

type Direction int
type Color int

func New(program *intcode.Program) *Robot {
	r := Robot{
		panels:      make(map[Location]Color),
		painted:     make(map[Location]bool),
		program:     program,
		outputstate: 0,
		direction:   NORTH,
	}
	// Start on white panel.
	r.panels[Location{0, 0}] = 1
	return &r
}

func (r *Robot) PanelsPainted() int {
	return len(r.painted)
}

func (r *Robot) Navigate(turn int) {
	switch turn {
	case 0:
		r.direction = Direction((r.direction - 1) % 4)
	case 1:
		r.direction = Direction((r.direction + 1) % 4)
	}
	if r.direction < 0 {
		r.direction = r.direction + 4
	}
	r.currentLocation = r.currentLocation.Move(r.direction)
}

func (r *Robot) Run() {
	color, direction := 0, 0
	for !r.program.Halted() {
		err := r.program.Run()
		if err != nil {
			fmt.Println(err)
			r.Errors <- err
			break
		}
		if r.program.RequiresInput() {
			r.program.Input(int(r.panels[r.currentLocation]))
		}
		if r.program.HasOutput() {
			switch r.outputstate {
			case 0:
				color = r.program.Output()
				r.outputstate = 1
			case 1:
				direction = r.program.Output()
				r.panels[r.currentLocation] = Color(color)
				r.painted[r.currentLocation] = true
				r.Navigate(direction)
				r.outputstate = 0
			}
		}
	}
}

func (r *Robot) Paint() {
	var maxwidth, maxheight float64
	for location := range r.painted {
		if maxwidth < math.Abs(float64(location.X)) {
			maxwidth = math.Abs(float64(location.X))
		}
		if maxheight < math.Abs(float64(location.Y)) {
			maxheight = math.Abs(float64(location.Y))
		}
	}
	w, h := int(maxwidth), int(maxheight)

	hull := image.NewRGBA(image.Rect(-w, -h, w, 2*h))
	var pixel color.Color
	for location, c := range r.panels {
		switch c {
		case BLACK:
			pixel = color.Black
		case WHITE:
			pixel = color.White
		}
		hull.Set(location.X, -location.Y, pixel)
	}

	r.hull = hull
}

func (r *Robot) PrintHull(filename string) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = png.Encode(outputFile, r.hull)
	if err != nil {
		return err
	}
	outputFile.Close()
	return nil
}
