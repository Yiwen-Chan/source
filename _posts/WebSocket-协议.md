---
title: WebSocket 协议
date: 2021-06-26 07:20:07
top_img: https://pic3.zhimg.com/80/v2-8d5eb57c8d6e2b3867408688c5124aad_720w.png
cover: https://pic3.zhimg.com/80/v2-8d5eb57c8d6e2b3867408688c5124aad_720w.png
tags: 
    - HTTP
    - WebSocket
---

### WebSocket
WebSocket 是 HTML5 开始提供的一种在单个 TCP 连接上进行全双工通讯的协议。

### 连接握手

`WebSocket` 客户端发起基于 `HTTP` 握手的数据

```
GET / HTTP/1.1
Host: 127.0.0.1:8000
Connection: Upgrade
Upgrade: websocket
Origin: http://127.0.0.1:8000
Sec-WebSocket-Version: 13
Sec-WebSocket-Key: Bt4+Nfq12qxyxHslV2iFFg==
Sec-WebSocket-Protocol: chat
```
`WebSocket` 服务端响应握手的数据
```
HTTP/1.1 101 Switching Protocols
Sec-WebSocket-Accept: MK6YmuGMF81B+0zEjhayzUlnqxg=
Connection: Upgrade
Upgrade: websocket
```

### 基础帧协议

```
      0                   1                   2                   3
      0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
     +-+-+-+-+-------+-+-------------+-------------------------------+
     |F|R|R|R| opcode|M| Payload len |    Extended payload length    |
     |I|S|S|S|  (4)  |A|     (7)     |             (16/64)           |
     |N|V|V|V|       |S|             |   (if payload len==126/127)   |
     | |1|2|3|       |K|             |                               |
     +-+-+-+-+-------+-+-------------+ - - - - - - - - - - - - - - - +
     |     Extended payload length continued, if payload len == 127  |
     + - - - - - - - - - - - - - - - +-------------------------------+
     |                               |Masking-key, if MASK set to 1  |
     +-------------------------------+-------------------------------+
     | Masking-key (continued)       |          Payload Data         |
     +-------------------------------- - - - - - - - - - - - - - - - +
     :                     Payload Data continued ...                :
     + - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - +
     |                     Payload Data continued ...                |
     +---------------------------------------------------------------+

```
- FIN: 1 bit
​值为1则表示最后的`frame`。

- RSV1，RSV2，RSV3: 每个1 bit
必须设置为0，除非扩展了非0值含义的扩展。

- Opcode: 4 bit
`Payload data` 的操作码。
    + %x0 表示一个持续帧
    + %x1 表示一个文本帧
    + %x2 表示一个二进制帧
    + %x3-7 预留给以后的非控制帧
    + %x8 表示一个连接关闭包
    + %x9 表示一个ping包
    + %xA 表示一个pong包
    + %xB-F 预留给以后的控制帧

- Mask: 1 bit
是否使用掩码。如果设置为1，那么掩码的键值存在于Masking-Key中。

- Payload length: 7 bits, 7+16 bits, or 7+64 bits
`Payload data` 长度

- Masking-Key: 0 or 4 bytes
发送的数据与同一帧中的掩码进行过了运算，用于解码 `Payload data`
运算公式为：payload[i] = origin_data[i] ^ masking_key[i%4] 。

- Payload data: (x+y) bytes
`Payload data` 包括 `Extension data` 和 `Application data`。

### 关闭连接
为了使用一个状态码关闭websocket，一端必须发送一个关闭的控制帧，当两端都发送了关闭数据帧时，双方都要关闭所有的连接资源。控制帧为一个“状态码”和一个“原因说明”，当关闭之后，双方处于CLOSED状态。