def main(lower, upper):
    # for num in range(lower, upper):
        # print(num, ascending_digits(str(num)), same_adjacent(str(num)), larger_group(str(num)))
    return sum(ascending_digits(str(num)) and same_adjacent(str(num)) and larger_group(str(num)) for num in range(lower, upper))

def ascending_digits(num):
    return all(int(num[i]) <= int(num[i+1]) for i in range(len(num)-1))

def same_adjacent(num):
    return any(int(num[i]) == int(num[i+1]) for i in range(len(num)-1))

def larger_group(num):
    num = "0"+num+"0"
    # print(num)
    return any(int(num[i-1]) != int(num[i]) and int(num[i]) == int(num[i+1]) and int(num[i+1]) != int(num[i+2]) for i in range(1, len(num)-2))


print(main(240920, 789857))
