inp = list(map(lambda x: x.strip(), open('day03.in').readlines()))

L = []
for line in inp:
  half = len(line)//2
  fhalf, shalf = line[:half], line[half:]
  L.extend(set(fhalf).intersection(set(shalf)))

L2 = []
for i in range(0, len(inp)-2, 3):
  bags = inp[i:i+3]
  L2.extend(set(bags[0]).intersection(set(bags[1])).intersection(bags[2]))

prioritize = lambda x: sum([ord(l)-96 if l == l.lower() else ord(l)-38 for l in x])
p1 = prioritize(L)
p2 = prioritize(L2)
print(f'part1: {p1}\npart2: {p2}')
