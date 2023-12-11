package main

import (
	"fmt"
	"os"
	"strings"
)

func calculate (s []string, i string, m map[string][]string) int {
	// TODO: this is brute force it can be done by doing lcm
	c := 0 
	for zcount(s) != len(s) {
		for x, ck := range s {
			if ck[len(ck)-1] == 'Z'{ 

			}
			if i[c%len(i)] == 'L' {
				s[x] = m[ck][0]
			} else {
				s[x] = m[ck][1]
			}
		}
		c++
	}
	return c
}

func zcount(s []string) int {
	c := 0
	for _, v := range s {
		if v[len(v)-1] == 'Z' {
			c += 1
		}
	}
	return c
}

func remove(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func main() {
	f, err := os.ReadFile("day08.in")
	if err != nil {
		fmt.Print(err)
	}
	fs := string(f)
	ps := strings.Split(fs, "\n\n")
	i := ps[0]

	m := map[string][]string{}
	snodes := []string{}
	for _, x := range strings.Split(ps[1], "\n") {
		if len(x) == 0 {
			continue
		}
		p := strings.Split(x, "=")
		k := strings.Fields(p[0])[0]
		v := strings.Fields(p[1])
		m[k] = v
		if k[len(k)-1] == 'A' {
			snodes = append(snodes, k) 
		}
	}



	fmt.Printf("Part 1: %v\n", calculate([]string{"AAA"}, i, m))
	fmt.Printf("Part 2: %v\n", calculate(snodes, i, m))
}
