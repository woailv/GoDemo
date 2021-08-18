package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	write()
	f := open()
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	end := info.Size()
	r := bufio.NewReader(f)
	var start int64
	var readLen int64
	flag := false
	var suf []byte
	delData := "0\n"
	for {
		start = readLen
		data, err := r.ReadSlice('\n')
		readLen += int64(len(data))
		if string(data) == delData {
			flag = true
			suf = make([]byte, end-int64(len(data)))
			_, err = r.Read(suf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			break
		}
	}
	if flag {
		_, err := f.Seek(start, 0)
		if err != nil {
			panic(err)
		}
		_, err = f.Write(suf)
		if err != nil {
			panic(err)
		}
		err = f.Truncate(end - int64(len(delData)))
		if err != nil {
			panic(err)
		}
	}
}

var name = "./temp.txt"

func write() {
	os.Remove(name)
	f := open()
	f.WriteString("1\n")
	f.WriteString("2\n")
	f.WriteString("3\n")
	f.WriteString("4\n")
	f.WriteString("5\n")
	f.Close()
}

func open() *os.File {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return f
}
