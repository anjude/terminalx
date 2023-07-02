package chatgpt

import (
	"github.com/sashabaranov/go-openai"
	"testing"
)

func TestChat(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{msg: "hello"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Chat(tt.args.msg, []openai.ChatCompletionMessage{})
		})
	}
}
