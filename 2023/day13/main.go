package main

import (
    "fmt"
    "os"
    "strings"
)

func transpose(matrix [][]rune) [][]rune {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return matrix
    }

    numRows, numCols := len(matrix), len(matrix[0])

    result := make([][]rune, numCols)
    for i := range result {
        result[i] = make([]rune, numRows)
    }

    for i := 0; i < numRows; i++ {
        for j := 0; j < numCols; j++ {
            result[j][i] = matrix[i][j]
        }
    }

    return result
}



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

func isReflection(rm [][]rune, m, h int) bool {
	li, ri := m-1, m
	diff := 0
	for li >= 0 && ri < len(rm) {
		diff += hamming(string(rm[li]), string(rm[ri]))
		li--
		ri++
	}
	return diff == h
}

func hamming(str1, str2 string) int {
	distance := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			distance++
		}
	}
	return distance
}

func main() {
    f, err := os.ReadFile("day13.in")
    if err != nil {
	    fmt.Print(err)
    }
    fs := string(f)
    m := strings.Split(fs, "\n\n")
   
	for x := 0; x < 2; x++ {
		tot := 0
		rm := [][]rune{}
		N, M := 0, 0
		cr := 0
		for _, mg := range m {
			cr = 0
			rm = toRuneMap(mg)
			N = len(rm)
			M = len(rm[0])
			for i := 1; i < N; i++ {
				if isReflection(rm, i, x) {
					cr += i*100
					break
				}
			}
			tot += cr
			if cr > 0 {
				continue
			}

			rm = transpose(rm)
			for i := 1; i < M; i++ {
				if isReflection(rm, i, x) {
					cr += i
					break
				}
			}
			tot += cr
		}
		fmt.Printf("Part %v: %v\n", x+1, tot)
	}

    
}
