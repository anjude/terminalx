package base

type IHandler interface {
	GetCommand() (command string)
	GetDesc() (desc string)
	GetArgs(args []string) (curArgs []string, nextArgs []string)
	Handle(args []string)
}
