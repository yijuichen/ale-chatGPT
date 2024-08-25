//go:build wireinject
// +build wireinject

package gpt

import (
	"ale-chatGPT/environment"
	"ale-chatGPT/internal/common"
	"ale-chatGPT/internal/usecase"
	"context"
	"github.com/google/wire"
)

type App struct {
	Config    environment.Config
	GPTClient *common.GPT
	Usecase   usecase.InfUsecase
}

func Initialize(ctx context.Context, path string) (App, error) {
	wire.Build(
		wire.Struct(new(App), "*"),
		environment.New,
		common.NewGPTClient,
		usecase.NewGPTUsecase,
	)
	return App{}, nil
}
