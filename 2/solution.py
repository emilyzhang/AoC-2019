import copy 
# part 1
def main1():
    f = open('input.txt')
    program = [int(s) for s in f.readline().strip().split(',')]
    program[1] = 12
    program[2] = 2
    i = 0
    while program[i] != 99:
        opcode = program[i]
        if opcode == 1:
            program[program[i+3]] = program[program[i+1]]+program[program[i+2]]
        elif opcode == 2:
            program[program[i+3]] = program[program[i+1]]*program[program[i+2]]
        i+=4
    print(program)


# print(main1())

def pretty(program):
    c = ""
    for i, p in enumerate(program):
        c += str(p) + " "
        if (i+1) % 4 == 0:
            c+="\n"
    return c


# part 2
def main2():
    f = open('input.txt')
    program = [int(s) for s in f.readline().strip().split(',')]
    print(pretty(program))
    for i in range(100):
        for j in range(100):
            p = copy.copy(program)
            p[1] = i
            p[2] = j
            target = run(p)
            if target == 19690720:
                print("yay!")
                return [i, j]
    return []

    # print(program)

def run(program):
    i = 0
    while program[i] != 99:
        opcode = program[i]
        # print(i, opcode, program[i+1], program[i+2], program[i+3])
        if opcode == 1:
            program[program[i+3]] = program[program[i+1]]+program[program[i+2]]
        elif opcode == 2:
            program[program[i+3]] = program[program[i+1]]*program[program[i+2]]
        i += 4
    return program[0]


sol = main2()
s = sol[0]*100+sol[1]
print(sol, s)
