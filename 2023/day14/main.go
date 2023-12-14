package main

import (
	"fmt"
	"math"
	"os"
	"strings"

)

func toRuneMap(s string) [][]rune {
	rc := []rune{}
	rm := [][]rune{}
	for _, st := range strings.Split(s, "\n") {
		rc = []rune{}
		for _, r := range st {
			rc = append(rc, r)

		}
		rm = append(rm, rc)
	}
	return rm
}

// Used for part1 originally
func nload(m [][]rune) int {
    N := len(m)
    M := len(m[0])
    tot := 0
    for j := 0; j < M; j++ {
	colval := M
	for i := 0; i < N; i++ {
	    if m[i][j] == 'O' {
		tot += colval
		colval--
	    }
	    if m[i][j] == '#' {
		colval = N-i-1
	    }
	}
    }

    return tot
}

func r2s(m [][]rune) string {
    s := ""
    for _, rm := range m {
	s += string(rm)
    }
    return s
}

func rotate(m [][]rune) [][]rune {
	rows := len(m)
	cols := len(m[0])

	rotated := make([][]rune, cols)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-i-1] = m[i][j]
		}
	}

	return rotated
}

func roll(m [][]rune) [][]rune {
	for i := 1; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			for k := i - 1; k >= 0; k-- {
				if m[k][j] == '.' && m[k+1][j] == 'O' {
					m[k][j] = 'O'
					m[k+1][j] = '.'
				}
			}
		}
	}
	return m
}

func printr(m [][]rune) {
	for _, r := range m {
		fmt.Println(string(r))
	}
}

func score(m [][]rune) int {
	tot := 0
	R := len(m)
	C := len(m[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if m[r][c] == 'O' {
				tot += R-r
			}
		}
	}
	return tot
}
    
func main() {
    f, err := os.ReadFile("day14.in")
    if err != nil {
	    fmt.Print(err)
    }
    fs := string(f)
    m := [][]rune{}
    for _, rm := range strings.Split(fs, "\n") {
	a := []rune(rm)
	if len(a) > 0 {
	    m = append(m, []rune(rm))
	}
    }



    XX := map[string]int{}
	target := 1000000000
    for i := 0; i < target; i++ {
		for k := 0; k < 4; k++ {
			m = roll(m)
			if k==0 && i == 0 {
				fmt.Printf("Part 1: %v\n", score(m))
			}
			m = rotate(m)
		}
		x := r2s(m)
		t, exists := XX[x]
		if exists {
			cycle := i-t
			skip := int(math.Floor(float64(target-i)/float64(cycle)))
			i += skip * cycle
		}
		XX[x] = i
    }
	fmt.Printf("Part 2: %v\n", score(m))
}
	
