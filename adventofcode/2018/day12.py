inp = open('day12in.txt').read().split('\n')
init = inp[0].split(': ')[-1]
inlen = len(init)

rules = {}
for r in inp[1:-1]:
    r = r.split(' => ')
    rules[r[0]] = r[1]

def next(pre, rules):
    new = ''
    for i in range(len(pre)-4):
        pat = pre[i:i+5]
        new += rules.get(pat, '.')
    
    return '..' + new + '..'

def sumplant(init):
    count = 0
    for i, p in enumerate(init):
        if p == '#':
            count += i-10000
    return count

init = '.'*10000 + init + '.'*10000
    
for t in range(20):
    init = next(init, rules)
    print(t+1, len(init), sumplant(init))

#part2 iternum*21 + 480
print('part2: ', 50000000000*21 + 480)
