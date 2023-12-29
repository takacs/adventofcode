package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type E struct {
	name string
	path []Edge
}

type Edge struct {
	s, e string
}

type EdgeSlice struct {
	e Edge
	v int
}

func nedge(s, e string) Edge {
	s1 := s
	s2 := e
	if s1 > s2 {
		s1, s2 = s2, s1
	}
	return Edge{s1, s2}
}
func findpath(s, t string, G map[string]map[string]int) (bool, []Edge) {
	Q := []E{}
	Q = append(Q, E{name: s, path: []Edge{}})
	S := map[Edge]bool{}
	for len(Q) > 0 {
		curr := Q[0]
		Q = Q[1:]
		for k := range G[curr.name] {
			edge := nedge(curr.name, k)
			if k == t {
				curr.path = append(curr.path, edge)
				return true, curr.path
			}
			exists, _ := S[edge]
			if exists {
				continue
			}
			S[edge] = true
			p := curr.path
			p = append(p, edge)
			Q = append(Q, E{name: k, path: p})
		}

	}

	return false, []Edge{}
}

func main() {
	f, err := os.ReadFile("day25.in")
	if err != nil {
		fmt.Print(err)
	}
	fs := string(f)

	G := map[string]map[string]int{}
	for _, l := range strings.Split(fs, "\n") {
		if len(l) < 2 {
			continue
		}
		s := strings.Split(l, ":")
		for _, g := range strings.Fields(s[1]) {
			_, exists := G[s[0]]
			if !exists {
				G[s[0]] = make(map[string]int)
			}
			_, exists = G[g]
			if !exists {
				G[g] = make(map[string]int)
			}

			G[s[0]][g] = 1
			G[g][s[0]] = 1
		}
	}

	n := []string{}
	for k := range G {
		n = append(n, k)
	}

	NC := map[Edge]int{}
	for i := 0; i < len(n); i++ {
		if i == 100 {
			break
		}
		for j := i + 1; j < len(n); j++ {
			fmt.Println(i,j)
			exists, path := findpath(n[i], n[j], G)
			if exists {
				for _, edge := range path {
					NC[edge]++
				}
			}
		}
	}

	es := []EdgeSlice{}
	for k, v := range NC {
		es = append(es, EdgeSlice{k, v})
	}

	sort.Slice(es, func(i, j int) bool {
		return es[i].v > es[j].v
	})
	x := es[:3]
	for _, v := range x {
		delete(G[v.e.s], v.e.e)
		delete(G[v.e.e], v.e.s)
	}

	Q := []string{n[0]}
	S := map[string]bool{}
	S[n[0]] = true
	for len(Q) > 0 {
		curr := Q[0]
		Q = Q[1:]
		for k := range G[curr] {
			_, exists := S[k]
			if exists {
				continue
			}
			S[k] = true
			Q = append(Q, k)
		}
	}
	fmt.Println(len(S), len(G)-len(S))
}
