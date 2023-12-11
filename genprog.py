import os

days = [ f'day{i:02d}.py' for i in range(1,26) ]

years = ['2015', '2016', '2017', '2018', '2019', '2020']

d = {}

for year in years:
    direct = os.listdir('./' + year)
    d[year] = sum([ 1 for day in direct if day in days ])

print('Advent of Code progress')

for k, v in d.items():
    print(k + ': ', end='')
    print('[' + v * '#', end='')
    print((25-v) * '.' + '] ')

print()
print('Ovrl: ', end='')
print('[' + sum(d.values()) // 4 * '#', end='')
print((25 - sum(d.values())//4) * '.' + ']')
