from math import sqrt

# https://www.redblobgames.com/grids/hexagons/

ds = open('day11in.txt', 'r').read().strip().split(',')

def distance(x, y, z):
    return max(abs(x), abs(y), abs(z))

furthest = []
x = y = z = 0
for d in ds:
    if d == 'n':
        y += 1
        z -= 1
    if d == 's':
        y -= 1
        z += 1
    if d == 'ne':
        x += 1
        z -= 1
    if d == 'sw':
        x -= 1
        z += 1
    if d == 'nw':
        x -= 1
        y += 1
    if d == 'se':
        x += 1
        y -= 1

    furthest.append(distance(x,y,z))

print(max(furthest), furthest[-1])
