package version

import (
	"fmt"
	"github.com/anjude/terminalx/config"
	"github.com/anjude/terminalx/handler/base"
)

var _ base.IHandler = Handler{}

type Handler struct {
	Command string
	Desc    string
}

func (v Handler) GetCommand() string {
	return v.Command
}

func (v Handler) GetDesc() string {
	return v.Desc
}

func (v Handler) GetArgs(args []string) (curArgs []string, nextArgs []string) {
	return curArgs, args
}

func (v Handler) Handle(strings []string) {
	fmt.Println(config.BotConf.Version)
}
