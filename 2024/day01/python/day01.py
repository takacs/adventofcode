from collections import Counter

x = map(lambda x: x.split(), open("day01.in", "r").readlines())
a = [int(y[0]) for y in x]
x = map(lambda x: x.split(), open("day01.in", "r").readlines())
b = [int(y[1]) for y in x]

print(a, b)
tot = 0

for i, j in zip(sorted(a), sorted(b)):
    tot += abs(i - j)

print(tot)


c = Counter(b)

tot = 0
for i in a:
    tot += i * c.get(i, 0)

print(tot)
