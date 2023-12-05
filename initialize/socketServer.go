package initialize

import (
	"bufio"
	"fmt"
	"log"
	"loopy-manager/controller"
	"net"
)

type Task struct {
	Address     string
	BufSize     int
	ProcessFunc func(conn net.Conn, bufSize int)
	DataConfig
}

type DataConfig struct {
	DLen        int
	Bits        int
	Omit        int
	SensorHead  int
	CrcLen      int
	Complement  bool
	TimeFormat  bool
	SensorMerge bool
	Unit        bool
	ValidCode   bool
	Crc         bool
	Sprintf     string
	Report      string
}

// 配置任务发送
func sendTasks(taskQueue chan<- Task) {
	SliceTask := []Task{
		{
			Address:     "0.0.0.0:8004",
			BufSize:     1024,
			ProcessFunc: controller.Test1,
			DataConfig: DataConfig{
				DLen:        0,
				Bits:        0,
				Omit:        0,
				SensorHead:  0,
				CrcLen:      0,
				Complement:  false,
				TimeFormat:  false,
				SensorMerge: false,
				Unit:        false,
				ValidCode:   false,
				Crc:         false,
				Sprintf:     "",
				Report:      "",
			},
		},
		{
			Address:     "0.0.0.0:8010",
			BufSize:     2048,
			ProcessFunc: controller.Test2,
			DataConfig: DataConfig{
				DLen:        0,
				Bits:        0,
				Omit:        0,
				SensorHead:  0,
				CrcLen:      0,
				Complement:  false,
				TimeFormat:  false,
				SensorMerge: false,
				Unit:        false,
				ValidCode:   false,
				Crc:         false,
				Sprintf:     "",
				Report:      "",
			},
		},
		{
			Address:     "0.0.0.0:8019",
			BufSize:     4096,
			ProcessFunc: controller.Test3,
			DataConfig: DataConfig{
				DLen:        0,
				Bits:        0,
				Omit:        0,
				SensorHead:  0,
				CrcLen:      0,
				Complement:  false,
				TimeFormat:  false,
				SensorMerge: false,
				Unit:        false,
				ValidCode:   false,
				Crc:         false,
				Sprintf:     "",
				Report:      "",
			},
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
		processTask(task.Address, task.BufSize, task.ProcessFunc, task.DataConfig)
	}
}

// 处理任务1
func processTask(Address string, BufSize int, Func func(conn net.Conn, bufSize int), Data DataConfig) {
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
		go process(conn, BufSize, Data)
	}
}

// 处理任务2
func process(conn net.Conn, bufSize int, Data DataConfig) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(conn)
	for {
		reader := bufio.NewReader(conn) // 创建一个带缓冲的读取器，用于从连接中读取数据
		buf := make([]byte, bufSize)
		//var buf [bufSize]byte           // 创建一个长度为 128 的字节数组，用于存储读取的数据
		n, err := reader.Read(buf[:]) // 从读取器中读取数据，并将数据存储到 buf 中，同时返回读取的字节数和可能的错误
		if err != nil {               // 若读取过程中发生错误
			fmt.Println("read from client failed, err:", err) // 打印错误信息
			break                                             // 结束循环，退出处理函数
		}
		//recvStr := string(buf[:n])                                // 将读取的字节转换为字符串
		//fmt.Println(conn.LocalAddr(), "收到client端发来的数据：", recvStr) // 打印接收到的字符串数据
		//conn.Write([]byte(recvStr))                               // 将接收到的数据通过连接发送回客户端
		parseDTUDataNew(buf[:], n, Data)
	}
}

// 处理任务3
func parseDTUDataNew(buf []byte, n int, Data DataConfig) {

}
