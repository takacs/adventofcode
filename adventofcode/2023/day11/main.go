package main

import (
    "fmt"
    "os"
    "bufio"
)

type C struct {
    x,y int
}

func (c C) distance(c2 C) int {
    dx := c.x - c2.x
    dy := c.y - c2.y
    if dx < 0 {
	dx = dx*-1
    }
    if dy < 0 {
	dy = dy*-1
    }
    return dx+dy
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

func extend(m [][]rune) [][]rune {
    rr := [][]rune{}
    for _, r := range m {
	rr = append(rr, r)
	if alldot(r) {
	    rr = append(rr, r)
	}
    }
    return rr
}

func getG(m [][]rune) []C{
    cs := []C{}
    for i := 0; i < len(m); i++ {
	for j := 0; j < len(m[i]); j++ {
	    if m[i][j] == '#' {
		cs = append(cs, C{x:i, y:j})
	    }
	}
    }
    return cs
}

func main() {
    file, err := os.Open("day11.test")
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

    m = extend(m)
    m = transpose(m)
    m = extend(m)

    gs := getG(m)
    fmt.Println(gs)
    
    s := 0
    for i := 0; i < len(gs); i++ {
	for j := i+1; j < len(gs); j++ {
	    s += gs[i].distance(gs[j])
	}
    }

    fmt.Printf("Part 1: %v\n", s)

}

