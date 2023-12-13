package main

import (
    "fmt"
    "os"
    "strings"
)

func transpose(matrix [][]rune) [][]rune {
	fmt.Println("M->", matrix)
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return matrix
    }

    numRows, numCols := len(matrix), len(matrix[0])
	fmt.Println(numRows, numCols)

    result := make([][]rune, numCols)
    for i := range result {
        result[i] = make([]rune, numRows)
    }

    for i := 0; i < numRows; i++ {
        for j := 0; j < numCols; j++ {
            result[j][i] = matrix[i][j]
        }
    }
	fmt.Println("END")

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

func isReflection(rm [][]rune, m int) bool {
	li, ri := m-1, m
	for li >= 0 && ri < len(rm) {
		if string(rm[li]) != string(rm[ri]) {
			return false
		}
		li--
		ri++
	}
	return true

}

func main() {
    f, err := os.ReadFile("day13.in")
    if err != nil {
	    fmt.Print(err)
    }
    fs := string(f)
    m := strings.Split(fs, "\n\n")
   
	tot := 0
	rm := [][]rune{}
	N, M := 0, 0
	for _, mg := range m {
		rm = toRuneMap(mg)
		N = len(rm)
		M = len(rm[0])
		for i := 1; i < N; i++ {
			if isReflection(rm, i) {
				fmt.Printf("Validrow: %d\n", i)
				tot += i*100
				break
			}
		}

		rm = transpose(rm)
		for i := 1; i < M; i++ {
			if isReflection(rm, i) {
				fmt.Printf("Validcol: %d\n", i)
				tot += i
				break
			}
		}
	}


	fmt.Println(tot)
    
}
