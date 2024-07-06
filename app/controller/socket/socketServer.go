package controller

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Test1(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(conn)
	for {
		reader := bufio.NewReader(conn) // 创建一个带缓冲的读取器，用于从连接中读取数据
		var buf [128]byte               // 创建一个长度为 128 的字节数组，用于存储读取的数据
		n, err := reader.Read(buf[:])   // 从读取器中读取数据，并将数据存储到 buf 中，同时返回读取的字节数和可能的错误
		if err != nil {                 // 若读取过程中发生错误
			fmt.Println("read from client failed, err:", err) // 打印错误信息
			break                                             // 结束循环，退出处理函数
		}
		recvStr := string(buf[:n])              // 将读取的字节转换为字符串
		fmt.Println("收到client端发来的数据：", recvStr) // 打印接收到的字符串数据
		conn.Write([]byte(recvStr))             // 将接收到的数据通过连接发送回客户端
	}
}

func Test2(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn) // 创建一个带缓冲的读取器，用于从连接中读取数据
		var buf [128]byte               // 创建一个长度为 128 的字节数组，用于存储读取的数据
		n, err := reader.Read(buf[:])   // 从读取器中读取数据，并将数据存储到 buf 中，同时返回读取的字节数和可能的错误
		if err != nil {                 // 若读取过程中发生错误
			fmt.Println("read from client failed, err:", err) // 打印错误信息
			break                                             // 结束循环，退出处理函数
		}
		recvStr := string(buf[:n])              // 将读取的字节转换为字符串
		fmt.Println("收到client端发来的数据：", recvStr) // 打印接收到的字符串数据

		conn.Write([]byte(recvStr)) // 将接收到的数据通过连接发送回客户端
	}
}

func Test3(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn) // 创建一个带缓冲的读取器，用于从连接中读取数据
		var buf [128]byte               // 创建一个长度为 128 的字节数组，用于存储读取的数据
		n, err := reader.Read(buf[:])   // 从读取器中读取数据，并将数据存储到 buf 中，同时返回读取的字节数和可能的错误
		if err != nil {                 // 若读取过程中发生错误
			fmt.Println("read from client failed, err:", err) // 打印错误信息
			break                                             // 结束循环，退出处理函数
		}
		recvStr := string(buf[:n])              // 将读取的字节转换为字符串
		fmt.Println("收到client端发来的数据：", recvStr) // 打印接收到的字符串数据
		conn.Write([]byte(recvStr))             // 将接收到的数据通过连接发送回客户端
	}
}
