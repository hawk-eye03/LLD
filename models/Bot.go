package models

type Bot struct {
	player             Player
	botDifficultyLevel BotDifficultyLevel
	botPlayingStrategy BotPlayingStrategy
}

func NewBot(botDifficultyLevel BotDifficultyLevel, player *Player) *Bot {
	BotPlayingStrategyFactory := BotPlayingStrategyFactory{}
	return &Bot{
		player:             *player,
		botDifficultyLevel: botDifficultyLevel,
		botPlayingStrategy: BotPlayingStrategyFactory.GetBotPlayingStrategyForDifficultyLevel(botDifficultyLevel),
	}
}

func (b *Bot) GetBotPlayerInfo() Player {
	return b.player
}

func (b *Bot) GetBotDifficultyLevel() BotDifficultyLevel {
	return b.botDifficultyLevel
}

func (b *Bot) SetBotDifficultyLevel(bdl BotDifficultyLevel) {
	b.botDifficultyLevel = bdl
}

func (b *Bot) MakeMove(game *Game) Cell {
	return b.botPlayingStrategy.MakeMove(game)
}
