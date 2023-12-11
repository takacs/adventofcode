from recordtype import recordtype
from collections import Counter

road = open('wot.txt').read().split('\n')

Cart = recordtype('Cart', ['x', 'y', 'dir', 'iter'])
carts = []
ogroad = [ list(x) for x in road ]
for i,r in enumerate(ogroad):
    for j,c in enumerate(r):
        if c == '>' or c == '<':
            carts.append(Cart(i, j, c, 0))
            ogroad[i][j] = '-'
        if c == '^' or c == 'v':
            ogroad[i][j] = '|'
            carts.append(Cart(i, j, c, 0))

rdict = {'-':'>', '\\':'v', '/':'^'}
ldict = {'-':'<', '\\':'^', '/':'v'}
ddict = {'|':'v', '\\':'>', '/':'<'}
udict = {'|':'^', '\\':'<', '/':'>'}

rt = ['^', '>', 'v']
lt = ['v', '<', '^']
dt = ['>', 'v', '<']
ut = ['<', '^', '>']

def hitcheck(carts):
    ind = [ str(c.y) + ',' + str(c.x) for c in carts ]
    if len(set(ind)) != len(ind):
        return True
    return False

def locatecrash(carts):
    ind = [ str(c.y) + ',' + str(c.x) for c in carts ]
    count = Counter(ind)
    ret = []
    for key, val in count.items():
        if val >= 2:
            print(val)
            ret.append(key)
    return ret

for i in range(15000000000):
    print('+'*15)
    for c in carts:
        direct = c.dir
        if direct == '>':
            c.y += 1
            try:
                c.dir = rdict[ogroad[c.x][c.y]]
            except KeyError:
                c.dir = rt[c.iter%3]
                c.iter += 1
        if direct == '<':
            c.y -= 1
            try:
                c.dir = ldict[ogroad[c.x][c.y]] 
            except KeyError:
                c.dir = lt[c.iter%3]
                c.iter += 1
        if direct == '^':
            c.x -= 1
            try:
                c.dir = udict[ogroad[c.x][c.y]] 
            except KeyError:
                c.dir = ut[c.iter%3]
                c.iter += 1
        if direct == 'v':
            c.x += 1
            try:
                c.dir = ddict[ogroad[c.x][c.y]] 
            except KeyError:
                c.dir = dt[c.iter%3]
                c.iter += 1
    # part 1
    if hitcheck(carts):
        print('HIT!!: ')
        hit = locatecrash(carts)
        print(hit)
        break

    # part 2
   # if hitcheck(carts):
   #     hit = locatecrash(carts)
   #     hits = [ h.split(',') for h in hit ]
   #     hints = [ [int(x), int(y)] for x, y in hits ]
   #     newcarts = []
   #     for c in carts:
   #         if [c.y, c.x] not in hints:
   #             newcarts.append(c)
   #     carts = newcarts
   #     if len(carts) == 1:
   #         print(carts)
   #         break
   # for d in carts:
   #     print(d)

#for line in ogroad:
#    print("".join(line))
