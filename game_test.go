package mockTest

import (
	"github.com/Nolions/mockTest/game"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockController *gomock.Controller

func TestGetGoVersion(t *testing.T) {
	mockGame := newMockGame(t)
	mockController.Finish()
	
	mockGame.EXPECT().GetType().Return("baseball")

	v := GetGameType(mockGame)

	assert.Equal(t, "baseball", v)
}

func TestGetNumberByBaseball(t *testing.T) {
	mockGame := newMockGame(t)
	mockController.Finish()

	mockGame.EXPECT().GetType().Return("baseball")

	v := GetNumber(mockGame)

	assert.Equal(t, 9, v)
}

func TestGetNumberByFootball(t *testing.T) {
	mockGame := newMockGame(t)
	mockController.Finish()

	mockGame.EXPECT().GetType().Return("football")

	v := GetNumber(mockGame)

	assert.Equal(t, 11, v)
}

func TestGetNumberNoMatch(t *testing.T) {
	mockGame := newMockGame(t)
	mockController.Finish()

	mockGame.EXPECT().GetType().Return("")

	v := GetNumber(mockGame)

	assert.Equal(t, 0, v)
}

func TestStartGame(t *testing.T) {
	mockGame := newMockGame(t)
	mockController.Finish()

	mockGame.EXPECT().Start().Times(1)

	StartGame(mockGame)
}

func TestEndGame(t *testing.T) {
	mockGame := newMockGame(t)
	mockController.Finish()

	mockGame.EXPECT().End().Times(1)

	EndGame(mockGame)
}

func newMockGame(t *testing.T) *game.MockGame {
	mockController = gomock.NewController(t)
	return game.NewMockGame(mockController)
}
