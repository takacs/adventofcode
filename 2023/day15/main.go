package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Focal struct {
	name string
	val int
}

func hash(s string) int {
	ascii := 0
	for _, c := range s {
		if c == '\n' {
			continue
		}
		ascii += int(c)
		ascii *= 17
		ascii = ascii%256
	}
	return ascii
}

func inBox(box []Focal, f Focal) bool {
	for _, b := range box {
		if b.name == f.name {
			return true
		}
	}
	return false
}

func removeFromBox(box []Focal, f Focal) []Focal {
	for i:=0; i < len(box); i++ {
		if box[i].name == f.name {
			box = remove(box, i)	
		}
	}
	return box
}

func updateBoxVal(box []Focal, f Focal) []Focal {
	for i:=0; i < len(box); i++ {
		if box[i].name == f.name {
			box[i].val = f.val	
		}
	}
	return box
}

func remove(slice []Focal, s int) []Focal {
    return append(slice[:s], slice[s+1:]...)
}

func main() {
    f, err := os.ReadFile("day15.in")
    if err != nil {
	    fmt.Print(err)
    }
    fs := string(f)

	tot := 0
	for _, s := range strings.Split(fs, ",") {
		tot += hash(s)
	}
	fmt.Printf("Part 1: %v\n", tot)

	b := []Focal{}
	for _, s := range strings.Split(fs, ","){
		if strings.Contains(s, "\n") {
			s = s[:len(s)-1]
		}
		if strings.Contains(s, "-"){
			p := strings.Split(s, "-")
			b = append(b, Focal{name: p[0], val:-1})
		} else if strings.Contains(s, "=") {
			p := strings.Split(s, "=")
			v, _ := strconv.Atoi(p[1])
			b = append(b, Focal{name: p[0], val:v})
		}
	}

	boxes := [][]Focal{}
	
	for i:=0; i<256 ;i++ {
		boxes = append(boxes, make([]Focal, 0))
	}

	for _, f := range b {
		boxnum := hash(f.name)
		if inBox(boxes[boxnum], f) {
			if f.val > 0 {
				boxes[boxnum] = updateBoxVal(boxes[boxnum], f)
			} else {
				boxes[boxnum] = removeFromBox(boxes[boxnum], f)
			}
		} else if f.val > 0 {
			boxes[boxnum] = append(boxes[boxnum], f)
		}
	}

	totp2 := 0
	for i:=0; i<len(boxes); i++{
		for j, l := range boxes[i] {
			vv:= (i+1)*(j+1)*l.val
			totp2 += vv
		}
	}
	fmt.Printf("Part 2: %v\n", totp2)

}
	
