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
	if x == "5" {
		return 0
	} else if x == "41" {
		return 1
	} else if x == "32" {
		return 2
	} else if x == "311" {
		return 3
	} else if x == "221" {
		return 4
	} else if x == "2111" {
		return 5
	} else if x == "1111" {
		return 6
	}
	return 7
}

func counter(s string) string {
    counter := make(map[rune]int)

    for _, r := range s {
        counter[r]++
    }
	
	keys := ""
	for _, v := range counter {
		vs := strconv.Itoa(v)
		keys += vs
	}
    return sortstring(keys)
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
    file, err := os.Open("day07.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    hs := []Hand{}
    scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		p := strings.Fields(line)
		p[0] = strings.Replace(p[0], "T", ":", -1)
		p[0] = strings.Replace(p[0], "J", ";", -1)
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
		fmt.Println(h)
		tot += (i+1)*h.bid
	}

	fmt.Printf("Part 1: %v\n", tot)
}
