import fileinput

def countAny():
    count = 0
    ans = []
    for line in fileinput.input():
        line = line.strip()
        if len(line) == 0:
            count += len(set(ans))
            ans = []
        else:
            ans.extend([char for char in line])
    return count

def countAll():
    count = 0
    ans = []
    people = 0
    for line in fileinput.input():
        line = line.strip()
        if len(line) == 0:
            for char in set(ans):
                if ans.count(char) == people:
                    count+=1
            people = 0
            ans = []
        else:
            people += 1
            ans.extend([char for char in line])
    return count

print(countAny())
print(countAll())
