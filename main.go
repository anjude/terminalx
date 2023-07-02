package main

import (
	"fmt"
	"github.com/anjude/terminalx/config"
	"github.com/anjude/terminalx/handler"
	"os"
)

func main() {
	// 解析命令行参数
	args := os.Args[1:]
	if len(args) == 0 {
		// 如果没有指定任何参数，则显示帮助信息
		fmt.Println("Usage: bot [options]")
		fmt.Println("Use -h for more information.")
	}

	// 加载配置
	if err := config.InitConfig(); err != nil {
	}

	// 执行命令
	handler.Execute(args)
}
