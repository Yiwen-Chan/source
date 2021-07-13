---
title: CGO Go 与 C 的互相调用
date: 2021-07-13 21:47:32
top_img: https://s3.jpg.cm/2021/07/13/IQEvIu.jpg
cover: https://s3.jpg.cm/2021/07/13/IQEvIu.jpg
tags: 
    - Golang
	- C
    - CGO
---

`CGO` 提供了 `Golang` 和 `C` 相互调用的机制。通过 `CGO` 技术我们可以在 `Go` 中调用 `C` 函数，也可以将 `Go` 函数导出为 `C` 函数。

### Go 调用 C 函数
```go
package main

/*
int add(int a, int b) {
    return a + b;
}
*/
import "C"
import "fmt"

func main() {
	var a, b = 1, 2
	c := int(C.add(C.int(a), C.int(b)))
	fmt.Println(c) // 3
}
```
* 上述代码中的 `import C` ， `"C"` 是一个伪包，在 `Go` 的标准库中并不存在 `"C"` 包， `CGO` 通过这个查找到对应引用 `C` 命名空间的。

### C 调用 Go 函数

`main.go` 文件，`main` 函数中调用 `add.h` 文件中的 `C.test` 函数。
```go
// main.go
package main

//#include <add.h>
import "C"

func main() {
	C.test()
}
```

`add.go` 文件，定义并导出了 `Add` 函数，提供给 `add.h` 文件调用。其中 `//export Add` 是导出 `Add` 为 `C` 函数。
```go
// add.go
package main

import "C"

//export Add
func Add(a, b C.int) C.int {
	return a + b
}
```

`add.h` 文件，通过 `extern` 导入 `add.go` 文件定义的 `Add` 函数，并在 `C.test` 函数中调用。
```c
// add.h
#include <stdio.h>

extern int Add(int a, int b);

void test()
{
    printf("%d", Add(1, 2)); // 3
}
```

* 使用命令 `go build main.go add.go` 编译
