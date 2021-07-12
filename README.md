## WIP 活跃开发中...

- [x] 实现op所有的方法调用 截至2021-06-03 op wiki中所有提到的接口都已实现
- [ ] 详细简洁的注释文档 
- [ ] 一键打包所有静态文件
- [ ] 包含基本功能的图形界面api

## 例子

```go
package main

import (
	"fmt"
	"github.com/kekeimiku/op-go"
)

func main() {
	op, _ := op.Load()
	fmt.Println("当前版本", op.Ver())
}
```

一个低配版后台鼠标模式的“自瞄”小游戏例子
![aimbooster](./_example/example.gif)

视频里的小游戏是 http://www.aimbooster.com/

## ChangeLog

### 2021-07-10
删除了大部分注释和cgo相关代码，注释是根据wiki生成的错误太多，准备抽时间重写，mavc的cgo有点难搞，放弃了，cgo和com感觉性能开销影响不大。（2021-07-12）基本确定了标准cgo性能开销暂时不如com...

### !!!2021-06-11
发现op-go注释里面不少东西都写错啦!!!不建议看注释，简直直接看op的wiki，我会尽快修复

### 2021-06-11
尝试使用cgo方式调用
