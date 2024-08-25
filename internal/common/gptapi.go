package common

import (
	"ale-chatGPT/environment"
	goOpenai "github.com/sashabaranov/go-openai"
)

type GPT struct {
	Client *goOpenai.Client
}

func NewGPTClient(env environment.Config) *GPT {
	return &GPT{
		Client: goOpenai.NewClient(env.Openai.ApiKey),
	}
}
