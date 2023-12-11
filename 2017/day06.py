inp = """
11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11
""".strip()

inp = list(map(int, inp.split()))

banks = []
b = inp
ln = len(b)

while b not in banks:
    banks.append(b.copy())
    m = max(b)
    mi = b.index(m)

    for i in range(mi+1, mi+1+b[mi%ln]):
        b[i%ln] += 1

    b[mi] -= m

print(f'Part 1: {len(banks)}')
print(f'Part 2: {len(banks) - banks.index(b)}')
