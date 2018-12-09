from string import ascii_uppercase
import numpy as np

graph = { ac:[] for ac in ascii_uppercase}

with open('day7in.txt', 'r') as f:
    for line in f:
        l = line.split(' ')
        b, a = l[1], l[-3]
        graph[a].append(b)

comp = []
while(len(comp)!=26):
    for key, value in graph.items():
        if set(value).issubset(set(comp)) and key not in comp:
            comp.append(key)
            break

print(''.join(comp))

comp = []
total = 0
workers = [0,0,0,0,0]
curr_work = [0,0,0,0,0]

while(len(comp)!=26):
    for key, value in graph.items():
        if set(value).issubset(set(comp)) and key not in comp and key not in curr_work:
            try:
                ind = workers.index(0)
                workers[ind] = 60+ascii_uppercase.find(key)+1
                curr_work[ind] = key
                continue
            except ValueError:
                break
    warr = np.array(workers)
    total += np.min(warr[np.nonzero(warr)])
    workers =  [w - np.min(warr[np.nonzero(warr)]) for w in workers]
    for i, c in enumerate(workers):
        if c == 0 and curr_work[i] != 0:
            comp.append(curr_work[i])
            curr_work[i] = 0
    workers = [ w if w >= 0 else 0 for w in workers ]
print(total)
