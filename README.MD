x`x`protoc编译命令

`protoc  --go_out=plugins=grpc:../pb/ ./hello.proto
`

发送文本消息
`{"otherSideId":"2","messageType":1,"otherSideType":1,"action":"sendMessage","messageContent":"你好"}`