package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Game struct {
    s string
    n []int
}


func s2i(s string) []int {
    sp := strings.Split(s, ",")
    i := []int{}
    for _, c := range sp {
	val, _ := strconv.Atoi(c)
	i = append(i, val)
    }
    return i
}

var DP = map[string]int{}
func f(dots string, blocks []int, i, bi, cur int) int {
    key := string(i)+"_"+string(bi)+"_"+string(cur)
    _, exists := DP[key]
    if exists {
        return DP[key]
    }
    if i == len(dots) {
	if bi == len(blocks) && cur==0 {
	    return 1
	} else if bi==len(blocks)-1 && blocks[bi] == cur {
	    return 1
	} else {
	    return 0
	}
    }
    ans := 0
    for _, c := range ".#" {
	if string(dots[i]) == string(c) || dots[i] == '?'{
	    if c == '.' && cur == 0 {
		ans += f(dots, blocks, i+1, bi, 0)
	    } else if c == '.' && c > 0 && bi < len(blocks) && blocks[bi] == cur {
		ans += f(dots, blocks, i+1, bi+1, 0)
	    } else if c == '#' {
		ans += f(dots, blocks, i+1, bi, cur+1)
	    }
	}
    }
    DP[key] = ans
    return ans
}

func rList(list []int, n int) []int {
    repeated := make([]int, 0, len(list)*n)
    for i := 0; i < n; i++ {
        repeated = append(repeated, list...)
    }
    return repeated
}

func rString(s string, n int) string {
    sr := ""
    for i := 0; i < n; i++ {
        sr += s + "?"
    }
    return sr[:len(sr)-1]
}

func main() {
    file, err := os.Open("day12.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    
    gs := []Game{}
    for scanner.Scan() {
        line := scanner.Text()
	ps := strings.Fields(line)
	gs = append(gs, Game{s:ps[0], n:s2i(ps[1])})
    }
    
    count := 0
    for _, g := range gs {
	DP = map[string]int{}
	count += f(g.s, g.n, 0, 0, 0)
    }
    fmt.Printf("Part 1: %v\n", count)

    count = 0
    for _, g := range gs {
	DP = map[string]int{}
	count += f(rString(g.s, 5), rList(g.n, 5) , 0, 0, 0)
    }
    fmt.Printf("Part 2: %v\n", count)

}

