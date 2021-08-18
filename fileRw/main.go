package main

import (
	"os"
)

/*
先把文件syscall.Mmap 到内存，
查找到该行的起止位置，
然后用copy把后面的数据往前移
file.Truncate 缩短文件
syscall.Munmap
*/
func main() {
	f, err := os.OpenFile("./temp.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString("123456")
	if err != nil {
		panic(err)
	}

	f.Truncate(2)
}
