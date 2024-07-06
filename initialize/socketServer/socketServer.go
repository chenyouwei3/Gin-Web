package socketServer

import (
	"fmt"
	controller "loopy-manager/app/controller/socket"
	"loopy-manager/app/model"
	"net"
	"sync"
)

func SocketServerStart() {
	taskQueue := make(chan model.SocketServerTask) // 创建任务队列
	done := make(chan bool)                        // 创建完成信号通道
	workerCount := 3                               // 设置工作 goroutine 的数量
	go sendSocketServerTasks(taskQueue)            // 启动生产任务的 goroutine
	// 启动一定数量的工作 goroutine 来处理任务
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go socketServerWorker(taskQueue, &wg)
	}
	// 等待工作 goroutine 完成处理
	go func() {
		wg.Wait()
		done <- true
	}()
	<-done // 等待完成信号

}

// 发送任务到通道中
func sendSocketServerTasks(taskQueue chan<- model.SocketServerTask) {
	SliceTask := []model.SocketServerTask{
		{
			Address:     "0.0.0.0:8004",
			ProcessFunc: controller.Test1,
		},
		{
			Address:     "0.0.0.0:8010",
			ProcessFunc: controller.Test2,
		},
		{
			Address:     "0.0.0.0:8019",
			ProcessFunc: controller.Test3,
		},
	}
	for _, task := range SliceTask {
		taskQueue <- task
	}
	close(taskQueue)
}

// 工作协程处理任务
func socketServerWorker(taskQueue <-chan model.SocketServerTask, wg *sync.WaitGroup) {
	for task := range taskQueue {
		processSocketServerTaskForTCP(task, task.ProcessFunc)
	}
	wg.Done()
}

//第一次()

// 创建连接(tcp)
func processSocketServerTaskForTCP(task model.SocketServerTask, Func func(conn net.Conn)) {
	listen, err := net.Listen("tcp", task.Address) //代表监听的地址端口
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	fmt.Println("正在等待建立连接.....", listen.Addr())
	for { //这个for循环的作用是可以多次建立连接
		conn, err := listen.Accept() //请求建立连接，客户端未连接就会在这里一直等待
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Println("连接建立成功.....")
		go Func(conn)
	}
}

// 创建连接(tcp)
func processSocketServerTaskForUDP(task model.SocketServerTask, Func func(conn net.Conn)) {
	listen, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	//fmt.Println("正在等待建立连接.....", listen.Addr())
	//for { //这个for循环的作用是可以多次建立连接
	//	conn, err := listen.Accept() //请求建立连接，客户端未连接就会在这里一直等待
	//	if err != nil {
	//		fmt.Println("accept failed, err:", err)
	//		continue
	//	}
	//	fmt.Println("连接建立成功.....")
	//	go Func(conn)
	//}
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err: ", err)
			continue
		}
		fmt.Println("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
	}

}
