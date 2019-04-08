def gc(fn):
    with open(fn, 'r') as f:
        vals = [val.strip() for val in f]
    
    d = {}
    cf = vals[0][1:]
    s = ''

    for val in vals:
        if val[0] != '>':
            s += val
        else:
            d[cf] = s
            s = ''
            cf = val[1:]
    d[cf] = s

    d = { k:((v.count('G')+v.count('C'))/len(v))*100 for k, v in d.items() }

    maxfn, maxgc = '', 0
    for k, v in d.items():
        if v > maxgc:
            maxgc = v
            maxfn = k

    return maxfn, maxgc
        
a = gc('rosalind_gc.txt')
print(a[0])
print(a[1])
