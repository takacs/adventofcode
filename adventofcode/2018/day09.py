#p: 400,  m:71864

g = [0,2,1]
cmarb = 1
players  = { p+1:0 for p in range(400)}
cplayer = 1

for m in range(3,71865*100):
    if m%23 != 0:
        cmarb += 2
        if cmarb >= len(g):
            cmarb -= len(g)
        g.insert(cmarb,m)
    else:
        score = m
        extra = cmarb - 7
        if extra < 0:
            extra += len(g)
        score += g[extra]
        g.remove(g[extra])
        cmarb = extra
        players[cplayer] += score
    cplayer += 1
    if cplayer > 400:
        cplayer = 1
    print(round(m/7186500*100,2))

print(max([ v for k,v in players.items()]))
