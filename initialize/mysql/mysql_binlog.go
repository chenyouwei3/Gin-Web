package mysql

import (
	"context"
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/sirupsen/logrus"
	"loopy-manager/initialize/messageQueue"
	"os"
	"os/signal"
	"syscall"
)

func MysqlBinlogInit() {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "43.138.32.203",
		Port:     3301,
		User:     "root",
		Password: "Cyw123456",
	}
	// 创建 BinlogSyncer 实例
	syncer := replication.NewBinlogSyncer(cfg)
	// 获取当前 MySQL 实例的最新 binlog 位置
	pos := syncer.GetNextPosition()

	// 开始监听 binlog 事件
	streamer, err := syncer.StartSync(pos)
	if err != nil {
		logrus.Error("无法启动 binlog 同步器：%v", err)
	}
	// 捕获退出信号，优雅地关闭程序
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		fmt.Println("接收到退出信号，正在关闭程序...")
		syncer.Close()
		os.Exit(0)
	}()
	for {
		ev, err := streamer.GetEvent(context.Background())
		if err != nil {
			logrus.Error("获取 binlog 事件失败：%v", err)
		}
		switch e := ev.Event.(type) {
		case *replication.RowsEvent:
			switch ev.Header.EventType {
			case replication.WRITE_ROWS_EVENTv1, replication.WRITE_ROWS_EVENTv2:
				messageQueue.RabbitCache.PublishSimple(string(e.Table.Table))
			case replication.UPDATE_ROWS_EVENTv1, replication.UPDATE_ROWS_EVENTv2:
				messageQueue.RabbitCache.PublishSimple(string(e.Table.Table))
			case replication.DELETE_ROWS_EVENTv1, replication.DELETE_ROWS_EVENTv2:
				messageQueue.RabbitCache.PublishSimple(string(e.Table.Table))
			}
		case *replication.QueryEvent:
			logrus.Error("SQL 查询语句：%s\n", e.Query)
		default:
			logrus.Error("未知的事件类型：%T\n", ev.Event)
		}
	}
}
