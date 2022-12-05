import re

crates = [
  'WBDNCFJ',
  'PZVQLST',
  'PZBGJT',
  'DTLJZBHC',
  'GVBJS',
  'PSQ',
  'BVDFLMPN',
  'PSMFBDLR',
  'VDTR'
]
crates = [list(c) for c in crates]

inp = list(map(lambda x: x.strip(), open('day05.in').readlines()))

for line in inp:
  i, f, t  = map(int, re.findall('move (\d+) from (\d+) to (\d+)', line)[0])
  crates[t-1].extend(crates[f-1][-i:][::-1])
  crates[f-1] = crates[f-1][:-i]
  
p1 = ''.join([c[-1] for c in crates])

crates = [
  'WBDNCFJ',
  'PZVQLST',
  'PZBGJT',
  'DTLJZBHC',
  'GVBJS',
  'PSQ',
  'BVDFLMPN',
  'PSMFBDLR',
  'VDTR'
]
crates = [list(c) for c in crates]

for line in inp:
  i, f, t  = map(int, re.findall('move (\d+) from (\d+) to (\d+)', line)[0])
  crates[t-1].extend(crates[f-1][-i:])
  crates[f-1] = crates[f-1][:-i]
  

p2 = ''.join([c[-1] for c in crates])

print(f'part1: {p1}\npart2: {p2}')

