// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"gpt-telegran-bot/internal/di/config"
	"gpt-telegran-bot/internal/domain/service"
	"gpt-telegran-bot/internal/domain/service/editor"
	"gpt-telegran-bot/internal/domain/service/generator"
	"gpt-telegran-bot/internal/domain/usecase"
	"gpt-telegran-bot/internal/infrastructure/client/openAi"
	"gpt-telegran-bot/internal/infrastructure/service/cache"
	openAi4 "gpt-telegran-bot/internal/infrastructure/service/editor/openAi"
	openAi3 "gpt-telegran-bot/internal/infrastructure/service/generator/openAi"
	"gpt-telegran-bot/internal/infrastructure/service/messenger"
	openAi2 "gpt-telegran-bot/internal/infrastructure/service/speech/openAi"
)

// Injectors from wire.go:

func InitialiseMessaging() (*usecase.Messaging, error) {
	telegramConfig := config.ProvideTelegramBotConfig()
	telegram, err := messenger.NewTelegram(telegramConfig)
	if err != nil {
		return nil, err
	}
	memory := cache.NewMemory()
	clientConfig := config.ProvideOpenAiClientConfig()
	client := openAi.NewClient(clientConfig)
	speech := openAi2.NewSpeech(client)
	chat := openAi3.NewChat(client)
	text := openAi3.NewText(client)
	image := openAi3.NewImage(client)
	openAiText := openAi4.NewText(client)
	code := openAi4.NewCode(client)
	openAiImage := openAi4.NewImage(client)
	messaging := usecase.NewMessaging(telegram, memory, speech, chat, text, image, openAiText, code, openAiImage)
	return messaging, nil
}

// wire.go:

var cacheSet = wire.NewSet(cache.NewMemory, wire.Bind(new(service.Cache), new(*cache.Memory)))

var messengerSet = wire.NewSet(config.ProvideTelegramBotConfig, messenger.NewTelegram, wire.Bind(new(service.Messenger), new(*messenger.Telegram)))

var openAiSet = wire.NewSet(config.ProvideOpenAiClientConfig, openAi.NewClient, openAi3.NewChat, wire.Bind(new(generator.Chat), new(*openAi3.Chat)), openAi3.NewText, wire.Bind(new(generator.Text), new(*openAi3.Text)), openAi3.NewImage, wire.Bind(new(generator.Image), new(*openAi3.Image)), openAi4.NewText, wire.Bind(new(editor.Text), new(*openAi4.Text)), openAi4.NewCode, wire.Bind(new(editor.Code), new(*openAi4.Code)), openAi4.NewImage, wire.Bind(new(editor.Image), new(*openAi4.Image)), openAi2.NewSpeech, wire.Bind(new(service.Speech), new(*openAi2.Speech)))
