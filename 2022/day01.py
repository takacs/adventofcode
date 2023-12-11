p1 = max([sum(map(int, l.split())) for l in open('day01.in').read().split('\n\n')])
p2 = sum(sorted([sum(map(int, l.split())) for l in open('day01.in').read().split('\n\n')],key=int)[-3:])
print(f'part1: {p1}\npart2: {p2}')
