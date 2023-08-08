package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	size    int
	count   int
	icmp    *ICMP = &ICMP{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Id:          1,
		SequenceNum: 1,
	}
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Id          uint16
	SequenceNum uint16
}

// 解析命令行参数
func parseCommandArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长，单位ms")
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小，单位byte")
	flag.IntVar(&count, "n", 4, "请求发送次数")
	flag.Parse()
}

func main() {
	parseCommandArgs()
	fmt.Println(timeout)
	fmt.Println(size)
	fmt.Println(count)
	destIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", destIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	data := make([]byte, size)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	buffer.Write(data)
	data = buffer.Bytes()
	fmt.Println(data)
}
