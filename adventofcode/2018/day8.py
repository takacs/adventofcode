import sys
sys.setrecursionlimit(15000)

inp = open('day8in.txt', 'r').read().strip().split(' ')
inp = [int(n) for n in inp]

class TreeData:
    curr = -1

    def __init__(self, data):
        self.data = data

    def next(self):
        self.curr += 1
        return self.data[self.curr]

tree = TreeData(inp)

def readtree():
    cn, mn = tree.next(), tree.next()
    children = []
    metadata = []
    
    for _ in range(cn):
        children.append(readtree())
    for _ in range(mn):
        metadata.append(tree.next())

    return(children, metadata)

def sumtree(root):
    ans = 0
    for m in root[1]:
        ans += m
    for c in root[0]:
        ans += sumtree(c)
    return ans

def rootval(root):
    ans = 0
    if len(root[0]) == 0:
        return sum(root[1])
    for m in root[1]:
        try:
            ans += rootval(root[0][m-1])
        except IndexError:
            continue
    return ans

root = readtree()
print(sumtree(root))
print(rootval(root))

