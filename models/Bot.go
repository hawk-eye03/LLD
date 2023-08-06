package models

import "github.com/hawk-eye03/LLD/TicTacToe/strategies/botPlayingStrategies"

type Bot struct {
	player             Player
	botDifficultyLevel BotDifficultyLevel
}

func (b *Bot) getBotDifficultyLevel() botPlayingStrategies.BotPlayingStrategy {
	return b.botDifficultyLevel
}

func (b *Bot) setBotDifficultyLevel(bdl BotDifficultyLevel) {
	b.botDifficultyLevel = bdl
}
