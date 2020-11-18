### channel相关的用法
1. 关闭channel
```go
// 关闭channel时一定得使用方法close
// 这样才可以使得所有接收该channel的监听者能够接收到关闭消息
stop := make(chan struct{})
close(stop)
```