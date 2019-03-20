import os

days = ['day01.py', 'day02.py', 'day03.py', 'day04.py', 'day05.py', 'day06.py', 'day07.py', 'day08.py',
        'day09.py', 'day10.py', 'day11.py', 'day12.py', 'day13.py', 'day14.py', 'day15.py', 'day16.py',
        'day17.py', 'day18.py', 'day19.py', 'day20.py', 'day21.py', 'day22.py', 'day23.py', 'day24.py',
        'day25.py']

years = ['2015', '2016', '2017', '2018']

d = {}

for year in years:
    direct = os.listdir('./'+year)
    d[year] = sum([1 for day in direct if day in days ])

print('Advent of Code progress')
for k, v in d.items():
    print(k + ': ', end='')
    print('[' + v*'#', end='')
    print((25-v)*'.' + '] ')


print()
print('Ovrl: ', end='')
print('[' + sum(d.values())//4*'#', end='')
print((25-sum(d.values())//4)*'.' + ']')
