### protoc编译命令

`protoc  --go_out=plugins=grpc:./pkg/pb/ ./pkg/proto/im.proto
`

### 发送文本消息
`{"otherSideId":"2","messageType":1,"otherSideType":1,"action":"sendMessage","messageContent":"你好","messageClientId":"58a78cf65856466b9956d00411f878e9"}`

### go build方法
参考https://blog.csdn.net/baidu_34122324/article/details/90183394


windows打包命令
```
set GOARCH=amd64

set GOOS=linux

go build main.go
```
linux修改权限

```
chmod 773 [文件名]
```