package mockTest

import (
	"github.com/Nolions/mockTest/game"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGoVersion(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockGame := game.NewMockGame(mockCtl)
	mockGame.EXPECT().GetType().Return("baseball")

	v := GetGameType(mockGame)

	assert.Equal(t, "baseball", v)
}

func TestGetNumberByBaseball(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockGame := game.NewMockGame(mockCtl)
	mockGame.EXPECT().GetType().Return("baseball")

	v := GetNumber(mockGame)

	assert.Equal(t, 9, v)
}

func TestGetNumberByFootball(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockGame := game.NewMockGame(mockCtl)
	mockGame.EXPECT().GetType().Return("football")

	v := GetNumber(mockGame)

	assert.Equal(t, 11, v)
}

func TestGetNumberNoMatch(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockGame := game.NewMockGame(mockCtl)
	mockGame.EXPECT().GetType().Return("")

	v := GetNumber(mockGame)

	assert.Equal(t, 0, v)
}
