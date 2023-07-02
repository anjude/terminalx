package handler

import (
	"fmt"
	"github.com/anjude/terminalx/handler/base"
	"github.com/anjude/terminalx/handler/chat"
	"github.com/anjude/terminalx/handler/shutdown"
	"github.com/anjude/terminalx/handler/version"
)

var OptionMap = map[string]base.IHandler{
	"-h": Handler{
		Command: "-h",
		Desc:    "show help info",
	},
	"-v": version.Handler{
		Command: "-v",
		Desc:    "show version info",
	},
	"-c": chat.Handler{
		Command: "-q",
		Desc:    "chat with the bot",
	},
	"-shutdown": shutdown.Handler{
		Command: "-shutdown",
		Desc:    "shutdown this computer",
	},
}

func Execute(args []string) {
	if len(args) == 0 {
		return
	}
	handler, ok := OptionMap[args[0]]
	if !ok {
		fmt.Printf("command [%v] not found\n", args[0])
		return
	}
	curArgs, nextArgs := handler.GetArgs(args[1:])
	handler.Handle(curArgs)
	Execute(nextArgs)
}
