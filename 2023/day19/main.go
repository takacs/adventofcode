package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)


func main() {
    f, err := os.ReadFile("day19.in")
    if err != nil {
	    fmt.Print(err)
    }
    fs := strings.Split(string(f), "\n\n")
	wfs := fs[0]
	rtgs := fs[1]

	w := map[string]string{}
	for _, s := range strings.Split(wfs, "\n") {
		p := strings.Split(s, "{")
		w[p[0]] = strings.ReplaceAll(p[1], "}", "")
	}

	pmap := []map[string]int{}
	for _, y := range strings.Split(rtgs, "\n") {
		if len(y) < 1 {
			continue
		}
		y = strings.ReplaceAll(y, "{", "")
		y = strings.ReplaceAll(y, "}", "")
		m := map[string]int{}
		for _, part := range strings.Split(y, ",") {
			x := strings.Split(part, "=")
			v, _ := strconv.Atoi(x[1])
			m[x[0]] = v

		}
		pmap = append(pmap, m)
	}

	A := 0
	tot := 0
	for _, p := range pmap {
		ckey := "in"
		for true {
			currentw := w[ckey]
			ins := strings.Split(currentw, ",")
			for i:=0; i < len(ins); i++ {
				// if last instruction break
				if !strings.Contains(ins[i], ":"){
					ckey = ins[i]
					break
				}
				dir := strings.Split(ins[i], ":")
				if strings.Contains(dir[0], "<") {
					pi := strings.Split(dir[0], "<")
					v, _ := strconv.Atoi(pi[1])
					if p[pi[0]] < v {
						ckey = dir[1]
						break
					}
				} else if strings.Contains(dir[0], ">") {
					pi := strings.Split(dir[0], ">")
					v, _ := strconv.Atoi(pi[1])
					if p[pi[0]] > v {
						ckey = dir[1]
						break
					}
				}
			}

			if ckey == "R" {
				fmt.Println(p, "R")
				break
			} else if ckey == "A" {
				tot += p["x"] + p["m"] + p["a"] + p["s"] 
				A++
				break
			}
		}
	}
	fmt.Println(A, tot)
}
	
