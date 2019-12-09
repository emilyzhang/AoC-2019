package intcode

type Opcode int

type Add struct {
	Parameter1 int
	Parameter2 int
	Address    int
	Next       int
}

func (a Add) isInstruction() {}

type Multiply struct {
	Parameter1 int
	Parameter2 int
	Address    int
	Next       int
}

func (m Multiply) isInstruction() {}

type Compare struct {
	Parameter1 int
	Parameter2 int
	Address    int
	Next       int
}

func (c Compare) isInstruction() {}

type Equals struct {
	Parameter1 int
	Parameter2 int
	Address    int
	Next       int
}

func (e Equals) isInstruction() {}

type Jump struct {
	Jump   bool
	JumpTo int
	Next   int
}

func (j Jump) isInstruction() {}

type Input struct {
	Address int
	Next    int
}

func (i Input) isInstruction() {}

type Output struct {
	Address int
	Next    int
}

func (o Output) isInstruction() {}

type Halt struct{}

func (h Halt) isInstruction() {}

type Instruction interface {
	isInstruction()
}
