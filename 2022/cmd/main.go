package main

import (
    "fmt"
    "github.com/katelynsalvatori/advent2022/pkg"
)

func main() {
    // cc, err := pkg.ParseCollectionFromFile("inputs/1.txt")
    // fmt.Println(cc.GetFormattedMaxCalories())

    //games, err := pkg.ParseGamesFromFile("inputs/2.txt")
    //fmt.Println(games.FormattedTotalScore())

    rc, err := pkg.ParseFileToRucksacks("inputs/3.txt")
    if err != nil {
        fmt.Errorf(err.Error())
    }
    fmt.Println(rc.GetFormattedBadgePriorities())
}
