from collections import defaultdict

d = defaultdict(list)

with open('day12in.txt', 'r') as f:
    for line in f:
        l = line.split(' <-> ')
        head = int(l[0])
        for p in l[1].replace(' ', '').split(','):
            d[head].append(int(p))

first = True
vis = set()
ans = 0
for i in range(2000):
    if i in vis:
        continue
    
    ans += 1
    q = [i]
    while q:
        for b in d[q.pop()]:
            if b not in vis:
                q.append(b) 
            vis.add(b)
    if first:
        first = False
        print(len(vis))
print(ans)
