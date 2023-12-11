import sys

def twoSum(vals, target):
    d = {}
    for val in vals:
        if val in d.keys():
            return val * (2020 - val)
        else:
            d[2020-val] = None
    return None
    
def threeSum(vals, target):
    size = len(vals)
    for i in range(0, size):
        for j in range(i+1, size):
            for k in range(j+1, size):
                currSum = vals[i] + vals[j] + vals[k]
                if currSum == 2020:
                    return vals[i] * vals[j] * vals[k] 
    return None

nums = list(map(int,open(sys.argv[1]).read().strip().split()))
print(twoSum(nums, 2020))
print(threeSum(nums, 2020))
