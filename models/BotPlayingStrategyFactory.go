package models

type BotPlayingStrategyFactory struct{}

func (botStrategyFactory *BotPlayingStrategyFactory) GetBotPlayingStrategyForDifficultyLevel(botPlayingDifficultyLevel BotDifficultyLevel) BotPlayingStrategy {
	return &EasyBotPlayingStrategy{}

	// we can implement multiple strategies according to difficulty level of bot, for simplicity only simple strategy has been choosen
	//        return switch (difficultyLevel) {
	//        case EASY -> new EasyBotPlayingStrategy();
	//        case MEDIUM -> new MediumBotPlayingStrategy();
	//        case HARD -> new HardBotPlayingStrategy();
}
