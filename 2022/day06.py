inp = list(map(lambda x: x.strip(), open('day06.in').readlines()))[0]

def makezip(nr, inp):
  return zip(*[inp[n:-nr+n] for n in range(nr)])

marker = 4
for i, a in enumerate(makezip(marker, inp)):
  if len(set(a)) == len(a):
    p1 = i+marker
    break

marker = 14
for i, a in enumerate(makezip(marker, inp)):
  if len(set(a)) == len(a):
    p2 = i+marker
    break


print(f'part1: {p1}\npart2: {p2}')

