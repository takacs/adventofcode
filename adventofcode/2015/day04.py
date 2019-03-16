import hashlib

def solve(word, n=5):

    for i in range(0, 100000000):
        m = '{}{}'.format(word, i)
        h = hashlib.md5(m.encode('ascii')).hexdigest()
        if h[:n] == '0'*n:
            return i

    return -1


p1 = solve('ckczppom')
print(f'Part 1: {p1}')
p2 = solve('ckczppom', 6)
print(f'Part 1: {p2}')