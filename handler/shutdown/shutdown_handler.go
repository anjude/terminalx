package shutdown

import (
	"fmt"
	"github.com/anjude/terminalx/handler/base"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var _ base.IHandler = Handler{}

type Handler struct {
	Command string
	Desc    string
}

func (s Handler) GetCommand() (command string) {
	return s.Command
}

func (s Handler) GetDesc() (desc string) {
	return s.Desc
}

func (s Handler) GetArgs(args []string) (curArgs []string, nextArgs []string) {
	return curArgs, args
}

func (s Handler) Handle(args []string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		if is64Bit() {
			cmd = exec.Command("shutdown", "/s", "/t", "30")
		} else {
			cmd = exec.Command("shutdown", "/s", "/t", "30", "/f")
		}
	case "linux":
		if is64Bit() {
			cmd = exec.Command("shutdown", "-h", "+30")
		} else {
			cmd = exec.Command("shutdown", "-h", "+30", "-f")
		}
	case "darwin":
		cmd = exec.Command("sudo", "shutdown", "-h", "+30")
	default:
		fmt.Println("unsupported platform:", runtime.GOOS)
		os.Exit(1)
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to shutdown:", err)
	}
}

func is64Bit() bool {
	return strings.Contains(runtime.GOARCH, "64")
}
