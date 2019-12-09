package main

import (
	"./pkg/intcode"
	"fmt"
)

func main() {
	program, err := intcode.New("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(program.Code)
	max, err := largestAmplifiedOutput(program)
	if err != nil {
		fmt.Println(err)
	}
	println(max)
	// _, err = intcode.Run(program, []int{1, 4})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// largest, err := largestOutput(program)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// print(largest)
}

func feedbackloop(program intcode.Program, phaseSequence []int) (int, error) {
	var err error
	p1 := spawn(program)
	p1.Input = append(p1.Input, 0)
	p1.Input = append(p1.Input, phaseSequence[0])
	p1, err = intcode.Run(p1)
	if err != nil {
		return 0, err
	}

	p2 := spawn(program)
	p2.Input = append(p2.Input, phaseSequence[1])
	p2.Input = append(p2.Input, p1.Output[0])
	p2, err = intcode.Run(p2)
	if err != nil {
		return 0, err
	}

	p3 := spawn(program)
	p3.Input = append(p3.Input, phaseSequence[2])
	p3.Input = append(p3.Input, p2.Output[0])
	p3, err = intcode.Run(p3)
	if err != nil {
		return 0, err
	}

	p4 := spawn(program)
	p4.Input = append(p4.Input, phaseSequence[3])
	p4.Input = append(p4.Input, p3.Output[0])
	p4, err = intcode.Run(p4)
	if err != nil {
		return 0, err
	}

	p5 := spawn(program)
	p5.Input = append(p5.Input, phaseSequence[4])
	p5.Input = append(p5.Input, p4.Output[0])
	p5, err = intcode.Run(p5)
	if err != nil {
		return 0, err
	}

	o1, o2, o3, o4, o5 := 0, 1, 1, 1, 1
	for true {
		if !p1.Complete {
			p1, o1, err = runAmplifier(p1, p5, o1)
			if err != nil {
				return 0, err
			}
		}
		if !p2.Complete {
			p2, o2, err = runAmplifier(p2, p1, o2)
			if err != nil {
				return 0, err
			}
		}
		if !p3.Complete {
			p3, o3, err = runAmplifier(p3, p2, o3)
			if err != nil {
				return 0, err
			}
		}
		if !p4.Complete {
			p4, o4, err = runAmplifier(p4, p3, o4)
			if err != nil {
				return 0, err
			}
		}
		if !p5.Complete {
			p5, o5, err = runAmplifier(p5, p4, o5)
			if err != nil {
				return 0, err
			}
		}
		if p1.Complete && p2.Complete && p3.Complete && p4.Complete && p5.Complete {
			break
		}
	}
	return p5.Output[o5], nil
}

func runAmplifier(program *intcode.Program, prev *intcode.Program, outputIndex int) (*intcode.Program, int, error) {
	program.Input = append(program.Input, prev.Output[outputIndex])
	program.Input = append(program.Input, prev.Output[outputIndex])
	p, err := intcode.Run(program)
	return p, outputIndex + 2, err
}

func spawn(program intcode.Program) *intcode.Program {
	p := intcode.Program{}

	c := make([]int, len(program.Code))
	copy(c, program.Code)
	i := make([]int, len(program.Input))
	copy(i, program.Input)
	o := make([]int, len(program.Output))
	copy(i, program.Output)

	p.Code = c
	p.Input = i
	p.Output = o
	p.CurrentIndex = program.CurrentIndex
	p.InputIndex = program.InputIndex
	p.Complete = program.Complete
	return &p
}

func largestAmplifiedOutput(program intcode.Program) (int, error) {
	maxOutput := 0
	for i := 5; i <= 9; i++ {
		for j := 5; j <= 9; j++ {
			if j == i {
				continue
			}
			for k := 5; k <= 9; k++ {
				if k == j || k == i {
					continue
				}
				for l := 5; l <= 9; l++ {
					if l == k || l == i || l == j {
						continue
					}
					for m := 5; m <= 9; m++ {
						if m == k || m == i || m == j || m == l {
							continue
						}
						o, err := feedbackloop(program, []int{i, j, k, l, m})
						if err != nil {
							return 0, err
						}
						if o > maxOutput {
							maxOutput = o
						}
					}
				}
			}
		}
	}
	return maxOutput, nil
}

// func largestOutput(program []int) (int, error) {
// 	maxOutput := 0
// 	for i := 0; i <= 4; i++ {
// 		p := make([]int, len(program))
// 		copy(p, program)
// 		input := []int{i, 0}
// 		output, err := intcode.Run(program, input)
// 		if err != nil {
// 			return maxOutput, err
// 		}
// 		o := output[0]
// 		for j := 0; j <= 4; j++ {
// 			if j == i {
// 				continue
// 			}
// 			p := make([]int, len(program))
// 			copy(p, program)
// 			input := []int{j, o}
// 			output, err := intcode.Run(program, input)
// 			if err != nil {
// 				return maxOutput, err
// 			}
// 			o := output[0]
// 			for k := 0; k <= 4; k++ {
// 				if k == j || k == i{
// 					continue
// 				}
// 				p := make([]int, len(program))
// 				copy(p, program)
// 				input := []int{k, o}
// 				output, err := intcode.Run(program, input)
// 				if err != nil {
// 					return maxOutput, err
// 				}
// 				o := output[0]
// 				for l := 0; l <= 4; l++ {
// 					if l == k || l == i || l == j {
// 						continue
// 					}
// 					p := make([]int, len(program))
// 					copy(p, program)
// 					input := []int{l, o}
// 					output, err := intcode.Run(program, input)
// 					if err != nil {
// 						return maxOutput, err
// 					}
// 					o := output[0]
// 					for m := 0; m <= 4; m++ {
// 						if m == k || m == i || m == j || m == l{
// 							continue
// 						}
// 						println(i, j, k, l, m)
// 						p := make([]int, len(program))
// 						copy(p, program)
// 						input := []int{m, o}
// 						output, err := intcode.Run(program, input)
// 						if err != nil {
// 							return maxOutput, err
// 						}
// 						o := output[0]
// 						// println("at the end", o)
// 						if o > maxOutput {
// 							maxOutput = o
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return maxOutput, nil
// }
