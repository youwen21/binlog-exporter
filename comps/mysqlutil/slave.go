package mysqlutil

import (
	"binlog_exporter/conf"
	"database/sql"
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
)

type BinLogPosition struct {
	File            string
	Position        uint32
	BinlogDoDb      string
	BinlogIgnoreDb  string
	ExecutedGtidSet string
}

func GetMysqlPosition(dbConf conf.Mysql) (*BinLogPosition, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local&parseTime=true",
		dbConf.Username,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database,
		dbConf.Charset,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rows := db.QueryRow("show master status")
	if rows.Err() != nil {
		panic(rows.Err())
	}

	pos := new(BinLogPosition)
	err = rows.Scan(&pos.File, &pos.Position, &pos.BinlogDoDb, &pos.BinlogIgnoreDb, &pos.ExecutedGtidSet)
	if err != nil {
		panic(err)
	}
	return pos, nil
}

func GetBinlogSyncerConfig(conf conf.Mysql) replication.BinlogSyncerConfig {
	cfg := replication.BinlogSyncerConfig{
		ServerID: uint32(conf.ServerId),
		Flavor:   "mysql",
		Host:     conf.Host,
		Port:     uint16(conf.Port),
		User:     conf.Username,
		Password: conf.Password,
	}

	return cfg
}
