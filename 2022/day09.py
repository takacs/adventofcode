import math

import pandas as pd
import numpy as np

MOVE_HEAD = {
    'U': (0, 1),
    'D': (0, -1),
    'L': (-1, 0),
    'R': (1, 0)
}

MOVE_TAIL = {
    (0, 0): {'L': (0, 0), 'R': (0, 0), 'U': (0, 0), 'D': (0, 0)},
    (0, 1): {'L': (0, 0), 'R': (0, 0), 'U': (0, 1), 'D': (0, 0)},
    (1, 1): {'L': (0, 0), 'R': (1, 1), 'U': (1, 1), 'D': (0, 0)},
    (1, 0): {'L': (0, 0), 'R': (1, 0), 'U': (0, 0), 'D': (0, 0)},
    (1, -1): {'L': (0, 0), 'R': (1, -1), 'U': (0, 0), 'D': (1, -1)},
    (0, -1): {'L': (0, 0), 'R': (0, 0), 'U': (0, 0), 'D': (0, -1)},
    (-1, -1): {'L': (-1, -1), 'R': (0, 0), 'U': (0, 0), 'D': (-1, -1)},
    (-1, 0): {'L': (-1, 0), 'R': (0, 0), 'U': (0, 0), 'D': (0, 0)},
    (-1, 1): {'L': (-1, 1), 'R': (0, 0), 'U': (-1, 1), 'D': (0, 0)}
}


def place(inp, n=2):
    # Solution only works for n=2 :)
    n = 2
    head = [0, 0]
    tail = [0, 0]
    visited = set()
    for line in inp:
        d, nr = line.split()
        for i in range(int(nr)):
            vkey = (head[0] - tail[0], head[1] - tail[1])
            head = [MOVE_HEAD[d][0] + head[0], MOVE_HEAD[d][1] + head[1]]
            tail = [MOVE_TAIL[vkey][d][0] + tail[0], MOVE_TAIL[vkey][d][1] + tail[1]]
            visited.add((tail[0], tail[1]))

    return len(visited)

def main():
    inp = list(map(lambda x: x.strip(), open('day09.in').readlines()))

    p1 = place(inp, 2)
    p2 = place(inp, 10)
    print(f'part1: {p1}\npart2: {p2}')


if __name__ == '__main__':
    main()
