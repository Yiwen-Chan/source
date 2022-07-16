---
title: luasocket 安装
date: 2021-08-16 17:04:22
top_img: https://pic2.zhimg.com/80/v2-744c9eca037b02be297b3a37df9b2309_720w.jpeg
cover: https://pic2.zhimg.com/80/v2-744c9eca037b02be297b3a37df9b2309_720w.jpeg
tags: 
    - Lua
    - LuaRocks
    - luasocket
---

### 环境
- Lua(version 3.3.6)
- LuaRocks(version 3.7.0)

本篇博文主要详细讲述 luasocket 安装过程中踩到过的坑。

### 安装流程
`CMD` 内输入 `luarocks install luasocket`

以下为安装过程中遇到的几个坑，记录一下

#### 网络波动
```
Warning: Failed searching manifest: Failed downloading https://luarocks.org/manifest-5.3 - host or service not provided, or not known
Warning: Failed searching manifest: Failed downloading https://raw.githubusercontent.com/rocks-moonscript-org/moonrocks-mirror/master/manifest-5.3 - host or service not provided, or not known
Warning: Failed searching manifest: Failed downloading https://luafr.org/luarocks/manifest-5.3 - host or service not provided, or not known
Warning: Failed searching manifest: Failed downloading http://luarocks.logiceditor.com/rocks/manifest-5.3 - host or service not provided, or not known

Error: No results matching query were found for Lua 5.3.
To check if it is available for other Lua versions, use --check-lua-versions.
```

解决办法：
多试几次就好了，实在不行手动下载 `luasocket-3.0rc1-1.src.rock` ，并在 `luasocket-3.0rc1-1.src.rock` 目录下使用 `CMD` 命令 `luarocks install luasocket-3.0rc1-1.src.rock`。

#### 提示 mingw32-gcc.exe
```
'mingw32-gcc' 不是内部或外部命令，也不是可运行的程序或批处理文件。
```

解决办法：
在 GCC 环境中 `bin` 文件夹中，找到 `x86_64-w64-mingw32-gcc-9.2.0.exe`，复制一份并改名成 `mingw32-gcc.exe`。

#### C 编译不通过
```
src/luasocket.c: In function 'global_skip':
src/luasocket.c:67:18: warning: implicit declaration of function 'luaL_checkint'; did you mean 'luaL_checkany'? [-Wimplicit-function-declaration]
   67 |     int amount = luaL_checkint(L, 1);
      |                  ^~~~~~~~~~~~~
      |                  luaL_checkany
```

解决办法：
1. 在 `luasocket-3.0rc1-1.src.rock` 目录下使用 `CMD` 命令 `luarocks unpack luasocket-3.0rc1-1.src.rock` 
```
|-- luasocket-3.0rc1-1
    |-- luasocket-3.0rc1-1
        |-- src
        |-- doc
        |-- luasocket-3.0rc1-1.rockspec
    |-- luasocket-3.0rc1-1.rockspec
    |-- v3.0-rc1.zip
```
2. 增加 Lua 版本 3.3.6 兼容
用文本编辑器打开里层的 `luasocket-3.0rc1-1.rockspec` 文件，找到 `mingw32` 并修改成以下
```
mingw32 = {
    "LUASOCKET_DEBUG",
    "LUASOCKET_INET_PTON",
    "LUA_COMPAT_5_2",
    "WINVER=0x0501",
    "LUASOCKET_API=__declspec(dllexport)",
    "MIME_API=__declspec(dllexport)"
}
```
其中新增的 `LUA_COMPAT_5_2` 为 Lua 5.3.6 向下兼容
3. 安装
在里层 `luasocket-3.0rc1-1` 目录下使用 `CMD` 命令  `luarocks make luasocket-3.0rc1-1.rockspec`

#### GCC 炸了
```
C:/TDM-GCC-64/bin/../lib/gcc/x86_64-w64-mingw32/9.2.0/../../../../x86_64-w64-mingw32/bin/ld.exe: skipping incompatible c:/windows/system32/ws2_32.dll when searching for -lws2_32
C:/TDM-GCC-64/bin/../lib/gcc/x86_64-w64-mingw32/9.2.0/../../../../x86_64-w64-mingw32/bin/ld.exe: skipping incompatible c:/windows/system32/msvcrt.dll when searching for -lmsvcrt
C:/TDM-GCC-64/bin/../lib/gcc/x86_64-w64-mingw32/9.2.0/../../../../x86_64-w64-mingw32/bin/ld.exe: skipping incompatible c:/windows/system32/advapi32.dll when searching for -ladvapi32
C:/TDM-GCC-64/bin/../lib/gcc/x86_64-w64-mingw32/9.2.0/../../../../x86_64-w64-mingw32/bin/ld.exe: skipping incompatible c:/windows/system32/shell32.dll when searching for -lshell32
C:/TDM-GCC-64/bin/../lib/gcc/x86_64-w64-mingw32/9.2.0/../../../../x86_64-w64-mingw32/bin/ld.exe: skipping incompatible c:/windows/system32/user32.dll when searching for -luser32
C:/TDM-GCC-64/bin/../lib/gcc/x86_64-w64-mingw32/9.2.0/../../../../x86_64-w64-mingw32/bin/ld.exe: skipping incompatible c:/windows/system32/kernel32.dll when searching for -lkernel32
C:/TDM-GCC-64/bin/../lib/gcc/x86_64-w64-mingw32/9.2.0/../../../../x86_64-w64-mingw32/bin/ld.exe: skipping incompatible c:/windows/system32/msvcrt.dll when searching for -lmsvcrt
```

解决办法：
低版本的 GCC 无法完整编译 luasocket ，更换高版本的 GCC ，博主这里用的是 `x86_64-w64-mingw32-gcc-9.2.0.exe`

#### Lua 找不到安装的依赖
```
Lua 5.3.6  Copyright (C) 1994-2020 Lua.org, PUC-Rio
> require("socket")
stdin:1: module 'socket' not found:
        no field package.preload['socket']
        no file 'E:\kanri program\Lua\bin\lua\socket.lua'
        no file 'E:\kanri program\Lua\bin\lua\socket\init.lua'
        no file 'E:\kanri program\Lua\bin\socket.lua'
        no file 'E:\kanri program\Lua\bin\socket\init.lua'
        no file 'E:\kanri program\Lua\bin\..\share\lua\5.3\socket.lua'
        no file 'E:\kanri program\Lua\bin\..\share\lua\5.3\socket\init.lua'
        no file '.\socket.lua'
        no file '.\socket\init.lua'
        no file 'C:\Program Files (x86)\Lua\5.3.6\lua\socket.luac'
        no file 'E:\kanri program\Lua\bin\socket.dll'
        no file 'E:\kanri program\Lua\bin\..\lib\lua\5.3\socket.dll'
        no file 'E:\kanri program\Lua\bin\loadall.dll'
        no file '.\socket.dll'
stack traceback:
        [C]: in function 'require'
        stdin:1: in main chunk
        [C]: in ?
```

解决办法：
`CMD` 内输入 `luarocks path --bin` ，将路径添加到环境变量中
注意 `CMD` 内使用 `SET` 命令配置的环境变量只适用于当前 `CMD` 窗口