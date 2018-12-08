# Both these were solved by looking at the data with plt.imshow()

import numpy as np
import matplotlib.pyplot as plt

FIELD_SIZE = 400

field = [[None for x in range(FIELD_SIZE)] for x in range(FIELD_SIZE)]
coords = []

with open('day6in.txt', 'r') as f:
    for i, line in enumerate(f):
        x,y = line.split(',')
        x,y = int(x), int(y)
        coords.append((x,y))

def findclosest(x,y,coords):
    cid = 0
    for i, c in enumerate(coords):
        dist = abs(c[0]-x) + abs(c[1]-y)
        try: 
            if dist == closest:
                cid = 100
            if dist < closest:
                closest = dist
                cid = i
        except UnboundLocalError:
            closest = dist
    return cid

def findsafest(x,y,coords):
    slist = []
    for i, c in enumerate(coords):
        dist = abs(c[0]-x) + abs(c[1]-y)
        slist.append(dist)

    return sum(slist)

slists = []
count = 0
for x in range(FIELD_SIZE):
    for y in range(FIELD_SIZE):
        field[x][y] = findclosest(x,y,coords)
        slists.append(findsafest(x,y,coords))
        if slists[-1] < 10000:
            count+=1

print(count)
far = np.array(field, dtype=np.float32)

unique, counts = np.unique(far, return_counts=True)
print(dict(zip(unique,counts)))

plt.imshow(far)
plt.colorbar()
plt.show()
