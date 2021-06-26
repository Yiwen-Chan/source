---
title: Go http 标准库的 SNI 修改
date: 2021-06-21 07:20:07
top_img: https://s3.jpg.cm/2021/06/21/IRV4hz.jpg
cover: https://s3.jpg.cm/2021/06/21/IRV4hz.jpg
tags: 
    - Golang
    - HTTP
    - SNI
---

### SNI 阻断
SNI (Server Name Indication)，为 `TLS` 连接中客户端发起的第一个握手包 `Client Hello` 中的即将访问的域名信息数据的字段。该字段是为了解决某些服务器同时含有多个域名站点，在加密传输之前，它需要知道客户端访问的是哪个域名。因此某些防火墙能对报文中的 `SNI` 识别并进行阻断。

### SNI 修改
本文以 `P站 (https://pixiv.net)` 为例，示例如何修改 `Go http 标准库` 请求中的 `SNI` 字段数据。造成 P站 无法访问的原因有两方面，一方面是 `DNS` 污染，本文不多加叙述，另一个则为 `SNI` 阻断。因为 P站 服务器无需 `SNI` 字段即可正常访问，因此我们可以通过修改其 `SNI` 字段的值防止 `TLS` 连接被阻断。

通过 `WireShark` 抓包能看见 `Server Name` 字段的值为 `pixiv.net`
```
Extension: server_name (len=14)
    Type: server_name (0)
    Length: 14
    Server Name Indication extension
        Server Name list length: 12
        Server Name Type: host_name (0)
        Server Name length: 9
        Server Name: pixiv.net

```

Golang http 标准库实现 SNI 隐藏
- `ServerName` 即 `SNI` 中的 `Server Name`
- `InsecureSkipVerify` 跳过证书校验

```go
client := &http.Client{
    Transport: &http.Transport{
        // 隐藏 SNI
        TLSClientConfig: &tls.Config{
            ServerName:         "-",
            InsecureSkipVerify: true,
        },
        // 更改 IP
        Dial: func(network, addr string) (net.Conn, error) {
            return net.Dial("tcp", "210.140.131.223:443")
        },
    },
}
```

修改后从抓包数据可以看见， `Server Name` 字段的值变为了 `-`
```
Extension: server_name (len=6)
    Type: server_name (0)
    Length: 6
    Server Name Indication extension
        Server Name list length: 4
        Server Name Type: host_name (0)
        Server Name length: 1
        Server Name: -
```
