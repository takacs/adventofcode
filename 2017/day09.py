inp = (open('day09in.txt','r').read().strip())

idx, lvl, score = 0, 0, 0
gcnt = 0
garbage = False

while idx != len(inp):
    cur = inp[idx]

    if garbage and (cur != '>'):
        gcnt += 1
        
    if cur == '<':
        garbage = True

    if cur == '>':
        garbage = False

    if not garbage:
        if cur == '{':
            lvl += 1
            score += lvl

        elif cur == '}':
            lvl -=1

    if cur == '!':
        gcnt -= 1
        idx += 2
    else:
        idx += 1

print(f'Part 1: {score}')
print(f'Part 2: {gcnt}')
