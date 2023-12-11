def main():
    inp = list(map(lambda x: x.strip(), open('day10.in').readlines()))
    isLit = lambda x, crt: '#' if abs(x - int(crt % 40)) < 2 else '.'

    X = 1
    current_goal = 20
    signal_strength = 0
    screen = [[0 for x in range(40)] for y in range(6)]
    crt = 0
    for cmd in inp:
        if cmd == 'noop':
            screen[int(crt / 40)][crt % 40] = isLit(X, crt)
            crt += 1
        if 'add' in cmd:
            val = int(cmd.split()[1])
            if crt + 2 >= current_goal:
                signal_strength += current_goal * X
                current_goal += 40
            for i in range(2):
                screen[int(crt / 40)][crt % 40] = isLit(X, crt)
                crt += 1
            X += val

    p1 = signal_strength
    p2 = '\n' + '\n'.join(''.join(map(str, row)) for row in screen)
    print(f'part1: {p1}\npart2: {p2}')


if __name__ == '__main__':
    main()
