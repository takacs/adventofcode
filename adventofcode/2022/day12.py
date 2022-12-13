from collections import deque
from copy import deepcopy


def bfs(inp, orig, part=1):
    Q = deque()
    for x in range(len(inp)):
        for y in range(len(inp[0])):
            if (part == 1 and orig[x][y] == 'S') or (part == 2 and inp[x][y] == 1):
                Q.append(((x, y), 0))
    visited = set()
    while Q:
        (x, y), d = Q.popleft()
        if (x, y) in visited:
            continue
        visited.add((x, y))
        if orig[x][y] == 'E':
            return d
        for dx, dy in [(-1, 0), (0, 1), (1, 0), (0, -1)]:
            xdx = x + dx
            ydy = y + dy
            if 0 <= xdx < len(inp) and 0 <= ydy < len(inp[0]) and inp[xdx][ydy]<=1+inp[x][y]:
                Q.append(((x + dx, y + dy), d + 1))


def main():
    inp = list(map(lambda x: list(x.strip()), open('day12.in').readlines()))
    new = deepcopy(inp)
    for i in range(len(inp)):
        for j in range(len(inp[0])):
            if inp[i][j] == 'S':
                new[i][j] = 1
            elif inp[i][j] == 'E':
                new[i][j] = 26
            else:
                new[i][j] = ord(inp[i][j]) - ord('a') + 1

    p1 = bfs(new, inp, 1)
    p2 = bfs(new, inp, 2)
    print(f'part1: {p1}\npart2: {p2}')


if __name__ == '__main__':
    main()
