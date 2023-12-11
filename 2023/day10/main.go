package main

import (
    "fmt"
    "os"
    "bufio"
)

type Point struct {
    x,y int
}

type Direction struct {
    x,y int
}


func (p Point) move(d Direction) Point {
    return Point{x:p.x+d.x, y:p.y+d.y}
}

func (p *Point) update(s Point) {
    p.x = s.x
    p.y = s.y
}

func findS(m [][]rune) Point {
    for i:=0; i < len(m); i++ {
	for j:=0; j < len(m[i]); j++ {
	   if m[i][j] == 'S' {
		return Point{x: i, y: j}
	    } 
	}
    }
    return Point{}
}

func vins(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
func main() {
    file, err := os.Open("day10.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    
    p := [][]rune{}
    for scanner.Scan() {
        line := []rune(scanner.Text())
	p = append(p, line)
    }

    s := findS(p)
    dirs := []Direction{{x:1,y:0},{x:0,y:1},
		      {x:-1,y:0},{x:0,y:-1}}
    steps := 0
    sd := false
    su := false
    sl := false
    sr := false
    for true {

	us := s.move(dirs[2])
	rs := s.move(dirs[1])
	ds := s.move(dirs[0])
	ls := s.move(dirs[3])
	if !su && ( p[s.x][s.y] == 'S' || p[s.x][s.y] == 'L' || p[s.x][s.y] == 'J' || p[s.x][s.y] == '|' ) && (p[us.x][us.y] == '|' || p[us.x][us.y] == '7' || p[us.x][us.y] == 'F' || p[us.x][us.y] == 'S') {
	    if steps > 2 && p[us.x][us.y] == 'S' {
		fmt.Println("up", string(p[us.x][us.y]), s)
		steps++
		break
	    } 
	    sd = true 
	    su = false
	    sl = false
	    sr = false
	    s.update(us)
	} else if !sr && ( p[s.x][s.y] == 'S' || p[s.x][s.y] == '-' || p[s.x][s.y] == 'L' || p[s.x][s.y] == 'F' ) && (p[rs.x][rs.y] == '-' || p[rs.x][rs.y] == '7' || p[rs.x][rs.y] == 'J' || p[rs.x][rs.y] == 'S') {
	    if steps > 2 && p[rs.x][rs.y] == 'S' {
		fmt.Println("right", string(p[rs.x][rs.y]), s)
		steps++
		break
	    } 
	    sd = false 
	    su = false
	    sl = true
	    sr = false
	    s.update(rs)
	} else if !sd && ( p[s.x][s.y] == 'S' || p[s.x][s.y] == '|' || p[s.x][s.y] == 'F' || p[s.x][s.y] == '7' ) && (p[ds.x][ds.y] == '|' || p[ds.x][ds.y] == 'J' || p[ds.x][ds.y] == 'L' || p[ds.x][ds.y] == 'S'){
	    if steps > 2 && p[ds.x][ds.y] == 'S' {
		fmt.Println("down", string(p[ds.x][ds.y]), s)
		steps++
		break
	    } 
	    sd = false 
	    su = true
	    sl = false
	    sr = false
	    s.update(ds)
	} else if !sl && ( p[s.x][s.y] == 'S' || p[s.x][s.y] == '7' || p[s.x][s.y] == '-' || p[s.x][s.y] == 'J' ) && (p[ls.x][ls.y] == '-' || p[ls.x][ls.y] == 'L' || p[ls.x][ls.y] == 'F' || p[ls.x][ls.y] == 'S'){
	    if steps > 2 && p[ls.x][ls.y] == 'S' {
		fmt.Println("left", string(p[ls.x][ls.y]), s)
		steps++
		break
	    } 
	    sd = false 
	    su = false
	    sl = false
	    sr = true
	    s.update(ls)
	}

	steps += 1
    }
    fmt.Println(steps/2)
}

