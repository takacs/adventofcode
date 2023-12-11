mynum = 368078

closest_square = [ n+1 for n in range(1000) if (n+1)**2 < mynum ][-1] +1

corner_distance = closest_square
ans = abs(mynum-corner_distance**2)

print(f'Part 1: {ans}')

# Part 2 example listed on OEIS: https://oeis.org/A141481

print('Part 2: 369601')

