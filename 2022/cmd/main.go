package main

import (
    "fmt"
    "github.com/katelynsalvatori/advent2022/pkg"
)

func main() {
    fileLines, err := pkg.ReadFileLinesForDay(5)

    if err != nil {
        fmt.Errorf(err.Error())
    }
    cm := pkg.ParseLinesToCrateMovement(fileLines)
    newCm := cm.ExecuteInstructions()
    fmt.Println(newCm.GetTopCrates())
}
