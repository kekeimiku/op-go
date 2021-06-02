# WIP 活跃开发中...

- [ ] 实现op所有的方法调用
- [ ] 详细简洁的注释文档 
- [ ] 一键打包所有静态文件

# 例子

```go
package main

import (
	"fmt"
	"github.com/kekeimiku/op-go"
)

func main() {
	op, _ := op.LoadOP()
	fmt.Println("当前版本", op.Ver())
}
```