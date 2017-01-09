package main

import (
	"flag"
	"fmt"
	
	"bufio"
	"io"
	"os"
	"strconv"
)

//*string 指针 (输入变量 如 ./sorter -i unsorted.dat -o sorted.dat -a bubblesort)
var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()
	
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	
	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}
}

func readValues(infile string)(values []int, err error) {
	file, err := os.Open(infile)//打开文件
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
	}
	
	defer file.Close()//下面发生错误时关闭文件
	
	br := bufio.NewReader(file)
	
	values = make([]int, 0)
	
	for {
		line, isPrefix, err1 := br.ReadLine()
		
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		
		str := string(line) //转换字符数组为字符串
		
		//strconv: http://www.cnblogs.com/golove/p/3262925.html
		value, err1 := strconv.Atoi(str)
		
		if err1 != nil {
			err = err1
			return
		}
		
		//为数组切片中添加元素
		values = append(values, value)
	}
	return
}