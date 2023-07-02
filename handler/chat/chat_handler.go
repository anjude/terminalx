package chat

import (
	"bufio"
	"fmt"
	"github.com/anjude/terminalx/handler/base"
	"github.com/anjude/terminalx/third_party/chatgpt"
	"github.com/sashabaranov/go-openai"
	"os"
)

var _ base.IHandler = Handler{}

type Handler struct {
	Command string
	Desc    string
}

func (c Handler) GetCommand() (command string) {
	return c.Command
}

func (c Handler) GetDesc() (desc string) {
	return c.Desc
}

func (c Handler) GetArgs(args []string) (curArgs []string, nextArgs []string) {
	if len(args) == 0 {
		fmt.Println("no chat content")
		return
	}
	return args[:1], args[1:]
}

func (c Handler) Handle(args []string) {
	var dialog []openai.ChatCompletionMessage
	fmt.Println(`[input "quit" to exit]`)
	if len(args) != 0 {
		dialog = chatgpt.Chat(args[0], dialog)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "quit" {
			break
		}
		dialog = chatgpt.Chat(text, dialog)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Goodbye!")
}
