package binlog_metrics

import (
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var labelTags = []string{"schema", "table", "event"}

var EventCount = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "binlog_count"}, labelTags)
var DurationSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{Name: "binlog_sum"}, labelTags)
var DurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "binlog_his"}, labelTags)
var SizeSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{Name: "binlog_size"}, labelTags)

func init() {
	prometheus.MustRegister(EventCount)
	prometheus.MustRegister(DurationHistogram)
	prometheus.MustRegister(DurationSummary)
}

func InsertAction(e *replication.BinlogEvent, ev *replication.RowsEvent) {
	RowEvent(e, ev, "insert")
}

func UpdateAction(e *replication.BinlogEvent, ev *replication.RowsEvent) {
	RowEvent(e, ev, "update")
}

func DeleteAction(e *replication.BinlogEvent, ev *replication.RowsEvent) {
	RowEvent(e, ev, "delete")
}

func RowEvent(e *replication.BinlogEvent, ev *replication.RowsEvent, action string) {
	labels := map[string]string{
		"schema": string(ev.Table.Schema),
		"table":  string(ev.Table.Table),
		"event":  action,
	}

	du := float64(time.Now().Unix() - int64(e.Header.Timestamp))

	EventCount.With(labels).Inc()
	DurationHistogram.With(labels).Observe(du)
	DurationSummary.With(labels).Observe(du)
	SizeSummary.With(labels).Observe(float64(e.Header.EventSize))
}
