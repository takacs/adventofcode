package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func getNums (l string) []int {
    nums := []int{} 
    for _, n := range strings.Fields(l) {
	v, _ := strconv.Atoi(n)
	nums = append(nums, v)
    }
    return nums
}

func solve (l string, reverse bool) int {
    if len(l) == 0 {
	return 0
    }
    nums := getNums(l)
    if reverse {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
	    nums[i], nums[j] = nums[j], nums[i]
    }
    }
    s := f(nums)
    return s
}

func f (n []int) int {
    nn := []int{}
    for i := 0; i < len(n)-1; i++ {
	nn = append(nn, n[i+1]-n[i])
    }
    if all_zero(nn) {
	return n[len(n)-1]
    }
    return n[len(n)-1] + f(nn)
}


func all_zero(n []int) bool {
    for _, nu := range n {
	if nu != 0 {
	    return false
	}
    }
    return true
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    f, err := os.ReadFile("day09.in")
    if err != nil {
	    fmt.Print(err)
    }
    fs := string(f)	
    h := strings.Split(fs, "\n")
    tot := 0
    tot2 := 0
    for _, l := range h {
	tot += solve(l, false)
	tot2 += solve(l, true)
    }
    fmt.Printf("Part 1: %v\n", tot) 
    fmt.Printf("Part 2: %v\n", tot2) 
}
