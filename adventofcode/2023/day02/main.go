package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type Set struct {
    cubes map[string]int
}

func (s *Set) setColor(color string, value int) {
    s.cubes[color] = value
} 

type Game struct {
    id int
    sets []Set
}


func createGame(line string) (Game, error) {
    parts := strings.Split(line, ":")
    if len(parts) != 2 {
        return Game{}, fmt.Errorf("invalid input")
    }

    game := Game{}
    id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
    if err != nil {
        return Game{}, fmt.Errorf("no game id")
    }
    game.id = id

    var sets []Set
    for _, set := range strings.Split(parts[1], ";") {
        currset := Set{ cubes: make(map[string]int) }
        colors := strings.Split(set, ",")
        for _, color := range colors {
            values := strings.Split(color, " ")
            value, _ := strconv.Atoi(values[1])
            currset.setColor(values[2], value)
        }
        sets = append(sets, currset)
    }
    game.sets = sets

    return game, nil
}

func (g Game) maxColor(color string) int {
    var maxCube int
    for _, set := range g.sets {
        if maxCube < set.cubes[color] {
            maxCube = set.cubes[color]
        }
    }
    return maxCube
}



func main() {
    file, err := os.Open("day02.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    games := []Game{}
    for scanner.Scan() {
        line := scanner.Text()
        game, _ := createGame(line)
        games = append(games, game)
    }

    var validGames int
    for _, game := range games {
        if game.maxColor("red") <= 12 && game.maxColor("green") <= 13 && game.maxColor("blue") <= 14 {
            validGames += game.id
        }  
    }

    fmt.Printf("Part 1: %d\n", validGames)

    var powerCount int
    for _, game := range games {
        powerCount += game.maxColor("red") * game.maxColor("blue") * game.maxColor("green")

    }
    fmt.Printf("Part 2: %d\n", powerCount)

}
