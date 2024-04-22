学生成绩管理系统
已经部署，可通过校园网预览
`10.12.180.78:3000`


本地部署oceanbase

下载
```
obd demo
```
每次开机需要启动集群
```
obd cluster restart demo
```

启动client

```
cd ./client
yarn i
yarn serve
```

启动serve

```
cd ./server
go mod tidy
go rn server.go
```