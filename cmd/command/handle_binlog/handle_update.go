package handle_binlog

import (
	"binlog_exporter/binlog_metrics"
	"github.com/go-mysql-org/go-mysql/replication"
)

func handleUpdateEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)

	binlog_metrics.UpdateAction(e, ev)

}
