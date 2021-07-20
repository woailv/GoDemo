package main

import (
	"fmt"
	"net"
	"os/exec"
	"regexp"
)

func main() {
	fmt.Println(getMACAddress())
	fmt.Println("2:", getMac())
}

func getMACAddress() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)
				//fmt.Println(ipnet.IP)
				if ok && ipnet.IP.IsGlobalUnicast() {
					// 如果IP是全局单拨地址，则返回MAC地址
					return netInterfaces[i].HardwareAddr.String()
				}
			}
		}
	}
	return ""
}

func getMac() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr.String() //获取本机MAC地址
		fmt.Println("agc:", mac)
	}
	//	fmt.Println("MAC = ", mac)
	return ""
}

func getCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.Println(string(out))
	str := string(out)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}
