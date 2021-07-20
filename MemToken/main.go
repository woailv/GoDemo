package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	token := TVLoginGenToken()
	go func() {
	HERE:
		for {
			status := GetTokenStatus(token)
			fmt.Println("status:", status)
			switch status {
			case TVLoginStatusSuccess, TVLoginStatusOutOfTime:
				break HERE
			}
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	time.Sleep(time.Second)
	TVLoginByMobile(token, "1")
	wg.Wait()
}

const TVLoginStatusNotExist = "notExist"
const TVLoginStatusWait = "wait"
const TVLoginStatusLogging = "logging"
const TVLoginStatusSuccess = "success"
const TVLoginStatusFailed = "failed"
const TVLoginStatusOutOfTime = "outOfTime"

var gTokenToStatusMap = &sync.Map{}

func GetTokenStatus(token string) string {
	v, ok := gTokenToStatusMap.Load(token)
	if !ok {
		return TVLoginStatusNotExist
	}
	lastStatus := v.(string)
	// 最终状态直接返回
	switch lastStatus {
	case TVLoginStatusSuccess:
		return TVLoginStatusSuccess
		//...
	}
HERE:
	v, ok = gTokenToStatusMap.Load(token)
	if !ok {
		return TVLoginStatusNotExist
	}
	currentStatus := v.(string)
	if lastStatus != currentStatus {
		return currentStatus
	} else {
		time.Sleep(time.Millisecond * 200)
		goto HERE
	}
}

func TVLoginByMobile(token string, mobileDeviceId string) error {
	time.Sleep(time.Second * 5)
	_, ok := gTokenToStatusMap.Load(token)
	if !ok {
		return errors.New("tokenNotExist")
	}
	gTokenToStatusMap.Store(token, TVLoginStatusLogging)
	// 判断是否允许登录
	if mobileDeviceId == "" {
		gTokenToStatusMap.Store(token, TVLoginStatusFailed)
		gTokenToStatusMap.Delete(token)
		return errors.New("mobileAuthFailed")
	}
	gTokenToStatusMap.Store(token, TVLoginStatusSuccess)
	return nil
}

var i = 0

func TVLoginGenToken() string {
	i++
	token := strconv.Itoa(i)
	gTokenToStatusMap.Store(token, TVLoginStatusWait)
	return token
}
