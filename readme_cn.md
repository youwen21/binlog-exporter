# mysql binlog exporter
为 Prometheus 上报 Mysql binlog 指标.


## Installation and Usage
运行项目前先根据 config.env.examples 生成config.env文件
编辑配置config.env文件中的配置项

### Building and running the exporter
```bash
git clone git@gitee.com:youwen21/binlog-exporter.git
cd binlog_exporter
go build .
./binlog_exporter

-------

git clone git@gitee.com:youwen21/binlog-exporter.git
cd binlog_exporter
go run main.go

```

### docker
// todo make docker image

## metrics
默认metric搜集地址：
> http://127.0.0.1:9900/metrics

- binlog_count
- binlog_sum
- binlog_his
- binlog_size

# other mysql binlog project











