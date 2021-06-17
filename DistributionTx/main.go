package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// https://xiaomi-info.github.io/2020/01/02/distributed-transaction/
//消息表实现 分布式事务
//当系统 A 被其他系统调用发生数据库表更操作，首先会更新数据库的业务表，其次会往相同数据库的消息表中插入一条数据，两个操作发生在同一个事务中
//系统 A 的脚本定期轮询本地消息往 mq 中写入一条消息，如果消息发送失败会进行重试
//系统 B 消费 mq 中的消息，并处理业务逻辑。如果本地事务处理失败，会在继续消费 mq 中的消息进行重试，如果业务上的失败，可以通知系统 A 进行回滚操作
func main() {
	wg.Add(2)
	go handleMsg()
	go retry()
	business()
	wg.Wait()
	log.Println(list)
}

var wg = sync.WaitGroup{}

var msgChan = make(chan int, 10)

type Item struct {
	Num    int
	Status int //1 等待回调 2 发消息失败 3 处理成功
}

var retryChan = make(chan int, 10)

var list []Item

func business() {
	txFn := func(i int) {
		status := 1
		if i%2 == 0 { //模拟发消息失败, 也可能会成功之后中断, 需提供补偿机制
			retryChan <- i
			status = 2
		} else {
			msgChan <- i
		}
		list = append(list, Item{
			Num:    i,
			Status: status,
		})
	}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		txFn(i)
	}
	close(retryChan)
}

func handleMsg() {
	for msg := range msgChan {
		for k, item := range list {
			if item.Num == msg {
				fmt.Println("handle item:", item)
				item.Status = 3
				list[k] = item
			}
		}
	}
	wg.Done()
}

func retry() {
	for msg := range retryChan {
		msgChan <- msg
		fmt.Println("retry msg:", msg)
	}
	close(msgChan)
	wg.Done()
}
