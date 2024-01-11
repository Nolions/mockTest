package mockTest

import (
	"github.com/Nolions/mockTest/game"
)

// GetGameType 取得比賽類型
func GetGameType(s game.Game) string {
	return s.GetType()
}

// StartGame 比賽開始
func StartGame(s game.Game) {
	s.Start()
}

// EndGame 比賽結束
func EndGame(s game.Game) {
	s.End()
}

// GetNumber 取得比賽各隊上場人數
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
