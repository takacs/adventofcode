import fileinput

def getPasses():
    passes = []
    for line in fileinput.input():
        line = line.strip()
        l, r  = 0, 127
        lc, rc = 0, 7
        row = line[:7]
        col = line[-3:-1]
        for c in row:
            dist = (r-l)//2
            if c == 'F':
                r -=dist+1
            if c == 'B':
                l += dist+1
        for c in col:
            dist = (rc-lc)//2
            if c == 'R':
                lc += dist+1
            if c == 'L':
                rc -= dist+1
        if line[-1] == 'R':
            cr = rc
        else:
            cr = lc
        sid = r*8 + cr
        passes.append(sid)
    
    return passes
    return m

def findMissing(passes):
    for val in range(0, max(passes)):
        if val not in passes and val > 7:
            return val

passes = getPasses()
print(max(passes))
print(findMissing(passes))
