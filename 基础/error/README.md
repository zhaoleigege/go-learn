### Error
1. 原则
   不能忽略掉错误返回值，并且一直返回错误值到最上层的调用处，并且只在最上层的调用出打印日志。

2. 新建一个error
   ```go
   import (
	"errors"
	"fmt"
    )

    func main() {
	    err := errors.New("测试错误")
	    fmt.Println(err)
    }
   ```
3. 调用的时候记录错误堆栈追踪
   
