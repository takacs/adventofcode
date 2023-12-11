from itertools import repeat
from itertools import islice
from functools import reduce

nums = """
102,255,99,252,200,24,219,57,103,2,226,254,1,0,69,216
""".strip()
inp = list(map(int, nums.split(',')))
cp, skip = 0, 0

ls = list(range(0,256))
l = len(ls)

for length in inp:
    end = cp + length
    if end < l:
        ls[cp:end] = ls[cp:end][::-1]
    else:
        lop = (end)%l
        temp = (ls[cp:] + ls[:lop])[::-1]
        fir, sec = temp[len(ls[cp:]):], temp[:len(ls[cp:])]
        ls[cp:] = sec
        ls[:lop] = fir

    cp =  (cp + length + skip)%l
    skip += 1

    
print(f'Part 1: {ls[0]*ls[1]}')

inp = list(map(ord, list(nums)))
inp += [17, 31, 73, 47, 23]
cp, skip = 0, 0

ls = list(range(0,256))
l = len(ls)

for lengths in repeat(inp,64):
    for length in lengths:
        end = cp + length
        if end < l:
            ls[cp:end] = ls[cp:end][::-1]
        else:
            lop = (end)%l
            temp = (ls[cp:] + ls[:lop])[::-1]
            fir, sec = temp[len(ls[cp:]):], temp[:len(ls[cp:])]
            ls[cp:] = sec
            ls[:lop] = fir

        cp =  (cp + length + skip)%l
        skip += 1


def chunk(it, size):
    it = iter(it)
    return iter(lambda: tuple(islice(it, size)), ())

hexval = []
for ch in chunk(ls, 16):
    val = reduce(lambda a,b: a^b, ch)

    hexval.append(f'{val:02x}')

hexval = ''.join(hexval)
print(f'Part 2: {hexval}')
