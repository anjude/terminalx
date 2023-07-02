### TerminalX
terminal x

### 简单使用
1. 准备好golang环境
2. 下载运行项目
```shell
git clone https://github.com/anjude/terminalx.git
cd terminalx
go mod tidy
go build -o bot.exe main.go
```
3. 把生成的bot.exe（mac、linux没有.exe扩展名）的目录放到环境变量下
4. 修改config.yaml文件，需要设置自己的api-key，按需设置代理端口
5. 在终端运行bot -h查看帮助

![img.png](doc/img.png)