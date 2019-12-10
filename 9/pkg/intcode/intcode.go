package intcode

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// Program .
type Program struct {
	state        []int
	pointer      int
	status       Status
	input        int
	output       int
	relativeBase int
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
			p.state[i.InputAddress] = p.input
			p.status = DEFAULT
			p.pointer += i.Next
		}
	case Output:
		i := instruction.(Output)
		p.output = i.Output
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
	case AdjustRelativeBase:
		i := instruction.(AdjustRelativeBase)
		p.relativeBase += i.Adjustment
		p.pointer += i.Next
	case Halt:
		p.status = HALTED
	default:
		return errors.New("unexpected instruction")
	}
	return nil
}

func (p *Program) provisionMemory(p1, p2, p3, p1mode, p2mode, p3mode int) {
	if len(p.state) < p.pointer+3 {
		space := make([]int, p.pointer+4)
		p.state = append(p.state, space...)
	}
	p1check, p2check, p3check := 0, 0, 0
	switch p1mode {
	case 0:
		p1check = p1
	case 2:
		p1check = p1 + p.relativeBase
	}
	switch p2mode {
	case 0:
		p2check = p2
	case 2:
		p2check = p2 + p.relativeBase
	}
	switch p3mode {
	case 0:
		p3check = p3
	case 2:
		p3check = p3 + p.relativeBase
	}
	requiredMem := max(p1check, p2check, p3check)
	if len(p.state) < requiredMem {
		space := make([]int, requiredMem-len(p.state)+5)
		p.state = append(p.state, space...)
	}
}

func (p *Program) CalculateParameters(opcode Opcode, p1mode, p2mode, p3mode int) (p1 int, p2 int, p3 int) {
	p1, p2, p3 = p.state[p.pointer+1], p.state[p.pointer+2], 0
	if opcode == 1 || opcode == 2 || opcode == 7 || opcode == 8 {
		p3 = p.state[p.pointer+3]
	}
	p.provisionMemory(p1, p2, p3, p1mode, p2mode, p3mode)

	switch opcode {
	case 1, 2, 7, 8:
		switch p1mode {
		case 0:
			p1 = p.state[p1]
		case 2:
			p1 = p.state[p1+p.relativeBase]
		}
		switch p2mode {
		case 0:
			p2 = p.state[p2]
		case 2:
			p2 = p.state[p2+p.relativeBase]
		}
		switch p3mode {
		case 2:
			p3 += p.relativeBase
		}
	case 3:
		if p1mode == 2 {
			p1 += p.relativeBase
		}
	case 4:
		switch p1mode {
		case 0:
			p1 = p.state[p1]
		case 2:
			p1 = p.state[p1+p.relativeBase]
		}
	case 5, 6:
		switch p1mode {
		case 0:
			p1 = p.state[p1]
		case 2:
			p1 = p.state[p1+p.relativeBase]
		}
		switch p2mode {
		case 0:
			p2 = p.state[p2]
		case 2:
			p2 = p.state[p2+p.relativeBase]
		}
	case 9:
		switch p1mode {
		case 0:
			p1 = p.state[p1]
		case 2:
			p1 = p.state[p1+p.relativeBase]
		}
	}
	return p1, p2, p3
}

func (p *Program) DecodeInstruction() (Instruction, error) {
	instr := p.state[p.pointer]
	opcode := Opcode(instr % 100)
	p1mode := (instr / 100) % 10
	p2mode := (instr / 1000) % 10
	p3mode := (instr / 10000) % 10

	if opcode == 99 {
		return Halt{}, nil
	}

	p1, p2, p3 := p.CalculateParameters(opcode, p1mode, p2mode, p3mode)

	var instruction Instruction
	switch opcode {
	case 1:
		instruction = Add{Parameter1: p1, Parameter2: p2, Address: p3, Next: 4}
	case 2:
		instruction = Multiply{Parameter1: p1, Parameter2: p2, Address: p3, Next: 4}
	case 3:
		instruction = Input{InputAddress: p1, Next: 2}
	case 4:
		instruction = Output{Output: p1, Next: 2}
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
		instruction = Compare{Parameter1: p1, Parameter2: p2, Address: p3, Next: 4}
	case 8:
		instruction = Equals{Parameter1: p1, Parameter2: p2, Address: p3, Next: 4}
	case 9:
		instruction = AdjustRelativeBase{Adjustment: p1, Next: 2}
	default:
		return nil, errors.New("unrecognized instruction")
	}

	return instruction, nil
}

func max(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
