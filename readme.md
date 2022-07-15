# mysql binlog exporter
Prometheus exporter for Mysql binlog metrics.  
[中文文档](readme_cn.md)

## Installation and Usage
cp config.env.examples config.env and edit the config.env file before running

### Building and running the exporter
#### github
```bash
git clone git@github.com:youwen21/binlog-exporter.git
```
#### gitee
```bash
git clone git@gitee.com:youwen21/binlog-exporter.git
```

```bash
cd binlog_exporter
go run main.go

```

### docker
// todo make docker image


## metrics
default metric url：
> http://127.0.0.1:9900/metrics

 - binlog_count
 - binlog_sum
 - binlog_his
 - binlog_size


# other mysql binlog project

## github

## gitee











