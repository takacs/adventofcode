package main

import (
	"fmt"
	"os"
	"bufio"
)

type Beam struct {
	pos Position
	dir Direction
}

func (b *Beam) move() {
	b.pos.x += b.dir.x
	b.pos.y += b.dir.y
}

func (b Beam) valid(m [][]rune) bool {
	X := len(m)
	Y := len(m[0])
	return b.pos.x >= 0 && b.pos.x < X && b.pos.y >= 0 && b.pos.y < Y
}


type Position struct {
	x, y int
}

type Direction struct {
	x, y int
}

func max(x []int) int {
   ln := x[0]
   for i := 1; i < len(x); i++ {
      if x[i] > ln {
         ln= x[i]
      }
   }
   return ln
}



func main() {
    file, err := os.Open("day16.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    m := [][]rune{}
    for scanner.Scan() {
        l := scanner.Text()
        row := []rune(l)
		m = append(m, row)
    }
    spos := []Beam{}
	for i := 0; i < len(m); i++ {
		spos = append(spos, Beam{pos: Position{x:i, y:-1}, dir: Direction{x:0,y:1}})
		spos = append(spos, Beam{pos: Position{x:i, y:len(m[0])}, dir: Direction{x:0,y:-1}})
	}
	for i := 0; i < len(m[0]); i++ {
		spos = append(spos, Beam{pos: Position{x:-1, y:i}, dir: Direction{x:1,y:0}})
		spos = append(spos, Beam{pos: Position{x:len(m), y:i}, dir: Direction{x:-1,y:0}})
	}

	energies := []int{}
	for _, sp := range spos {
		queue := []Beam{}
		queue = append(queue, sp)
		s := map[Beam]bool{}
		sol := map[Position]bool{}
		for len(queue) > 0 {
			beam := queue[0] 
			queue = queue[1:]

			beam.move()
			if !beam.valid(m) {
				continue
			}
			_, exists := s[beam]
			if exists {
				continue
			}
			s[beam] = true
			sol[beam.pos] = true
			
			bx := beam.pos.x
			by := beam.pos.y
			bdx := beam.dir.x
			bdy := beam.dir.y
			if m[bx][by] == '|' && bdx == 0 {
				queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:1, y:0}})
				queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:-1, y:0}})
			} else if m[bx][by] == '-' && bdy == 0 {
				queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:0, y:1}})
				queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:0, y:-1}})
			} else if m[bx][by] == '\\' {
				if bdx == 0 && bdy == 1 {
					queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:1, y:0}})
				} else if bdx == 0 && bdy == -1 {
					queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:-1, y:0}})
				} else if bdx == 1 && bdy == 0 {
					queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:0, y:1}})
				} else if bdx == -1 && bdy == 0 {
					queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:0, y:-1}})
				}
			} else if m[bx][by] == '/' {
				if bdx == 0 && bdy == 1 {
					queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:-1, y:0}})
				} else if bdx == 0 && bdy == -1 {
					queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:1, y:0}})
				} else if bdx == 1 && bdy == 0 {
					queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:0, y:-1}})
				} else if bdx == -1 && bdy == 0 {
					queue = append(queue, Beam{pos: Position{x:bx, y:by}, dir: Direction{x:0, y:1}})
				}
			} else {
				queue = append(queue, beam)
			}
			
		}
		energies = append(energies, len(sol))
	}
	fmt.Println(energies[0])
	fmt.Println(max(energies))

}
	
