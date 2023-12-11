inp = """
R5, L2, L1, R1, R3, R3, L3, R3, R4, L2, R4, L4, R4, R3, L2, L1, L1, R2, R4, R4, L4, R3, L2, R1, L4, R1, R3, L5, L4, L5, R3, L3, L1, L1, R4, R2, R2, L1, L4, R191, R5, L2, R46, R3, L1, R74, L2, R2, R187, R3, R4, R1, L4, L4, L2, R4, L5, R4, R3, L2, L1, R3, R3, R3, R1, R1, L4, R4, R1, R5, R2, R1, R3, L4, L2, L2, R1, L3, R1, R3, L5, L3, R5, R3, R4, L1, R3, R2, R1, R2, L4, L1, L1, R3, L3, R4, L2, L4, L5, L5, L4, R2, R5, L4, R4, L2, R3, L4, L3, L5, R5, L4, L2, R3, R5, R5, L1, L4, R3, L1, R2, L5, L1, R4, L1, R5, R1, L4, L4, L4, R4, R3, L5, R1, L3, R4, R3, L2, L1, R1, R2, R2, R2, L1, L1, L2, L5, L3, L1
""".strip().replace(' ', '').split(',')

way = [(0,1), (1,0), (0,-1), (-1,0)]
wayp = ['N', 'E', 'S', 'W']
pos = [0,0]
curr = 0

visited = [[0,0]]
first = True

for d in inp:
    if d[0] == 'R':
        curr+=1
    else:
        curr-=1
    pos[0] += way[curr%4][0]*int(d[1:])
    pos[1] += way[curr%4][1]*int(d[1:])

    for i in range(1,int(d[1:])+1):
        nextpos = [visited[-1][0] + way[curr%4][0], visited[-1][1] + way[curr%4][1]]
        if first and nextpos in visited:
            first = False
            print(f'Part 2: {abs(nextpos[0]) + abs(nextpos[1])}')
        visited.append(nextpos)


print(f'Part 1: {abs(pos[0]) + abs(pos[1])}')



