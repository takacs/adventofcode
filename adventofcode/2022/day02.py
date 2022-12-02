inp = open('day02.in').readlines()

points = {
'A X':1+3,
'A Y':2+6,
'A Z':3+0,
'B X':1+0,
'B Y':2+3,
'B Z':3+6,
'C X':1+6,
'C Y':2+0,
'C Z':3+3,
}

points_strat = {
'A X':3+0,
'A Y':1+3,
'A Z':2+6,
'B X':1+0,
'B Y':2+3,
'B Z':3+6,
'C X':2+0,
'C Y':3+3,
'C Z':1+6,
}

p1 = sum(map(lambda x: points[x.strip()], inp))
p2 = sum(map(lambda x: points_strat[x.strip()], inp))
print(f'part1: {p1}\npart2: {p2}')
