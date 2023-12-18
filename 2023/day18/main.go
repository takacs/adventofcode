package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"image"
)

type Instruction struct {
	dir string
	val int
	hex string
}

type Hexstruction struct {
	dir string
	val int
}

func convertInstruction(i Instruction) Hexstruction {
	ndir, _ := strconv.ParseInt(i.hex[2:7], 16, 0)
	dirmap := map[string]string{"0":"R", "1":"D", "2":"L", "3":"U"}
	dir := dirmap[string(i.hex[len(i.hex)-2])]
	hex := Hexstruction{dir: dir, val: int(ndir)}
	return hex
}

func in(element image.Point, collection []image.Point,) bool {
	for _, value := range collection {
		if element == value {
			return true
		}
	}
	return false
}

func createMatrix(points []image.Point) [][]rune {
	var maxX, maxY int
	for _, p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	matrix := make([][]rune, maxY+1)
	for i := range matrix {
		matrix[i] = make([]rune, maxX+1)
		for j := range matrix[i] {
			matrix[i][j] = '.'
		}
	}

	for _, p := range points {
		matrix[p.Y][p.X] = '#'
	}

	for i:=0; i<len(matrix); i++ {
		fmt.Println(string(matrix[i]))
	}

	return matrix
}


func floodfill(cs []image.Point, h image.Point) int {
	queue := []image.Point{h}
	v := image.Point{}
	visited := []image.Point{h}
	for len(queue) > 0 {
		v = queue[0]
		queue = queue[1:]
		if !in(image.Point{X:v.X+1, Y:v.Y}, cs) && !in(image.Point{X:v.X+1, Y:v.Y}, queue) && !in(image.Point{X:v.X+1, Y:v.Y}, visited) {
			cs = append(cs, image.Point{X:v.X+1, Y:v.Y})
			queue = append(queue, image.Point{X:v.X+1, Y:v.Y})
		}
		if !in(image.Point{X:v.X-1, Y:v.Y}, cs) && !in(image.Point{X:v.X-1, Y:v.Y}, queue) && !in(image.Point{X:v.X-1, Y:v.Y}, visited){
			cs = append(cs, image.Point{X:v.X-1, Y:v.Y})
			queue = append(queue, image.Point{X:v.X-1, Y:v.Y})
		}
		if !in(image.Point{X:v.X, Y:v.Y+1}, cs) && !in(image.Point{X:v.X, Y:v.Y+1}, queue) && !in(image.Point{X:v.X, Y:v.Y+1}, visited) {
			cs = append(cs, image.Point{X:v.X, Y:v.Y+1})
			queue = append(queue, image.Point{X:v.X, Y:v.Y+1})
		}
		if !in(image.Point{X:v.X, Y:v.Y-1}, cs) && !in(image.Point{X:v.X, Y:v.Y-1}, queue) && !in(image.Point{X:v.X, Y:v.Y-1}, visited) {
			cs = append(cs, image.Point{X:v.X, Y:v.Y-1})
			queue = append(queue, image.Point{X:v.X, Y:v.Y-1})
		}
		visited = append(visited, v)
	}

	return len(cs)
}

func main() {
	f, _ := os.ReadFile("day18.in")
	fs := string(f)

	ins := []Instruction{}
	for _, s := range strings.Split(fs, "\n") {
		if len(s) < 1 {
			continue
		}
		p := strings.Fields(s)
		v, _ := strconv.Atoi(p[1])
		ins = append(ins, Instruction{dir:p[0], val: v, hex: p[2]})
	}
	pos := image.Point{X:100, Y:100}
	coords := []image.Point{}
	coords = append(coords, pos)
	for _, i := range ins {
		if i.dir == "R" {
			for k := 0; k < i.val; k++ {
				pos.X += 1
				coords = append(coords, pos)
			}
		} else if i.dir == "L" {
			for k := 0; k < i.val; k++ {
				pos.X -= 1 
				coords = append(coords, pos)
			}
		} else if i.dir == "D" {
			for k := 0; k < i.val; k++ {
				pos.Y += 1
				coords = append(coords, pos)
			}
		} else if i.dir == "U" {
			for k := 0; k < i.val; k++ {
				pos.Y -= 1
				coords = append(coords, pos)
			}
		}	
	}

	fmt.Printf("Part 1: %v\n", floodfill(coords, image.Point{X:225, Y:125}))
	
	hinstr := []Hexstruction{}
	for _, instr := range ins {
		hinstr = append(hinstr, convertInstruction(instr))
	}

	pos = image.Point{X:1000, Y:1000}
	coords = []image.Point{}
	coords = append(coords, pos)
	b := 0
	for _, h := range hinstr {
		if h.dir == "R" {
			pos.X += h.val
		} else if h.dir == "L" {
			pos.X -= h.val 
		} else if h.dir == "D" {
			pos.Y += h.val
		} else if h.dir == "U" {
			pos.Y -= h.val
		}	
		coords = append(coords, pos)
		b += h.val
	}

	A := 0
	j := 0
	for i := 0; i < len(coords); i++ {
		j = i-1
		if i == 0 {
			j = len(coords)-1
		}
		A += coords[i].X * (coords[(i+1)%len(coords)].Y - coords[j].Y)
	}
	A /= 2
	fmt.Printf("Part 2: %v\n", A + (b/2) + 1)
}
