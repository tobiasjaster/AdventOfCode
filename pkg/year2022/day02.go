package year2022

import (
	"strings"
	//"sort"
)

type Day02 struct{}

type Points int64
type Rating int64

const (
	Rock    Points = 1
	Paper   Points = 2
	Scissor Points = 3
)

const (
	Loose Rating = 0
	Draw  Rating = 3
	Win   Rating = 6
)

func calcRating(opposite Points, me Points) Rating {
	var pointsRatingMap = map[Points]map[Points]Rating{
		Rock: {
			Rock:    Draw,
			Paper:   Win,
			Scissor: Loose,
		},
		Paper: {
			Rock:    Loose,
			Paper:   Draw,
			Scissor: Win,
		},
		Scissor: {
			Rock:    Win,
			Paper:   Loose,
			Scissor: Draw,
		},
	}
	return pointsRatingMap[opposite][me]
}

func calcPoints(opposite Points, rating Rating) Points {
	var pointsRatingMap = map[Points]map[Rating]Points{
		Rock: {
			Draw:  Rock,
			Win:   Paper,
			Loose: Scissor,
		},
		Paper: {
			Loose: Rock,
			Draw:  Paper,
			Win:   Scissor,
		},
		Scissor: {
			Win:   Rock,
			Loose: Paper,
			Draw:  Scissor,
		},
	}
	return pointsRatingMap[opposite][rating]
}

type Game struct {
	points int
}

func (g *Game) playRound(opposite string, my string) {
	var inputMap = map[string]Points{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
		"X": Rock,
		"Y": Paper,
		"Z": Scissor,
	}
	var rating Rating = calcRating(inputMap[opposite], inputMap[my])
	g.points += int(inputMap[my]) + int(rating)
}

func (g *Game) playRound2(opposite string, rating string) {
	var inputMapPoints = map[string]Points{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
	}
	var inputMapRating = map[string]Rating{
		"X": Loose,
		"Y": Draw,
		"Z": Win,
	}
	var my Points = calcPoints(inputMapPoints[opposite], inputMapRating[rating])
	g.points += int(my) + int(inputMapRating[rating])
}

func (p Day02) PartA(lines []string) any {
	var game Game
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		s := strings.Split(line, " ")
		game.playRound(s[0], s[1])
	}
	return game.points
}

func (p Day02) PartB(lines []string) any {
	var game Game
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		s := strings.Split(line, " ")
		game.playRound2(s[0], s[1])
	}
	return game.points
}
