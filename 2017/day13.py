# I hate this code

inp = {
0: 3  ,
1: 2  ,
2: 6  ,
4: 4  ,
6: 4  ,
8: 8  ,
10: 6 ,
12: 8 ,
14: 5 ,
16: 6 ,
18: 8 ,
20: 8 ,
22: 12,
24: 6 ,
26: 9 ,
28: 8 ,
30: 12,
32: 12,
34: 17,
36: 12,
38: 8 ,
40: 12,
42: 12,
44: 10,
46: 12,
48: 12,
50: 12,
52: 14,
54: 14,
56: 10,
58: 14,
60: 12,
62: 14,
64: 14,
66: 14,
68: 14,
70: 14,
72: 14,
74: 14,
76: 14,
86: 14,
94: 20,
96: 18,
}
test = {
0: 3,
1: 2,
4: 4,
6: 4,
        }
from itertools import cycle


j = 0
total = -1

caught = True

while caught:
    
    caught=False
    j+=1

    d = { k : cycle(list(range(v)) + list(range(v))[1:-1][::-1]) for k, v in inp.items() }

    for i in range(j+1):
        nums = { k:next(v) for k, v in d.items() }
    
    total = 0
    for i in range(max(d.keys())+1):
        if i in nums.keys():
            if nums[i] == 0:
                caught = True
                total += i*inp[i]
        nums = { k:next(v) for k, v in d.items() }

print(j)
