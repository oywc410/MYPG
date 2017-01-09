package main

import "os"// 用于获得命令行参数os.Args
import "fmt"
import "simplemath"
import "strconv"

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ... ")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.");
}

func main() {
	args := os.Args
	
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	
	switch args[1] {
		case "add":
			if len(args) != 4 {
				fmt.Println("USAGE: calc add <integer1><integer2>")
				return
			}
			v1, err1 := strconv.Atoi(args[2])
			v2, err2 := strconv.Atoi(args[3])
			if err1 != nil || err2 != nil {
				fmt.Println("USAGE: calc add <integer1><integer2>")
			}
			ret := simplemath.Add(v1, v2)
			fmt.Println("Result: ", ret)
			break
		case "sqrt":
			if len(args) != 3 {
				fmt.Println("USAGE: calc sqrt <integer>")
				return
			}
			v, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("USAGE: calc sqrt <integer>")
				return
			}
			ret := simplemath.Sqrt(v)
			fmt.Println("Result: ", ret);
			break
		default:
			Usage()
	}
}


/**

生成可执行文件:

	为了能够构建这个工程，需要先把这个工程的根目录加入到环境变量GOPATH中。假设calcproj
	目录位于~/goyard下，则应编辑~/.bashrc文件，并添加下面这行代码：
	export GOPATH=~/goyard/calcproj
	然后执行以下命令应用该设置：
	$ source ~/.bashrc
	GOPATH和PATH环境变量一样，也可以接受多个路径，并且路径和路径之间用冒号分割。
	设置完GOPATH后，现在我们开始构建工程。假设我们希望把生成的可执行文件放到
	calcproj/bin目录中，需要执行的一系列指令如下：
	
	$ cd ~/goyard/calcproj
	$ mkdir bin
	$ cd bin
	$ go build calc
	顺利的话，将在该目录下发现生成的一个叫做calc的可执行文件
	
单元测试:

	(bin目录下)
	go test simplemath
	
打印日志:
	
	fmt.Println("The value of fval is", fval)
	fmt.Printf("fval=%f, ival=%d, sval=%s\n", fval, ival, sval)	
	
GDB 试调??????
*/