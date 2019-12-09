package intcode

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Program .
type Program struct {
	Code         []int
	Output       []int
	Input        []int
	Complete     bool
	CurrentIndex int
	InputIndex   int
}

// New .
func New(input string) (Program, error) {
	file, err := os.Open(input)
	if err != nil {
		return Program{}, err
	}
	defer file.Close()

	var s []string
	var program []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = strings.Split(strings.TrimSpace(scanner.Text()), ",")
	}
	for _, str := range s {
		i, err := strconv.Atoi(str)
		if err != nil {
			return Program{}, err
		}
		program = append(program, i)
	}

	return Program{
		Code:         program,
		Output:       make([]int, 0),
		Input:        make([]int, 0),
		Complete:     false,
		CurrentIndex: 0,
		InputIndex:   0,
	}, nil
}

// Run .
func Run(program *Program) (*Program, error) {
	code := program.Code
	i := program.CurrentIndex
	inputindex := program.InputIndex
	output := make([]int, 0)
	var instruction, opcode, p1, p2, index, parameter1, parameter2 int
	for code[i] != 99 {
		instruction = code[i]
		opcode, p1, p2, index = 0, 0, 0, 0
		for instruction != 0 && index < 4 {
			switch index {
			case 0:
				opcode = instruction % 100
				instruction = instruction / 100
			case 1:
				p1 = instruction % 10
				instruction = instruction / 10
			case 2:
				p2 = instruction % 10
				instruction = instruction / 10
			}
			index++
		}

		parameter1, parameter2 = code[i+1], code[i+2]
		if opcode != 3 && opcode != 4 {
			if p1 == 0 {
				parameter1 = code[code[i+1]]
			}
			if p2 == 0 {
				parameter2 = code[code[i+2]]
			}
		}

		switch opcode {
		case 1:
			code[code[i+3]] = parameter1 + parameter2
			i += 4
		case 2:
			code[code[i+3]] = parameter1 * parameter2
			i += 4
		case 3:
			if inputindex > len(program.Input)-1 {
				program.Code = code
				program.CurrentIndex = i
				program.InputIndex = inputindex
				program.Output = output
				return program, nil
			}
			code[code[i+1]] = program.Input[inputindex]
			inputindex++
			i += 2
		case 4:
			fmt.Println("output:", code[code[i+1]])
			output = append(output, code[code[i+1]])
			i += 2
		case 5:
			if parameter1 != 0 {
				i = parameter2
			} else {
				i += 3
			}
		case 6:
			if parameter1 == 0 {
				i = parameter2
			} else {
				i += 3
			}
		case 7:
			if parameter1 < parameter2 {
				code[code[i+3]] = 1
			} else {
				code[code[i+3]] = 0
			}
			i += 4
		case 8:
			if parameter1 == parameter2 {
				code[code[i+3]] = 1
			} else {
				code[code[i+3]] = 0
			}
			i += 4
		default:
			o := strconv.Itoa(opcode)
			return program, errors.New("unexpected opcode " + o)
		}
	}
	program.Code = code
	program.Complete = true
	return program, nil
}

// Run .
// func Run(program []int, input []int) ([]int, error) {
// 	i := 0
// 	inputindex := 0
// 	output := make([]int, 0)
// 	var instruction, opcode, p1, p2, index, parameter1, parameter2 int
// 	for program[i] != 99 {
// 		instruction = program[i]
// 		opcode, p1, p2, index = 0, 0, 0, 0
// 		for instruction != 0 && index < 4 {
// 			switch index {
// 			case 0:
// 				opcode = instruction % 100
// 				instruction = instruction / 100
// 			case 1:
// 				p1 = instruction % 10
// 				instruction = instruction / 10
// 			case 2:
// 				p2 = instruction % 10
// 				instruction = instruction / 10
// 			}
// 			index++
// 		}

// 		parameter1, parameter2 = program[i+1], program[i+2]
// 		if opcode != 3 && opcode != 4 {
// 			if p1 == 0 {
// 				parameter1 = program[program[i+1]]
// 			}
// 			if p2 == 0 {
// 				parameter2 = program[program[i+2]]
// 			}
// 		}

// 		switch opcode {
// 		case 1:
// 			program[program[i+3]] = parameter1 + parameter2
// 			i += 4
// 		case 2:
// 			program[program[i+3]] = parameter1 * parameter2
// 			i += 4
// 		case 3:
// 			// fmt.Print("input? ")
// 			// var input string
// 			// fmt.Scanln(&input)
// 			// inputint, err := strconv.Atoi(input)
// 			// if err != nil {
// 			// 	return err
// 			// }
// 			program[program[i+1]] = input[inputindex]
// 			inputindex++
// 			i += 2
// 		case 4:
// 			fmt.Println("output:", program[program[i+1]])
// 			output = append(output, program[program[i+1]])
// 			i += 2
// 		case 5:
// 			if parameter1 != 0 {
// 				i = parameter2
// 			} else {
// 				i += 3
// 			}
// 		case 6:
// 			if parameter1 == 0 {
// 				i = parameter2
// 			} else {
// 				i += 3
// 			}
// 		case 7:
// 			if parameter1 < parameter2 {
// 				program[program[i+3]] = 1
// 			} else {
// 				program[program[i+3]] = 0
// 			}
// 			i += 4
// 		case 8:
// 			if parameter1 == parameter2 {
// 				program[program[i+3]] = 1
// 			} else {
// 				program[program[i+3]] = 0
// 			}
// 			i += 4
// 		default:
// 			o := strconv.Itoa(opcode)
// 			return output, errors.New("unexpected opcode " + o)
// 		}
// 	}
// 	return output, nil
// }
