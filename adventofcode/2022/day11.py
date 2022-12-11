import collections
import copy
from functools import reduce
import operator


def throw_bananas(monkeys, n, lcm=None):
    monkeys = copy.deepcopy(monkeys)
    op_count = collections.defaultdict(int)
    for i in range(n):
        for monkey in monkeys.keys():
            cmonkey = monkeys[monkey]
            for j in range(len(cmonkey['items'])):
                old = cmonkey['items'].pop()
                new = eval(cmonkey['operation'])
                if lcm:
                    new = new % lcm
                else:
                    new = new // 3
                if new % cmonkey['test'] == 0:
                    monkeys['Monkey ' + cmonkey['true']]['items'].append(new)
                else:
                    monkeys['Monkey ' + cmonkey['false']]['items'].append(new)
                op_count[monkey] += 1

    return reduce(operator.mul, [op_count[monkey]
                                 for monkey in sorted(op_count, key=op_count.get, reverse=True)[:2]])


def main():
    monkeys = {}
    inp = open('day11.in').read()
    lcm = 1
    for monkey in inp.split('\n\n'):
        monkey_data = list(map(lambda x: x.strip(), monkey.split('\n')))
        monkey_info = {'items': list(map(lambda x: int(x.replace(',', '')), monkey_data[1].split()[2:])),
                       'operation': monkey_data[2].split('=')[-1],
                       'test': int(monkey_data[3].split()[-1]),
                       'true': monkey_data[4].split()[-1],
                       'false': monkey_data[5].split()[-1]
                       }
        monkeys[monkey_data[0][:-1]] = monkey_info
        lcm *= monkey_info['test']
    monkeys = collections.OrderedDict(monkeys)

    p1 = throw_bananas(monkeys, n=20)
    p2 = throw_bananas(monkeys, n=10000, lcm=lcm)
    print(f'part1: {p1}\npart2: {p2}')


if __name__ == '__main__':
    main()
