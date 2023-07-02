package chatgpt

import (
	"context"
	"errors"
	"fmt"
	"github.com/anjude/terminalx/config"
	"github.com/anjude/terminalx/third_party/sizhi"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"net/url"
)

func Chat(msg string, dialog []openai.ChatCompletionMessage) []openai.ChatCompletionMessage {
	clientConfig := openai.DefaultConfig(config.BotConf.ChatGPT.ApiKey)
	proxyUrl, err := url.Parse(config.BotConf.ChatGPT.Proxy)
	if err != nil {
		fmt.Println(err)
		return dialog
	}
	clientConfig.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	client := openai.NewClientWithConfig(clientConfig)

	dialog = append(dialog, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})
	ctx := context.Background()
	stream, err := client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 1000,
		Stream:    true,
		Messages:  dialog,
	})
	if err != nil {
		sizhiMsg, _ := sizhi.GetSizhiMsg(msg, openai.ChatMessageRoleUser)
		fmt.Println("ChatGPT报错，切换小思: ", sizhiMsg)
		return getReturn(dialog, sizhiMsg)
	}
	defer stream.Close()

	answer := ""
	fmt.Printf("bot: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return getReturn(dialog, answer)
		}
		if err != nil {
			fmt.Println(err)
			return dialog
		}
		content := response.Choices[0].Delta.Content
		answer += content
		fmt.Printf(content)
	}
}

func getReturn(dialog []openai.ChatCompletionMessage, answer string) []openai.ChatCompletionMessage {

	dialog = append(dialog, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: answer,
	})
	chatRounds := config.BotConf.ChatGPT.ChatRounds
	if len(dialog) >= chatRounds {
		dialog = dialog[len(dialog)-chatRounds:]
	}
	return dialog
}
