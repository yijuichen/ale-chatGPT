package usecase

import (
	"ale-chatGPT/internal/common"
	"context"
	"github.com/sashabaranov/go-openai"
)

type usecase struct {
	gpt *common.GPT
	//logger   *log.Logger
}

func (u *usecase) GetGPTResponse(ctx context.Context, prompt string) (string, error) {

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4oMini,
		MaxTokens: 400,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt,
			},
		},
	}

	response, err := u.gpt.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	content := response.Choices[0].Message.Content
	return content, nil

}

type InfUsecase interface {
	GetGPTResponse(ctx context.Context, prompt string) (string, error)
}

func NewGPTUsecase(gpt *common.GPT) InfUsecase {
	return &usecase{
		gpt: gpt,
		//logger:   logger,
	}
}
