MAXSIZE = 100000
DISK_SIZE = 70000000
NEEDED_SPACE = 30000000


class Directory:
    def __init__(self, name, parent=None, size=None):
        self._parent = parent
        self.name = name
        self.children = {}
        self._size = size

    def insert(self, k, size=None):
        if k not in self.children.keys():
            self.children[k] = Directory(k, self, size)

    @property
    def isDir(self):
        return False if self._size else True

    @property
    def parent(self):
        return self._parent if self._parent else self

    @property
    def size(self):
        return self._size or sum([val.size for key, val in self.children.items()])


def sum_below_max(pwd, count=0, max_size=MAXSIZE):
    if pwd.isDir and pwd.size <= max_size:
        count += pwd.size
    return count + sum([sum_below_max(child) for child in pwd.children.values()])


def min_over(pwd, needed_space):
    size = DISK_SIZE
    if pwd.isDir and pwd.size >= needed_space:
        size = pwd.size
    return min([size] + [min_over(child, needed_space) for child in pwd.children.values()])


def main():
    inp = list(map(lambda x: x.strip(), open('day07.in').readlines()))

    pwd = Directory('/')
    root = pwd
    i = 1
    while i < len(inp):
        if 'cd' in inp[i]:
            cd = inp[i].split()[-1]
            if cd == '..':
                pwd = pwd.parent
            else:
                pwd.insert(cd)
                pwd = pwd.children[cd]
            i += 1
        if 'ls' in inp[i]:
            i += 1
            while i < len(inp) and inp[i][0] != '$':
                if 'dir' in inp[i]:
                    cd = inp[i].split()[-1]
                    pwd.insert(cd)
                else:
                    size, filename = inp[i].split()
                    pwd.insert(filename, size=int(size))
                i += 1

    p1 = sum_below_max(root, 0)
    space_diff = abs(root.size - DISK_SIZE + NEEDED_SPACE)
    p2 = min_over(root, space_diff)

    print(f'part1: {p1}\npart2: {p2}')


if __name__ == '__main__':
    main()
