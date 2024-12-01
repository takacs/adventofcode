from collections import Counter

inp = open("day01.in", "r").readlines()

L, R = [], []
for i, v in enumerate(inp):
    vs = v.split()
    L.append(int(vs[0]))
    R.append(int(vs[1]))

s1 = sum(map(lambda x: abs(x[0] - x[1]), zip(sorted(L), sorted(R))))
print(s1)

C = Counter(R)
s2 = sum(map(lambda x: x * C.get(x, 0), L))
print(s2)
