# part 1
def read_input(file):
    f = open(file)
    line1 = f.readline().strip().split(',')
    line2 = f.readline().strip().split(',')
    return line1, line2

def find_intervals(path):
    intervals = []
    x_intervals = {}
    y_intervals = {}
    x, y, = 0, 0
    for d in path:
        direction = d[0]
        val = int(d[1:])
        prevx = x
        prevy = y
        if direction == "R":
            x+=val
            y_intervals[(min(prevx, x), max(prevx, x))] = y
        elif direction == "D": 
            y-=val
            x_intervals[x] = (min(prevy, y), max(prevy, y))
        elif direction == "L": 
            x-=val
            y_intervals[(min(prevx, x), max(prevx, x))] = y
        elif direction == "U": 
            y+=val
            x_intervals[x] = (min(prevy, y), max(prevy, y))
        intervals.append([(prevx, prevy), (x, y)])
    # print(x_intervals, y_intervals)
    return x_intervals, y_intervals



def find_intersections(x_intervals, y_intervals):
    intersections = []
    for x in sorted(x_intervals):
        for x_pair in sorted(y_intervals, key=lambda x: x[0]):
            if x_pair[0] <= x and x_pair[1] >= x:
                # print(x_pair, x)
                # print(x_intervals[x], y_intervals[x_pair])
                # print("====")
                if y_intervals[x_pair] >= x_intervals[x][0] and y_intervals[x_pair] <= x_intervals[x][1]:
                    intersections.append((x, y_intervals[x_pair]))
            if x_pair[0] > x:
                continue
    return intersections



def distance(point):
    return abs(point[0]) + abs(point[1])

line1, line2 = read_input("input.txt")
x1, y1 = find_intervals(line1)
x2, y2 = find_intervals(line2)
i = find_intersections(x1, y2) + find_intersections(x2, y1)
distances = [distance(point) for point in i]
print(min(distances))


# part 2
def find_intervals_with_steps(path):
    x_intervals = {}
    y_intervals = {}
    x, y, = 0, 0
    steps = 0
    # seen = {}
    for d in path:
        direction = d[0]
        val = int(d[1:])

        # if seen.get((x, y)):
        #     seen[(x, y)] = min(steps, seen[(x, y)])
        # else:
        #     seen[(x, y)] = steps
        
        prevx = x
        prevy = y
        if direction == "R":
            # for i in range(val):
            #     if seen.get((x+i, y)):
            #         seen[(x+i, y)] = min(steps +i, seen[(x+i, y)])
            #     else:
            #         seen[(x+i, y)] = steps + i
            x += val
            y_intervals[(prevx, x, steps)] = y
        elif direction == "D":
            # for i in range(val):
            #     if seen.get((x, y+i)):
            #         seen[(x, y+i)] = min(steps + i, seen[(x, y+i)])
            #     else:
            #         seen[(x, y+i)] = steps + i
            y -= val
            x_intervals[x] = (prevy, y, steps)
        elif direction == "L":
            # for i in range(val):
            #     if seen.get((x+i, y)):
            #         seen[(x+i, y)] = min(steps + i, seen[(x+i, y)])
            #     else:
            #         seen[(x+i, y)] = steps + i
            x -= val
            y_intervals[(prevx, x, steps)] = y
        elif direction == "U":
            # for i in range(val):
            #     if seen.get((x, y+i)):
            #         seen[(x, y+i)] = min(steps + i, seen[(x, y+i)])
            #     else:
            #         seen[(x, y+i)] = steps + i
            y += val
            x_intervals[x] = (prevy, y, steps)
        
        
        # if seen.get((x, y)):
        #     steps = seen[(x,y)]
        steps += val
    return x_intervals, y_intervals


def find_intersections_with_steps(x_intervals, y_intervals):
    intersections_steps = []
    for x in sorted(x_intervals):
        for x_pair in sorted(y_intervals, key=lambda x: x[0]):
            smallx = min(x_pair[0], x_pair[1])
            largex = max(x_pair[0], x_pair[1])
            if smallx < x and largex > x:
                # print(x_pair, x)
                # print(x_intervals[x], y_intervals[x_pair])
                # print("====")
                smally = min(x_intervals[x][0], x_intervals[x][1])
                largey = max(x_intervals[x][0], x_intervals[x][1])
                if y_intervals[x_pair] > smally and y_intervals[x_pair] < largey:
                    steps = x_intervals[x][2]+x_pair[2]
                    extra_steps_to_intersection = abs(x-x_pair[0]) + \
                        abs(y_intervals[x_pair]-x_intervals[x][0])
                    intersections_steps.append(
                        steps+extra_steps_to_intersection)
    return intersections_steps


line1, line2 = read_input("input.txt")
# line1, line2 = read_input("ex1.txt")
x1, y1 = find_intervals_with_steps(line1)
x2, y2 = find_intervals_with_steps(line2)
i = find_intersections_with_steps(x1, y2) + find_intersections_with_steps(x2, y1)
print(min(i))
