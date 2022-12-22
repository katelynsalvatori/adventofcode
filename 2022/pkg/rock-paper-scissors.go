package pkg

import (
	"fmt"
	"os"
	"strings"
)

type Move int
type Result int

const (
	Rock     Move   = 1
	Paper    Move   = 2
	Scissors Move   = 3
	Win      Result = 6
	Loss     Result = 0
	Draw     Result = 3
)

type Game struct {
	GameResult   Result
	OpponentMove Move
}

type Games struct {
	GamesPlayed []*Game
	Score       int
}

func (g Games) FormattedTotalScore() string {
	return fmt.Sprintf("Your total score is %d", g.Score)
}

func ParseGamesFromFile(filename string) (*Games, error) {
	fileBytes, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return ParseGamesFromString(string(fileBytes)), nil
}

func ParseGamesFromString(gamesString string) *Games {
	gamesSplit := strings.Split(gamesString, "\n")
	var gamesPlayed []*Game
	totalScore := 0

	for _, g := range gamesSplit {
		game := ParseGameFromString(g)
		gamesPlayed = append(gamesPlayed, game)
		totalScore += game.YourScore()
	}

	return &Games{
		GamesPlayed: gamesPlayed,
		Score:       totalScore,
	}
}

func ParseGameFromString(gameString string) *Game {
	gameStrategy := strings.Split(gameString, " ")
	opponentMove := MoveFromStringRepresentation(gameStrategy[0])
	result := ResultFromStringRepresentation(gameStrategy[1])
	return &Game{
		OpponentMove: opponentMove,
		GameResult:   result,
	}
}

func MoveFromStringRepresentation(representation string) Move {
	if representation == "A" {
		return Rock
	}
	if representation == "B" {
		return Paper
	}
	// "C"
	return Scissors
}

func ResultFromStringRepresentation(representation string) Result {
	if representation == "X" {
		return Loss
	}
	if representation == "Y" {
		return Draw
	}
	// "Z"
	return Win
}

func (g Game) YourMove() Move {
	if g.GameResult == Win {
		if g.OpponentMove == Rock {
			return Paper
		}
		if g.OpponentMove == Paper {
			return Scissors
		}
		return Rock
	}

	if g.GameResult == Loss {
		if g.OpponentMove == Rock {
			return Scissors
		}
		if g.OpponentMove == Paper {
			return Rock
		}
		if g.OpponentMove == Scissors {
			return Paper
		}
	}
	return g.OpponentMove
}

func (g Game) YourScore() int {
	return int(g.YourMove()) + int(g.GameResult)
}
