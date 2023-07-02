package handler

import (
	"fmt"
	"github.com/anjude/terminalx/handler/base"
)

var _ base.IHandler = Handler{}

type Handler struct {
	Command string
	Desc    string
}

func (h Handler) GetCommand() string {
	return h.Command
}

func (h Handler) GetDesc() string {
	return h.Desc
}

func (h Handler) GetArgs(args []string) (curArgs []string, nextArgs []string) {
	return curArgs, args
}

func (h Handler) Handle(args []string) {
	msg := "bot [option] args [option] args\n"
	for command, handler := range OptionMap {
		msg += fmt.Sprintf("	[%v] %v\n", command, handler.GetDesc())
	}
	fmt.Println(msg)
}
