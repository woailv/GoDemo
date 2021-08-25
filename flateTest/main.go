package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// 一个缓存区压缩的内容
	buf := bytes.NewBuffer(nil)

	// 创建一个flate.Writer
	flateWrite, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		log.Fatalln(err)
	}
	defer flateWrite.Close()
	// 写入待压缩内容
	flateWrite.Write([]byte("akdflj看见对方"))
	flateWrite.Flush()
	fmt.Printf("len:%d,content：%s\n", buf.Len(), buf)
	// 解压刚压缩的内容
	flateReader := flate.NewReader(buf)
	defer flateWrite.Close()
	bs, err := ioutil.ReadAll(flateReader)
	if err != nil {
	}
	// 输出
	fmt.Print("解压后 len：", len(bs))
	fmt.Print("解压后 content：", string(bs))
}
