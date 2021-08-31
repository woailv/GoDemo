package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	f, err := os.OpenFile("mmap.txt", os.O_RDWR|os.O_CREATE, 0644)
	if nil != err {
		log.Fatalln(err)
	}
	if _, err := f.WriteAt([]byte("1234"), 1); nil != err {
		log.Fatalln(err)
	}
	// 从文件描述符中读取数据
	data, err := syscall.Mmap(int(f.Fd()), 0, 100, syscall.PROT_WRITE, syscall.MAP_SHARED)
	if nil != err {
		log.Fatalln(err)
	}
	// 此处文件已经关闭了
	if err := f.Close(); nil != err {
		log.Fatalln(err)
	}
	fmt.Println("data:", string(data))
	for i, v := range []byte("ii") {
		data[i] = v
	}
	fmt.Println("data:", string(data))
	// 文件关闭了依然可以使用系统调用写数据 写入的数据不能超过原来的大小 只能使用 获取到的data 其它变量会报错
	if err := syscall.Munmap(data); nil != err {
		log.Fatalln(err)
	}
}
