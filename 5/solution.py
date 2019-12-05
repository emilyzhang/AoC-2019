# part 1
def read_input(file):
    f = open(file)
    return [int(s) for s in f.readline().strip().split(',')]

def run(program):
    i = 0
    while program[i] != 99:
        instruction = program[i]
        opcode, p1, p2, p3, index = 0, 0, 0, 0, 0
        while instruction != 0 and index < 4:
            if index == 0:
                opcode = instruction % 100
                instruction = instruction // 100
            elif index == 1:
                p1 = instruction % 10
                instruction = instruction // 10
            elif index == 2:
                p2 = instruction % 10
                instruction = instruction // 10
            elif index == 3:
                p3 = instruction % 10
                instruction = instruction // 10
            index += 1

        
        parameter1, parameter2 = program[i+1], program[i+2]
        if opcode != 3 and opcode != 4:
            if p1 == 0:
                parameter1 = program[program[i+1]]
            if p2 == 0:
                parameter2 = program[program[i+2]]
        print(opcode, p1, p2, p3, parameter1, parameter2)

        
        if opcode == 1:
            program[program[i+3]] = parameter1+parameter2
            i += 4
        elif opcode == 2:
            program[program[i+3]] = parameter1*parameter2
            i += 4
        elif opcode == 3:
            program[program[i+1]] = input("input? ")
            i += 2
        elif opcode == 4:
            print(program[program[i+1]])
            i += 2
        elif opcode == 5:
            if parameter1 != 0:
                i = parameter2
            else:
                i+=3
        elif opcode == 6:
            if parameter1 == 0:
                i = parameter2
            else:
                i += 3
        elif opcode == 7:
            if parameter1 < parameter2:
                program[program[i+3]] = 1
            else:
                program[program[i+3]] = 0
            i += 4
        elif opcode == 8:
            if parameter1 == parameter2:
                program[program[i+3]] = 1
            else:
                program[program[i+3]] = 0
            i += 4


def pretty(program):
    c = ""
    for i, p in enumerate(program):
        c += str(p) + " "
        if (i+1) % 4 == 0:
            c += "\n"
    return c

if __name__ == "__main__":
    # program = read_input("test.txt")
    # run(program)

    program = read_input("input.txt")
    run(program)
