package models

import (
	"fmt"

	"github.com/hawk-eye03/LLD/TicTacToe/strategies/botPlayingStrategies"
)

type Bot struct {
	player             Player
	botDifficultyLevel BotDifficultyLevel
}

func NewBot(difficultyLevel BotDifficultyLevel, player *Player) *Bot {
	return &Bot{
		player:             *player,
		botDifficultyLevel: difficultyLevel,
	}
}

func (b *Bot) GetInfo() {
	fmt.Println("I am a bot")
}

func (b *Bot) GetBotPlayerInfo() Player {
	return b.player
}

func (b *Bot) GetBotDifficultyLevel() botPlayingStrategies.BotPlayingStrategy {
	return b.botDifficultyLevel
}

func (b *Bot) SetBotDifficultyLevel(bdl BotDifficultyLevel) {
	b.botDifficultyLevel = bdl
}
