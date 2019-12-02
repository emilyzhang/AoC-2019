# part 1
def main1():
    fuel = 0
    f = open('input.txt')
    for line in f:
        fuel+=int(line)//3-2
    return fuel

print(main1())



# part 2
def main2():
    fuel = 0
    f = open('input.txt')
    for line in f:
        s = int(line)
        while s > 0:
            s = s//3-2
            if s > 0:
                fuel+= s
    return fuel


print(main2())
