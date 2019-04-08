def rabbitfib(n,k,t,r):
    if n == 1:
        return r+t
    return rabbitfib(n-1, k, r*k, r+t)

print(rabbitfib(5,3,1,0))
print(rabbitfib(29,5,1,0))

