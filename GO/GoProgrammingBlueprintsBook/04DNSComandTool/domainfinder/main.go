package main
import (
	"os/exec"
	"os"
	"log"
)

var cmdChain = []*exec.Cmd{
	//exec.Command("lib/synonyms"),
	exec.Command("lib/sprinkle"),
	exec.Command("lib/coolify"),
	exec.Command("lib/domainify"),
	exec.Command("lib/available"),
}

//命令行操作
func main() {
	//指定各个程序的输入输出流
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Panicln(err)
		}
		nextCmd.Stdin = stdout
	}

	//运行程序
	//cmd.Start 非阻塞形式运行
	//cmd.Run 阻塞好形式运行
	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Panicln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}

	//cmd.Wait 阻塞等待运行结束
	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Panicln(err)
		}
	}
}
