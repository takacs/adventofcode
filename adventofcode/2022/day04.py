inp = list(map(lambda x: x.strip(), open('day04.in').readlines()))

def full_overlap(a, b):
  return 0 if set(a).difference(set(b)) else 1

def some_overlap(a, b):
  return 1 if set(a).intersection(set(b)) else 0

p1, p2 = 0, 0
for line in inp:
  a,b = map(lambda x: x.split('-'), line.split(','))
  a_r, b_r = range(int(a[0]), int(a[1])+1), range(int(b[0]), int(b[1])+1)
  p1 += max(full_overlap(a_r, b_r), full_overlap(b_r, a_r))
  p2 += max(some_overlap(a_r, b_r), some_overlap(b_r, a_r))

print(f'part1: {p1}\npart2: {p2}')
