package mockTest

import (
	"github.com/Nolions/mockTest/game"
)

func GetGameType(s game.Game) string {
	return s.GetType()
}

func GetNumber(s game.Game) int {
	n := 0
	switch s.GetType() {
	case "baseball":
		n = 9
	case "football":
		n = 11
	}

	return n
}
