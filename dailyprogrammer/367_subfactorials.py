# Derangement ~ !n

def derangement(n):
    if n == 0:
        return 1
    if n == 1:
        return 0
    return (n-1)*(derangement(n-1)+derangement(n-2))

assert derangement(5) == 44
assert derangement(6) == 265
assert derangement(9) == 133496
assert derangement(14) == 32071101049
