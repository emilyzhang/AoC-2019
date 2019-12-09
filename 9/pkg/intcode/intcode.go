package intcode

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Program .
type Program struct {
	state   []int
	pointer int
	status  Status
	input   int
	output  int
}

type Source []int
type Status int

const (
	DEFAULT Status = iota
	REQUIRES_INPUT
	HAS_INPUT
	HAS_OUTPUT
	HALTED
)

// Halted returns true if the program has halted,
// otherwise false.
func (p *Program) Halted() bool {
	return p.status == HALTED
}

func (p *Program) RequiresInput() bool {
	return p.status == REQUIRES_INPUT
}

func (p *Program) HasOutput() bool {
	return p.status == HAS_OUTPUT
}

func (p *Program) Status() Status {
	return p.status
}

// func (p *Program) State() []int {
// 	return p.state
// }

// Read takes in a filename and returns the source code
// specified in the file.
func Read(filename string) (Source, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	data := strings.Split(strings.TrimSpace(string(f)), ",")

	var source []int
	for _, n := range data {
		num, err := strconv.Atoi(string(n))
		if err != nil {
			return nil, err
		}
		source = append(source, num)
	}

	return Source(source), nil
}

func New(source Source) *Program {
	s := make([]int, len(source))
	copy(s, source)
	return &Program{state: s, pointer: 0, status: DEFAULT}
}

func (p *Program) Run() error {
	for p.status == DEFAULT || p.status == HAS_INPUT {
		i, err := p.DecodeInstruction()
		if err != nil {
			return err
		}
		err = p.Step(i)
		if err != nil {
			return err
		}
		fmt.Println(p.state, p.pointer, p.status)
	}
	return nil
}

func (p *Program) Input(input int) {
	p.input = input
	p.status = HAS_INPUT
}

func (p *Program) Output() int {
	p.status = DEFAULT
	return p.output
}

func (p *Program) Step(instruction Instruction) error {
	switch instruction.(type) {
	case Add:
		i := instruction.(Add)
		p.state[i.Address] = i.Parameter1 + i.Parameter2
		p.pointer += i.Next
	case Multiply:
		i := instruction.(Multiply)
		p.state[i.Address] = i.Parameter1 * i.Parameter2
		p.pointer += i.Next
	case Input:
		i := instruction.(Input)
		if p.status != HAS_INPUT {
			p.status = REQUIRES_INPUT
		} else {
			p.state[i.Address] = p.input
			p.status = DEFAULT
			p.pointer += i.Next
		}
	case Output:
		i := instruction.(Output)
		p.output = p.state[i.Address]
		p.status = HAS_OUTPUT
		p.pointer += i.Next
	case Jump:
		i := instruction.(Jump)
		if i.Jump {
			p.pointer = i.JumpTo
		} else {
			p.pointer += i.Next
		}
	case Equals:
		i := instruction.(Equals)
		if i.Parameter1 == i.Parameter2 {
			p.state[i.Address] = 1
		} else {
			p.state[i.Address] = 0
		}
		p.pointer += i.Next
	case Compare:
		i := instruction.(Compare)
		if i.Parameter1 < i.Parameter2 {
			p.state[i.Address] = 1
		} else {
			p.state[i.Address] = 0
		}
		p.pointer += i.Next
	case Halt:
		p.status = HALTED
	default:
		return errors.New("unexpected instruction")
	}
	return nil
}

func (p *Program) DecodeInstruction() (Instruction, error) {
	instr := p.state[p.pointer]
	opcode := Opcode(instr % 100)
	p1mode := (instr / 100) % 10
	p2mode := (instr / 1000) % 10

	if opcode == 99 {
		return Halt{}, nil
	}

	// println(instr, opcode, p1mode, p2mode)

	p1, p2 := p.state[p.pointer+1], p.state[p.pointer+2]
	if opcode != 3 && opcode != 4 {
		if p1mode == 0 {
			p1 = p.state[p1]
		}
		if p2mode == 0 {
			p2 = p.state[p2]
		}
	}

	var instruction Instruction
	switch opcode {
	case 1:
		instruction = Add{Parameter1: p1, Parameter2: p2, Address: p.state[p.pointer+3], Next: 4}
	case 2:
		instruction = Multiply{Parameter1: p1, Parameter2: p2, Address: p.state[p.pointer+3], Next: 4}
	case 3:
		instruction = Input{Address: p1, Next: 2}
	case 4:
		instruction = Output{Address: p1, Next: 2}
	case 5:
		jump := false
		if p1 != 0 {
			jump = true
		}
		instruction = Jump{Jump: jump, Next: 3, JumpTo: p2}
	case 6:
		jump := false
		if p1 == 0 {
			jump = true
		}
		instruction = Jump{Jump: jump, Next: 3, JumpTo: p2}
	case 7:
		instruction = Compare{Parameter1: p1, Parameter2: p2, Address: p.state[p.pointer+3], Next: 4}
	case 8:
		instruction = Equals{Parameter1: p1, Parameter2: p2, Address: p.state[p.pointer+3], Next: 4}
	// case 9: instruction = Input{Address: p1, Next: 2}
	default:
		return nil, errors.New("unrecognized instruction")
	}

	return instruction, nil
}

// Run .
// func Run(program *Program) (*Program, error) {
// 	code := program.Code
// 	i := program.CurrentIndex
// 	inputindex := program.InputIndex
// 	var instruction, opcode, p1, p2, index, parameter1, parameter2 int
// 	for code[i] != 99 {
// 		instruction = code[i]
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

// 		parameter1, parameter2 = code[i+1], code[i+2]
// 		if opcode != 3 && opcode != 4 {
// 			if p1 == 0 {
// 				parameter1 = code[code[i+1]]
// 			}
// 			if p2 == 0 {
// 				parameter2 = code[code[i+2]]
// 			}
// 		}

// 		switch opcode {
// 		case 1:
// 			code[code[i+3]] = parameter1 + parameter2
// 			i += 4
// 		case 2:
// 			code[code[i+3]] = parameter1 * parameter2
// 			i += 4
// 		case 3:
// 			if inputindex > len(program.Input)-1 {
// 				program.Code = code
// 				program.CurrentIndex = i
// 				program.InputIndex = inputindex
// 				return program, nil
// 			}
// 			code[code[i+1]] = program.Input[inputindex]
// 			inputindex++
// 			i += 2
// 		case 4:
// 			// fmt.Println("output:", code[code[i+1]])
// 			program.Output = append(program.Output, code[code[i+1]])
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
// 				code[code[i+3]] = 1
// 			} else {
// 				code[code[i+3]] = 0
// 			}
// 			i += 4
// 		case 8:
// 			if parameter1 == parameter2 {
// 				code[code[i+3]] = 1
// 			} else {
// 				code[code[i+3]] = 0
// 			}
// 			i += 4
// 		default:
// 			o := strconv.Itoa(opcode)
// 			return program, errors.New("unexpected opcode " + o)
// 		}
// 	}
// 	program.Code = code
// 	program.Complete = true
// 	return program, nil
// }
