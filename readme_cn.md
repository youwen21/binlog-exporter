# mysql binlog exporter
为 Prometheus 上报 Mysql binlog 指标.


## Installation and Usage
运行项目前先根据 config.env.examples 生成config.env文件
编辑配置config.env文件中的配置项

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
默认metric搜集地址：
> http://127.0.0.1:9900/metrics

- binlog_count
- binlog_sum
- binlog_his
- binlog_size

# other mysql binlog project

## binlog-spread 
一个基于mysql binlog的开发出来的程序员辅助工具

[gitee:binlog-spread ](https://gitee.com/youwen21/binlog-spread)
> https://gitee.com/youwen21/binlog-spread

[github:binlog-spread ](https://github.com/youwen21/binlog-spread)
> https://github.com/youwen21/binlog-spread







