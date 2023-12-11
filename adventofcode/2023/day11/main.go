package main

import (
	"bufio"
	"fmt"
	"os"
)

type C struct {
    x,y int
}

func (c C) distance(c2 C, rs, cs []int, mx int) int {
    xm := 0
    for _, ro := range rs {
	if between(c.y, c2.y, ro) {
	    xm += mx-1
	}
    }
    for _, co := range cs {
	if between(c.x, c2.x, co) {
	    xm += mx-1
	}
    }

    return manhattan(c,c2) + xm
}

func manhattan(c, c2 C) int {
    x := c.x - c2.x
    y := c.y - c2.y
    if x < 0 {
	x *= -1
    }
    if y < 0 {
	y *= -1
    }
    return x+y
}

func between(a, b, x int) bool {
    lo, hi := a, b
    if lo > hi {
	lo, hi = b, a
    }
    return lo < x && x < hi
}

func transpose(matrix [][]rune) [][]rune {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]rune, cols)
	for i := range transposed {
		transposed[i] = make([]rune, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func alldot(r []rune) bool {
    for _, c := range r {
	if c != '.' {
	    return false
	}
    } 
    return true
}

func extract(m [][]rune) []int {
    x := []int{}
    for i, r := range m {
	if alldot(r) {
	    x = append(x, i)
	}
    }
    return x
}

func getG(m [][]rune) []C{
    cs := []C{}
    for i := 0; i < len(m); i++ {
	for j := 0; j < len(m[i]); j++ {
	    if m[i][j] == '#' {
		cs = append(cs, C{x:j, y:i})
	    }
	}
    }
    return cs
}

func main() {
    file, err := os.Open("day11.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    
    m := [][]rune{}
    for scanner.Scan() {
        line := []rune(scanner.Text())
	m = append(m, line)
    }

    rs := extract(m)
    m = transpose(m)
    cs := extract(m)

    gs := getG(m)
    fmt.Println(gs)
    fmt.Println(rs,cs)
    
    s := 0
    s2 := 0
    for i := 0; i < len(gs); i++ {
	for j := i+1; j < len(gs); j++ {
	    d := gs[i].distance(gs[j], cs, rs, 2)
	    fmt.Println(gs[i], gs[j], d)
	    s += d
	    s2 += gs[i].distance(gs[j], cs, rs, 1000000)
	}
    }

    fmt.Printf("Part 1: %v\n", s)
    fmt.Printf("Part 2: %v\n", s2)

}

