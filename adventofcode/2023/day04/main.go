package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
)

func getScratchcardScore(line string) (int, int, error) {
    parts := strings.Split(line, ":")
    if len(parts) != 2 {
        return -1, -1,  fmt.Errorf("invalid input")
    }

    nums := strings.Split(parts[1], "|")
    gamenums := getNumbers(strings.Fields(nums[0]))
    mynums := getNumbers(strings.Fields(nums[1])) 
    
    games, score := getScore(gamenums, mynums)
    return games, score, nil
}

func getNumbers(s []string) []int {
    ps := []int{}
    for _, number := range s {
        v, _ := strconv.Atoi(number)
        ps = append(ps, v)
    }
    return ps
}

func getScore(gnum, mnum []int) (int, int) {
    tot := -1.0
    for _, num := range mnum {
        for _, gnum := range gnum {
            if num == gnum {
                tot += 1
            }
        }
    }
    if tot < 0 {
        return 0, 0
    }
    return int(tot)+1, int(math.Pow(2, tot))
}

func sum(n []int) int {
    tot := 0
    for _, x := range n {
        tot += x
    }
    return tot
}

func main() {
    file, err := os.Open("day04.in")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    scores := []int{}
    games := []int{}
    for scanner.Scan() {
        line := scanner.Text()
        game, score, _ := getScratchcardScore(line)
        scores = append(scores, score)
        games = append(games, game)
    }

    fmt.Printf("Part 1: %d\n", sum(scores))

    multiplier := make([]int, len(games))
    for i := range multiplier {
        multiplier[i] = 1
    }

    for i := 0; i < len(games); i++ {
        g := games[i]
        for j:= i+1; j < len(multiplier) && g != 0; j++ {
            multiplier[j] += 1*multiplier[i]
            g--
        }
    }
    fmt.Printf("Part 2: %d\n", sum(multiplier))
}
