from string import ascii_lowercase
import sys

sys.setrecursionlimit(1500)

with open('day5in.py', 'r') as f:
    fl = f.read().strip()

poly = [ c+c.upper() for c in ascii_lowercase ]
poly += [ c.upper()+c for c in ascii_lowercase ]

def replacepolys(poly, fn):
    replc = 0
    for p in poly:
        if p in fn:
            replc += 1
            fn = fn.replace(p, '')
    if replc == 0:
        return fn
    else:
        return replacepolys(poly, fn)

def impoly(poly, f):
    react = []
    for c in ascii_lowercase:
        print(c)
        cf = f.replace(c, '')
        cf = cf.replace(c.upper(), '')
        print(len(replacepolys(poly, cf)))

print(len(replacepolys(poly, fl)))
print(len(fl))

print(impoly(poly, fl))

