a,b,c,d,e,f = range(6)

turing = {

    (a, 0):(1,1,b),
    (a, 1):(1,-1,e),

    (b, 0):(1,1,c),
    (b, 1):(1,1,f),

    (c, 0):(1,-1,d),
    (c, 1):(0,1,b),

    (d, 0):(1,1,e),
    (d, 1):(0,-1,c),

    (e, 0):(1,-1,a),
    (e, 1):(0,1,d),

    (f, 0):(1,1,a),
    (f, 1):(1,1,c)

        }

step = a
check = 12459852
vals = {}
ind = 0

for i in range(check):
    cval = vals.get(ind, 0)
    wval, mov, cont = turing.get((step,cval))

    vals[ind] = wval
    ind += mov
    step = cont

print(sum(vals.values()))

