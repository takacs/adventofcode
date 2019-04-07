inp = open('day16in.txt', 'r').read().strip().split(',')

ps = [
        'a', 'b', 'c', 'd', 'e',
        'f', 'g', 'h', 'i', 'j',
        'k', 'l', 'm', 'n', 'o',
        'p'
        ]


#for d in inp:
#    w = d[0]
#    if w == 's':
#        swap = int(d[1:])
#        ps = ps[-swap:]+ps[:-swap]
#    if w == 'x':
#        a,b = map(int, d[1:].split('/'))
#        ps[a], ps[b] = ps[b], ps[a]
#    if w == 'p':
#        a,b = d[1:].split('/')
#        a,b = ps.index(a), ps.index(b)
#        ps[a], ps[b] = ps[b], ps[a]
#
#print(''.join(ps))

pos = []

for i in range(10**9):
    s = ''.join(ps) 
    if s in pos:
        print(pos[10**9%i])
        break

    pos.append(s)
    for d in inp:
        w = d[0]
        if w == 's':
            swap = int(d[1:])
            ps = ps[-swap:]+ps[:-swap]
        if w == 'x':
            a,b = map(int, d[1:].split('/'))
            ps[a], ps[b] = ps[b], ps[a]
        if w == 'p':
            a,b = d[1:].split('/')
            a,b = ps.index(a), ps.index(b)
            ps[a], ps[b] = ps[b], ps[a]
