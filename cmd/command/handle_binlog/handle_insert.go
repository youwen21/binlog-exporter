package handle_binlog

import (
	"binlog_exporter/binlog_metrics"
	"github.com/go-mysql-org/go-mysql/replication"
)

func handleWriteRowsEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)

	binlog_metrics.InsertAction(e, ev)

}
