---
title: Go Windows 系统下隐藏 CMD 窗口
date: 2021-07-13 17:28:42
top_img: https://s3.jpg.cm/2021/07/13/IQtCz4.jpg
cover: https://s3.jpg.cm/2021/07/13/IQtCz4.jpg
tags: 
    - Golang
---

### 隐藏 Golang 自身运行窗口

仅需在编译是增加 `-ldflags` 参数
```
go build -ldflags -H=windowsgui 
```

### 隐藏 Golang 调用 CMD 窗口

```go
cmd := exec.Command("ping", "127.0.0.1")
if runtime.GOOS == "windows" {
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
```