def sumrange(a,b):
    return sum([ i for i in range(a,b+1) if i%2 == 1])


print(sumrange(100,200))
print(sumrange(4291,8955))
