<!--
 * @Author: Kanri
 * @Date: 2022-01-11 16:28:57
 * @LastEditors: Kanri
 * @LastEditTime: 2022-07-16 19:47:52
 * @Description: 
-->
---
title: Go Windows 系统下隐藏 CMD 窗口
date: 2021-07-13 17:28:42
top_img: https://pic1.zhimg.com/80/v2-5a6b5833d6e69a064f1b6cd48d19d4ec_720w.png
cover: https://pic1.zhimg.com/80/v2-5a6b5833d6e69a064f1b6cd48d19d4ec_720w.png
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