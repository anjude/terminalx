package base

type IHandler interface {
	// GetCommand 获取handler的命令
	GetCommand() (command string)
	// GetDesc 获取该命令的详情
	GetDesc() (desc string)
	// GetArgs 获取该命令需要的参数
	GetArgs(args []string) (curArgs []string, nextArgs []string)
	// Handle 该命令的具体执行逻辑
	Handle(args []string)
}
