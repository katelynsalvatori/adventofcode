package main

import (
	"fmt"
	"github.com/katelynsalvatori/advent2022/pkg"
)

func main() {
	games, err := pkg.ParseGamesFromFile("inputs/2.txt")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(games.FormattedTotalScore())
}
