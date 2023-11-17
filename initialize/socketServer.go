package initialize

import (
	"fmt"
	"loopy-manager/controller"
	"net"
)

type Task struct {
	Address     string
	ProcessFunc func(conn net.Conn)
}

// 模拟发送任务到通道中
func sendTasks(taskQueue chan<- Task) {
	SliceTask := []Task{
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
}

func SocketServer() {
	taskQueue := make(chan Task) // 创建任务队列
	workerCount := 3             // 设置工作 goroutine 的数量
	go sendTasks(taskQueue)      // 启动生产任务的 goroutine
	for i := 0; i < workerCount; i++ {
		go worker(taskQueue)
	}
}

// 工作协程处理任务
func worker(taskQueue <-chan Task) {
	for task := range taskQueue {
		processTask(task.Address, task.ProcessFunc)
	}
}

// 处理任务
func processTask(Address string, Func func(conn net.Conn)) {
	listen, err := net.Listen("tcp", Address) //代表监听的地址端口
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
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
