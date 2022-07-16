---
title: Go HTTP 标准库 多线程下载
date: 2022-01-30 19:12:55
top_img: https://pic1.zhimg.com/80/v2-e4870b3695ef144db682c464ba21df02_720w.jpeg
cover: https://pic1.zhimg.com/80/v2-e4870b3695ef144db682c464ba21df02_720w.jpeg
tags: 
    - Golang
    - HTTP
---

本文以 `https://i.pximg.net/img-master/img/2022/01/30/00/50/14/95863886_p0_master1200.jpg` 的下载作为例子

```Go
func GetHeader(image string) (http.Header, error) {
	// P站特殊客户端
	client := &http.Client{
		// 解决中国大陆无法访问的问题
		Transport: &http.Transport{
			// 更改 dns
			Dial: func(network, addr string) (net.Conn, error) {
				return net.Dial("tcp", "210.140.92.142:443")
			},
			// 隐藏 sni 标志
			TLSClientConfig: &tls.Config{
				ServerName:         "-",
				InsecureSkipVerify: true,
			},
		},
	}
	// 请求 Header
	req, err := http.NewRequest("HEAD", image, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Host", "i.pximg.net")
	req.Header.Set("Referer", "https://www.pixiv.net/")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Header, nil
}
```

执行 `GetHeader("https://i.pximg.net/img-master/img/2022/01/30/00/50/14/95863886_p0_master1200.jpg")` ，我们将返回的 `Header` 打印得到

```HTTP HEADER
Cache-Control: max-age=31536000
Expires: Sun, 29 Jan 2023 15:50:15 GMT
Last-Modified: Sat, 29 Jan 2022 15:50:14 GMT
X-Content-Type-Options: nosniff
Server: nginx
Content-Type: image/jpeg
Age: 237119
Via: http/1.1 f001 (second)
Accept-Ranges: bytes
Date: Tue, 01 Feb 2022 09:45:03 GMT
Content-Length: 1093164
```

其中， `Content-Length: 1093164` 为该资源的大小， `Accept-Ranges: bytes` 表示该资源是可断点续传的，因而我们可以实现并发下载资源。

```Go
func DownSlice(image string, start, end int) ([]byte, error) {
	// P站特殊客户端
	client := &http.Client{
		// 解决中国大陆无法访问的问题
		Transport: &http.Transport{
			// 更改 dns
			Dial: func(network, addr string) (net.Conn, error) {
				return net.Dial("tcp", "210.140.92.142:443")
			},
			// 隐藏 sni 标志
			TLSClientConfig: &tls.Config{
				ServerName:         "-",
				InsecureSkipVerify: true,
			},
		},
	}
	// 请求 资源
	req, err := http.NewRequest("GET", image, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Host", "i.pximg.net")
	req.Header.Set("Referer", "https://www.pixiv.net/")
	// 设置下载范围
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end-1))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return b, nil
}
```

执行 `DownSlice("https://i.pximg.net/img-master/img/2022/01/30/00/50/14/95863886_p0_master1200.jpg", 0, 1024)` ，即可下载资源的 0-1024 的部分，知道原理后我们利用协程即可对资源进行并发下载。