package handle_binlog

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
)

func handleQueryEvent(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.QueryEvent)
	if string(ev.Schema) == "" {
		fmt.Println("event schema is empty")
		//e.Dump(os.Stdout)
		return
	}

	switch string(ev.Query) {
	case "BEGIN", "begin":
		return
	default:
	}

}
