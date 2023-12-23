package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QE struct {
	pos                                    string
	lox, hix, lom, him, loa, hia, los, his int
}

func curse(q QE, w map[string]string) int {
	if q.pos == "R" {
		return 0
	} else if q.pos == "A" {
		return (q.hix - q.lox + 1) * (q.him - q.lom + 1) * (q.hia - q.loa + 1) * (q.his - q.los + 1)
	}

	total := 0
	rules := strings.Split(w[q.pos], ",")

	ruleKey := "1"
	nextKey := "1"
	T := []int{}
	F := []int{}
	for _, rule := range rules {
		if strings.Contains(rule, "<") {
			ruleParts := strings.Split(rule, "<")
			rulePartsTwo := strings.Split(ruleParts[1], ":")
			ruleValue, _ := strconv.Atoi(rulePartsTwo[0])
			nextKey = rulePartsTwo[1]
			ruleKey = ruleParts[0]
			lo, hi := getkey(q, ruleKey)
			T = []int{lo, min(hi, ruleValue-1)}
			F = []int{max(ruleValue, lo), hi}
		} else if strings.Contains(rule, ">") {
			ruleParts := strings.Split(rule, ">")
			rulePartsTwo := strings.Split(ruleParts[1], ":")
			ruleValue, _ := strconv.Atoi(rulePartsTwo[0])
			nextKey = rulePartsTwo[1]
			ruleKey = ruleParts[0]
			lo, hi := getkey(q, ruleKey)
			T = []int{max(ruleValue+1, lo), hi}
			F = []int{lo, min(hi, ruleValue)}
		} else {
			q.pos = rule
			total += curse(q, w)
			break
		}
		if T[0] <= T[1] {
			qCopy := q
			qCopy.pos = nextKey
			qCopy = setkey(qCopy, T, ruleKey)
			total += curse(qCopy, w)
		}
		if F[0] <= F[1] {
			q = setkey(q, F, ruleKey)
		}

	}
	return total
}

func min(v1, v2 int) int {
	v := []int{v1, v2}
	minv := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < minv {
			minv = v[i]
		}
	}
	return minv
}
func max(v1, v2 int) int {
	v := []int{v1, v2}
	maxv := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > maxv {
			maxv = v[i]
		}
	}
	return maxv
}

func getkey(q QE, k string) (int, int) {
	if k == "x" {
		return q.lox, q.hix
	}
	if k == "m" {
		return q.lom, q.him
	}
	if k == "a" {
		return q.loa, q.hia
	}
	if k == "s" {
		return q.los, q.his
	}
	return 1, 1
}

func setkey(q QE, T []int, k string) QE {
	if k == "x" {
		q.lox = T[0]
		q.hix = T[1]
	}
	if k == "m" {
		q.lom = T[0]
		q.him = T[1]
	}
	if k == "a" {
		q.loa = T[0]
		q.hia = T[1]
	}
	if k == "s" {
		q.los = T[0]
		q.his = T[1]
	}
	return q
}

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
			for i := 0; i < len(ins); i++ {
				// if last instruction break
				if !strings.Contains(ins[i], ":") {
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
	fmt.Printf("Part 1: %v\n", tot)

	Q := QE{pos: "in", lox: 1, hix: 4000, loa: 1, hia: 4000, lom: 1, him: 4000, los: 1, his: 4000}
	fmt.Printf("Part 2: %v\n", curse(Q, w))
}
