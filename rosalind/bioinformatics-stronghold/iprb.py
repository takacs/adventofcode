from itertools import combinations

def comb(n, k):
    return len(list(combinations(range(n),k)))

def domprob(k,m,n):
    tot = 4*comb((k+m+n), 2)
    totrec = 4*comb(n,2) + 2*n*m + comb(m,2)
    return 1-totrec/tot

print(domprob(27,20,22))
