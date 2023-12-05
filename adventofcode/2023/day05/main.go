package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func between(value, bound1, bound2 int) bool {
	mn := bound1
	mx := bound2
	if bound1 > bound2 {
		mn = bound2
		mx = bound1
	}
	return (value >= mn) && (value <= mx)
}

func convert(s int, mm string) int {
	if mm == "" {
		return s
	}
	cm := strings.Split(mm, "\n")
	cm = cm[1:]
	for _, c := range cm {
		if len(c) == 0 {
			return s
		}
		cf := strings.Fields(c)
		de, _ := strconv.Atoi(cf[0])
		so, _ := strconv.Atoi(cf[1])
		of, _ := strconv.Atoi(cf[2])
		dif := so - de
		if between(s, so, so+of) {
			s = s - dif
			break
		}
	}
	return s
}

func min(nums []int) int {
	mn := nums[0]
	for _, num := range nums {
		if num < mn {
			mn = num
		}
	}
	return mn
}

func main() {
	f, err := os.ReadFile("day05.in")
	if err != nil {
		fmt.Print(err)
	}
	fs := string(f)
	sed := []int{}
	ls := strings.Split(fs, "\n\n")
	p := strings.Split(ls[0], ":")
	ss := strings.Fields(p[1])
	for _, se := range ss {
		v, _ := strconv.Atoi(se)
		sed = append(sed, v)
	}

	locs := []int{}
	for _, s := range sed {
		for _, m := range ls[1:] {
			s = convert(s, m)
		}
		locs = append(locs, s)
	}

	fmt.Printf("Part 1: %v\n", min(locs))

}
