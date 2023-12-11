inp = list(map(int, open('day05in.txt','r').readlines()))

pos = 0
ppos = 0
steps = 0

while pos < len(inp):
    ppos = pos
    pos += inp[pos]
    inp[ppos] += 1
    steps+=1

print(f'Part 1: {steps}')

inp = list(map(int, open('day05in.txt','r').readlines()))

pos = 0
ppos = 0
steps = 0

while pos < len(inp):
    ppos = pos
    off = inp[pos]
    pos += inp[pos]
    if off >= 3:
        inp[ppos] -= 1
    else:
        inp[ppos] += 1
    steps+=1

print(f'Part 2: {steps}')
