---
title: LRU
date: 2022-02-12 17:54:51
top_img: https://pica.zhimg.com/80/v2-a453bf0124677aa5f3dfbd471935a3ac_720w.png
cover: https://pica.zhimg.com/80/v2-a453bf0124677aa5f3dfbd471935a3ac_720w.png
tags: 
    - Golang
    - LRU
---

这是一个用 双向链表与哈希表实现的 `LRU` (Least Recently Used) 缓存淘汰算法

```go
package main

import (
	"container/list"
	"fmt"
)

type Data struct {
	Key string
	Val interface{}
}

type LRU struct {
	Cap  int
	List *list.List
	Hash map[string]*list.Element
}

func NewLRU(cap int) *LRU {
	return &LRU{Cap: cap, List: list.New(), Hash: make(map[string]*list.Element)}
}

func (lru *LRU) Put(key string, val interface{}) {
	if e, ok := lru.Hash[key]; ok {
		lru.List.Remove(e)
	}
	e := lru.List.PushBack(&Data{Key: key, Val: val})
	lru.Hash[key] = e
	if lru.List.Len() > lru.Cap {
		f := lru.List.Front()
		lru.List.Remove(f)
		delete(lru.Hash, f.Value.(*Data).Key)
	}
}

func (lru *LRU) Get(key string) (val interface{}) {
	if e, ok := lru.Hash[key]; ok {
        lru.List.MoveToBack(e)
		return e.Value.(*Data).Val
	}
	return nil
}
```