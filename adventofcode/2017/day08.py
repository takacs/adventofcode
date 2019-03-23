import operator

inp = open('day08in.txt', 'r').read().strip().split('\n')


ops = {'==':operator.eq, '!=': operator.ne, '>=': operator.ge, '<=':operator.le,
       '>': operator.gt, '<': operator.lt}

d = {}

maxs = []
for line in inp:
    a, op, val, _, b, con, cval = line.split(' ')
    if ops[con](d.get(b, 0), int(cval)):
        if op == 'inc':
            d[a] = d.get(a, 0) + int(val)
        else:
            d[a] = d.get(a, 0) - int(val)
        maxs.append(d.get(a))


print(f'Part 1: {max(d.values())}')
print(f'Part 2: {max(maxs)}')
