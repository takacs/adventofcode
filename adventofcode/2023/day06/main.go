package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	f, err := os.ReadFile("day06.in")
	if err != nil {
		fmt.Print(err)
	}
	fs := string(f)
	ps := strings.Split(fs, "\n")
	times := strings.Fields(strings.Split(ps[0], ":")[1])
	distances := strings.Fields(strings.Split(ps[1], ":")[1])
	
	t := []int{}
	d := []int{}
	for i := 0; i < len(times); i++ {
		tval, _ := strconv.Atoi(times[i])
		dval, _ := strconv.Atoi(distances[i])
		t = append(t, tval)
		d = append(d, dval)
	}
	
	res := 1
	for i, time := range t {
		distance := d[i]
		tot := 0
		for i := 0; i < time; i++ {
			if i * (time-i) > distance {
				tot++
			}
		}
		res *= tot
	}

	fmt.Printf("Part 1: %v\n", res)

	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(ps[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(ps[1], ":")[1], " ", ""))
	tot := 0
	for i := 0; i < time; i++ {
		if i * (time-i) > distance {
			tot++
		}
	}

	fmt.Printf("Part 2: %v\n", tot)
}
