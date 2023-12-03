package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Schematic struct {
    engineMap [][]Part
}

type Part struct {
    value rune
    part bool
    adjacencyCount int
}

type Direction struct {
    x, y int
}

func getDirections() []Direction{
    var directions []Direction
    directions = append(directions, Direction{x: -1, y:0})
    directions = append(directions, Direction{x: -1, y:1})
    directions = append(directions, Direction{x: 0, y:1})
    directions = append(directions, Direction{x: 1, y:1})
    directions = append(directions, Direction{x: 1, y:0})
    directions = append(directions, Direction{x: 1, y:-1})
    directions = append(directions, Direction{x: 0, y:-1})
    directions = append(directions, Direction{x: -1, y:-1})
    
    return directions
}

func (s Schematic) loadAdjacencies() {
    for i := 0; i < len(s.engineMap); i++ {
        for j := 0; j < len(s.engineMap[i]); j++ {
            s.engineMap[i][j].part = s.isPart(i, j)
        }
    }
}

func (s *Schematic) isPart(i, j int) bool {
    if !unicode.IsDigit(s.engineMap[i][j].value) {
        return false
    }
    var hasAdjacency bool
    for _, direction := range getDirections() {

        if !isIndexValid(s.engineMap, i+direction.x, j+direction.y) {
            continue
        }
        value := s.engineMap[i+direction.x][j+direction.y].value
        fmt.Println(string(value))
        if s.isSymbol(value) {
            s.engineMap[i+direction.x][j+direction.y].adjacencyCount += 1
            hasAdjacency = true
        }
    }
    return hasAdjacency
}

func (s *Schematic) isSymbol(value rune) bool {
    symbols := []rune{'=','$', '/', '*', '+', '@', '&', '-', '%', '#'}
    for _, symbol := range symbols {
        if value == symbol {
            return true
        }
    }
    return false
}

func (s Schematic) getPartSum() int {
    var curr, tot int
    multiplier := 10
    newNumber := true
    newNumberPart := false
    for i := 0; i < len(s.engineMap); i++ {
        for j := len(s.engineMap[i])-1; j >= 0; j-- {
            fmt.Println(curr, newNumber, s.engineMap[i][j])
            if unicode.IsDigit(s.engineMap[i][j].value) {
                if newNumber {
                    curr = int(s.engineMap[i][j].value - '0')
                    newNumberPart = s.engineMap[i][j].part || newNumberPart
                    newNumber = false
                } else {
                    curr += int(s.engineMap[i][j].value - '0') * multiplier
                    newNumberPart = s.engineMap[i][j].part || newNumberPart
                    multiplier *= 10
                }
            } else {
                if newNumberPart {
                    fmt.Printf("adding: %d\n", curr)
                    tot += curr
                }
                curr = 0
                newNumber = true
                multiplier = 10
                newNumberPart = false
            }

        }
        if newNumberPart {
            tot += curr
            fmt.Printf("adding: %d\n", curr)
        }
        curr = 0
        newNumber = true
        multiplier = 10
        newNumberPart = false
    }
    return tot
}

func isIndexValid(matrix [][]Part, row, col int) bool {
    numRows := len(matrix)
    if numRows == 0 {
        return false
    }

    numCols := len(matrix[0])

    return row >= 0 && row < numRows && col >= 0 && col < numCols
}

func (s Schematic) printSchematicPartMap() {
    for i := 0; i < len(s.engineMap); i++ {
        for j := 0; j < len(s.engineMap[i]); j++ {
            fmt.Printf(string(s.engineMap[i][j].value))
        }
        fmt.Printf("\n")
    }
    
}

func (s Schematic) printSchematicPartMapAdjacent() {
    for i := 0; i < len(s.engineMap); i++ {
        for j := 0; j < len(s.engineMap[i]); j++ {
            if s.engineMap[i][j].part {
                fmt.Printf("T")
            } else {
                fmt.Printf("F")
            }
        }
        fmt.Printf("\n")
    }
    
}


func main() {
    file, err := os.Open("day03.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    schematic := Schematic{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        var rowParts []Part
        row := []rune(line)
        for _, r := range row {
            part := Part{value:r}
            rowParts = append(rowParts, part)
            
        }
        schematic.engineMap = append(schematic.engineMap, rowParts)
    }
    schematic.loadAdjacencies()
    fmt.Println(schematic.getPartSum())
//    fmt.Printf("Part 1: %d\n", validGames)

}
