a = 591
b = 393

af = 16807
bf = 48271

total = 0


def makegen(start, fact):
    x = start
    while True:
        x = (x*start)%2147483647
        yield x

def makegenconst(start, fact, mod):
    x = start
    while True:
        x = (x*fact)%2147483647
        if x % mod == 0:
            yield x


#for i in range(4*(10**7)):
#agen = makegen(a, af)
#bgen = makegen(b, bf)

agen = makegenconst(a, af, 4)
bgen = makegenconst(b, bf, 8)
for i in range(5*(10**6)):
    a = next(agen)
    b = next(bgen)
    if f'{a:016b}'[-16:] == f'{b:016b}'[-16:]:
        total += 1

print(total)
