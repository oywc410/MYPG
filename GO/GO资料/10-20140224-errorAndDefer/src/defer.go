package src

import (
	"os"
)

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dstFile.Close()
	
	defer func() {
		// 做你复杂的清理工作
	} ()
	
	return io.Copy(dstFile, srcFile)	//即使copy错误 defer 中的语句照样被执行
	/*
		一个函数中可以存在多个defer语句，因此需要注意的是，defer语句的调用是遵照
		先进后出的原则，即最后一个defer语句将最先被执行
	*/
}