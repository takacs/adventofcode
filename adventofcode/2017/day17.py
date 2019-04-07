from collections import deque

step = 354

lock = deque([0])

for val in range(1, (5*10**7)+1):
    lock.rotate(-step)
    lock.append(val)

print(lock[lock.index(0)+1])
