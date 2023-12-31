package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Velocity struct {
	x, y, z int64
}

type Position struct {
	x, y, z int64
}

type Line struct {
	v Velocity
	p Position
}

func (l Line) a() int64 {
	return l.v.y
}

func (l Line) b() int64 {
	return -1 * l.v.x
}

func (l Line) c() int64 {
	return l.v.y * l.p.x - l.v.x * l.p.y
}

func toi(s string) int64 {
	v, _ := strconv.Atoi(s)
	return int64(v)
}


func main() {
	f, err := os.ReadFile("day24.test")
	if err != nil {
		fmt.Print(err)
	}
	fs := string(f)
	fs = strings.ReplaceAll(fs, ",", "")

	lines := []Line{}	
	for _, l := range strings.Split(fs, "\n") {
		if len(l) < 1 {
			continue
		}
		p := strings.Split(l, "@")
		ps := strings.Fields(p[0])
		pos := Position{x: toi(ps[0]), y:toi(ps[1]), z:toi(ps[2])} 
		vs := strings.Fields(p[1])
		vel := Velocity{x: toi(vs[0]), y:toi(vs[1]), z:toi(vs[2])}
		lines = append(lines, Line{p: pos, v:vel})
	}

	for i := 0; i < len(lines); i++ {
		for j := i+1; j < len(lines); j++ {
			l1 := lines[i]
			l2 := lines[j]
			if l2.a() * l1.b()  == l2.b() * l1.a() {
				continue
			}
			x := (l1.c()*l2.b() - l2.c()*l1.b()) / (l1.a() * l2.b() - l2.a() * l1.b())
			y := (l2.c()*l1.a() - l1.c()*l2.a()) / (l1.a() * l2.b() - l2.a() * l1.b())
			fmt.Println(x,y)
		}
	}
}
