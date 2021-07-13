---
title: CGO 类型转换
date: 2021-07-14 22:11:16
top_img: https://s3.jpg.cm/2021/07/13/IQEvIu.jpg
cover: https://s3.jpg.cm/2021/07/13/IQEvIu.jpg
tags: 
    - Golang
    - C
    - CGO
---

### 基本数值类型

* `Go` 的基本数值类型内存模型和 `C` 一样，传递数值类型时可以直接将 `Go` 的基本数值类型转换成对应的 `CGO` 类型然后传递给 `C` 函数调用。

* `Go` 和 `C` 的基本数值类型转换对照表如下:

    | C语言类型              | CGO类型     | Go语言类型 |
    | :--------------------- | :---------- | :--------- |
    | char                   | C.char      | byte       |
    | singed char            | C.schar     | int8       |
    | unsigned char          | C.uchar     | uint8      |
    | short                  | C.short     | int16      |
    | unsigned short         | C.ushort    | uint16     |
    | int                    | C.int       | int32      |
    | unsigned int           | C.uint      | uint32     |
    | long                   | C.long      | int32      |
    | unsigned long          | C.ulong     | uint32     |
    | long long int          | C.longlong  | int64      |
    | unsigned long long int | C.ulonglong | uint64     |
    | float                  | C.float     | float32    |
    | double                 | C.double    | float64    |
    | size_t                 | C.size_t    | uint       |

### 字符串
使用 `"C"` 提供的 `C.CString()` 将 `Go` 的字符串转换成 `C` char * 然后传递给 `C` 函数调用
```go
package main

/*
#include <stdio.h>
#include <stdlib.h>

void hello(char *str) {
	printf(str);
}
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("hello world")
	C.hello(cstr)
	C.free(unsafe.Pointer(cstr))
}
```
其中，`C.CString()` 返回的 `C` 字符串是在堆上新创建的并且不受 Go 的 GC 的管理，使用完后需要自行调用 `C.free()`，否则会造成内存泄露。