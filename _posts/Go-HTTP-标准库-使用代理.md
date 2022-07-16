---
title: Go HTTP 标准库 使用代理
date: 2022-02-03 10:52:19
top_img: https://pic1.zhimg.com/80/v2-f96d22e058b46469ad8443bbe518460c_720w.jpeg
cover: https://pic1.zhimg.com/80/v2-f96d22e058b46469ad8443bbe518460c_720w.jpeg
tags: 
    - Golang
    - HTTP
    - Proxy
---

Go HTTP 标准库不走系统代理，因此我们通过 `Fiddler` 抓包的时候，是抓不到标准库发送的请求的。因此我们需要设置其代理

```Go
func fiddler() error {
	proxy, err := url.Parse("http://127.0.0.1:8866")
	if err != nil {
		return err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
```

执行 `fiddler()` 我们可以在 `Fiddler` 里面抓到访问 `https://www.baidu.com` 的请求。其中 `8866` 为 `Fiddler` 的监听端口，可以通过 `Fiddler` 设置。


顺带一提，我们可以通过 `import "golang.org/x/sys/windows/registry"` 获取系统代理

```Go
key, err := registry.OpenKey(
	registry.CURRENT_USER,
	`SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`,
	registry.QUERY_VALUE)
if err != nil {
	panic(err)
}
val, _, err := key.GetStringValue("ProxyServer")
if err != nil {
	panic(err)
}
```