import numpy as np


def get_visibility_matrix(arr):
    v = []
    for i in range(len(arr)):
        l = []
        lmax = -1
        for j in range(len(arr)):
            if arr[i][j] > lmax:
                l.append(1)
                lmax = int(arr[i][j])
            else:
                l.append(0)
        v.append(l)
    return np.array(v)


def get_vision_matrix(arr):
    v = []
    for i in range(len(arr)):
        l = []
        for j in range(len(arr)):
            vd = 0
            for k in range(j - 1, -1, -1):
                vd += 1
                if arr[i][k] >= arr[i][j]:
                    break
            l.append(vd)
        v.append(l)
    return np.array(v)


def main():
    inp = np.array([list(map(int, list(row)))
                    for row in list(map(lambda x: x.strip(), open('day08.in').readlines()))])
    for i in range(4):
        if not i:
            visibility_mask = get_visibility_matrix(inp)
            vision_mask = get_vision_matrix(inp)
        else:
            visibility_rotated = np.rot90(get_visibility_matrix(np.rot90(inp, k=i)), k=-i)
            vision_rotated = np.rot90(get_vision_matrix(np.rot90(inp, k=i)), k=-i)
            visibility_mask = visibility_mask | visibility_rotated
            vision_mask = vision_mask * vision_rotated

    p1 = np.sum(visibility_mask)
    p2 = np.max(vision_mask)
    print(f'part1: {p1}\npart2: {p2}')


if __name__ == '__main__':
    main()
