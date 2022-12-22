package main

import (
	"fmt"
	"github.com/katelynsalvatori/advent2022/pkg"
)

func main() {
	// cc, err := pkg.ParseCollectionFromFile("inputs/1.txt")
	// if err != nil {
	// 	fmt.Errorf(err.Error())
	// }
	// fmt.Println(cc.GetFormattedMaxCalories())

	games, err := pkg.ParseGamesFromFile("inputs/2.txt")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(games.FormattedTotalScore())
}
