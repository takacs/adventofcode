import re

inp = open('day10in.txt', 'r').read().strip('').split('\n')

lights = []

for l in inp[:-1]:
    line = re.findall('<(.*?)>', l)
    ps = line[0].split(',')
    pos = [int(ps[0]), int(ps[1])]
    vs = line[1].split(',')
    vel = [int(vs[0]), int(vs[1])]
    lights.append([pos,vel])

def getvals(lights):
    mx = max([x for (x,y),(vx,vy) in lights ])
    my = max([y for (x,y),(vx,vy) in lights ])

    nx = min([x for (x,y),(vx,vy) in lights ])
    ny = min([y for (x,y),(vx,vy) in lights ])

    return mx, my, nx, ny

j = 0
while True:
    maxx, maxy, minx, miny = getvals(lights) 
    
    if abs(miny)+abs(maxy) < 300:
        print(j)
        
        sky = [['.' for i in range(abs(maxx)+abs(minx)+1)] for j in range(abs(maxy)+abs(miny)+1)]
        for (x,y),(vx,vy) in lights:
            sky[y][x] = '#'
        for line in sky:
            for c in line:
                print(c, end='')
            print('\n')
    for i,((x,y),(vx,vy)) in enumerate(lights):
        lights[i][0][0] += vx
        lights[i][0][1] += vy
    j+= 1
