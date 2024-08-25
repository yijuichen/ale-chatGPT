package usecase

import (
	"ale-chatGPT/internal/common"
	"bytes"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"text/template"
)

// add enum for mod
type EnumMod string

const (
	ModeProductOwner EnumMod = "ProductOwner"
)

type TemplateConfig struct {
	Prompt string
}

type TemplateData map[string]interface{}

var modMap = map[EnumMod]*TemplateConfig{
	ModeProductOwner: &TemplateConfig{
		Prompt: `### Role
		A PM who is eloquent and good at communication can express ideas accurately using simple words
		### Skill
		- Skill 1: Good at communicating with software engineers
		- Skill 2: Uses simple and easy-to-understand words
		- Skill 3: Proficient in translation
		- Skill 4: Simply translated into feature title and commit message and more commit description
		### Constraints
		1. Output content in JSON format.
		2. Use English for output.

		The above is the configuration for this response
		===
		{{ .Question }}
		===
		run Skill 4 and  Skill 3`,
	},
}

type usecase struct {
	gpt *common.GPT
	//logger   *log.Logger
}

func (u *usecase) GetGPTResponse(ctx context.Context, msg string) (string, error) {

	td := TemplateData{
		"Question": msg,
	}
	result, tErr := ProcessTemplate(ModeProductOwner, td)
	if tErr != nil {
		return "", tErr
	}

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4oMini,
		MaxTokens: 400,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: result,
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

func ProcessTemplate(mod EnumMod, data TemplateData) (string, error) {
	config, exists := modMap[mod]
	if !exists {
		return "", fmt.Errorf("no configuration found for mode %s", mod)
	}

	tmpl, err := template.New("prompt").Parse(config.Prompt)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
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
