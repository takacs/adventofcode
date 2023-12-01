package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

type Line struct {
    text string
}

func (l Line) getNumber(possible_values []string, first bool) int {
    var curr_index = l.getBaseIndex(first)
    var value int
    for i, number := range possible_values {
        index :=  l.matchIndex(first, l.text, number)
        if index == -1 {
            continue
        }
        if l.substitute(first, index, curr_index) {
            curr_index = index
            value = i % 10
        }
    }
    return value
}

func (l Line) substitute(ascending bool, index, curr_index int) bool {
    if ascending {
        return index <= curr_index 
    }
    return index >= curr_index 
}

func (l Line) getBaseIndex(ascending bool) int {
    if ascending {
        return len(l.text) + 1
    }
    return -1
}

func (l Line) matchIndex(first bool, base_string, substring string) int {
    if first {
        return strings.Index(base_string, substring)
    }
    return strings.LastIndex(base_string, substring)
}

func main() {
    file, err := os.Open("day01.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    lines := []Line{}
    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, Line{ text:line })
    }

    var tot int
    var tot2 int
    p1_values := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
    p2_values := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    for _, line := range(lines) {
        tot += line.getNumber(p1_values, true) * 10 + line.getNumber(p1_values, false)
        tot2 += line.getNumber(p2_values, true) * 10 + line.getNumber(p2_values, false)
    }

    fmt.Printf("Part 1: %d\n", tot)
    fmt.Printf("Part 2: %d\n", tot2)

}
