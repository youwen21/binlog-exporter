package command

import (
	"binlog_exporter/cmd/command/handle_binlog"
	"binlog_exporter/comps/mysqlutil"
	"binlog_exporter/conf"
	"context"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
)

type binlogClient struct{}

var (
	BinlogClient = binlogClient{}
)

//StartBinlogClient 消费mysql binlog
func (bc binlogClient) StartBinlogClient(c *cli.Context) error {
	go func() {
		http.Handle(conf.Config.TelemetryPath, promhttp.Handler())
		err := http.ListenAndServe(conf.Config.ListenAddress, nil)
		if err != nil {
			panic(err)
		}
	}()

	masterPosition, err := mysqlutil.GetMysqlPosition(conf.Config.DefaultMysql)
	if err != nil {
		panic(err)
	}

	// 生成binlog消费实例
	cfg := mysqlutil.GetBinlogSyncerConfig(conf.Config.DefaultMysql)
	syncer := replication.NewBinlogSyncer(cfg)
	streamer, err := syncer.StartSync(mysql.Position{Name: masterPosition.File, Pos: masterPosition.Position})
	if err != nil {
		panic(err)
	}

	// 启动消费binlog
	for {
		ev, err := streamer.GetEvent(context.Background())
		if err != nil {
			log.Println("sync binlog error", err)
		}
		// Dump event
		//ev.Dump(os.Stdout)
		handle_binlog.HandleEvent(ev)
	}

	return nil
}
