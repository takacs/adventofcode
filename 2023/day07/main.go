package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "sort"
)

type Hand struct {
	cards string
	bid int
	hand string
	orig string
}

func (h Hand) score() int {
	x := h.hand
	r := 0
	if x == "5" {
		r = 0
	} else if x == "41" {
		r = 1
	} else if x == "32" {
		r = 2
	} else if x == "311" {
		r = 3
	} else if x == "221" {
		r = 4
	} else if x == "2111" {
		r = 5
	} else if x == "11111" {
		r = 6
	}
	return r
}

func counter(s string) string {
    counter := make(map[rune]int)

    for _, r := range s {
        counter[r]++
    }
	
	keys := ""
	saveval := -1
	for k, v := range counter {
		if k == '1' {
			saveval = v
			continue
		}
		vs := strconv.Itoa(v)
		keys += vs
	}
	ss := sortstring(keys)
	ns := ss
	if saveval == 5 {
		return "5"
	} else if saveval > 0 {
		v, _ := strconv.Atoi(string(ss[0]))
		ns = strconv.Itoa(v+saveval) + ss[1:]
	} 
	return ns 
}

func sortstring(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
	last := len(s) - 1
    for i := 0; i < len(s)/2; i++ {
        s[i], s[last-i] = s[last-i], s[i]
    }
    return strings.Join(s, "")
}

func main() {

	parts := []bool{false,true}
    for _, p2 := range parts {
		file, err := os.Open("day07.in")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		scanner := bufio.NewScanner(file)
		hs := []Hand{}
		for scanner.Scan() {
			line := scanner.Text()
			p := strings.Fields(line)
			p[0] = strings.Replace(p[0], "T", ":", -1)
			if p2 {
				p[0] = strings.Replace(p[0], "J", "1", -1)
			} else {
				p[0] = strings.Replace(p[0], "J", ";", -1)
			}
			p[0] = strings.Replace(p[0], "Q", "<", -1)
			p[0] = strings.Replace(p[0], "K", "=", -1)
			p[0] = strings.Replace(p[0], "A", ">", -1)
			op := p[0]
			b, _ := strconv.Atoi(p[1])
			hs = append(hs, Hand{cards: sortstring(p[0]), bid: b, hand:counter(p[0]), orig: op})
		}
		
		sort.Slice(hs, func(i, j int) bool {
			is := hs[i].score()
			js := hs[j].score()
			if is != js {
				return is > js
			}
			ih := hs[i].orig
			jh := hs[j].orig
			return ih < jh
		})

		tot := 0
		for i, h := range hs {
			if p2{
				fmt.Println(i, h, h.score())
			}
			tot += (i+1)*h.bid
		}
		
		if p2 {
			fmt.Printf("Part 2: %v\n", tot)
		} else {
			fmt.Printf("Part 1: %v\n", tot)
		}
		file.Close()
    }
}


