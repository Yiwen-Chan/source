---
title: Lua LuaRocks 安装
date: 2021-08-13 10:31:07
top_img: https://s3.jpg.cm/2021/08/13/IcGJdL.jpg
cover: https://s3.jpg.cm/2021/08/13/IcGJdL.jpg
tags: 
    - Lua
    - LuaRocks
---

鉴于网上对安装过程陈述过于模糊，本篇博文主要详细讲述 LuaRocks(version 3.7.0) 搭配 Lua(version 3.3.6) ，在 Windows 环境下的安装。

### LuaRocks
LuaRocks 是 Lua 的模组管理器。能很方便在 Unix 以及 Windows 下载以及安装 Lua 模组。

### 安装流程

#### 前置环境
- `TDM-GCC` 或 `Mingw64`
如果没有需要自行安装

#### 安装 Lua(version 3.3.6)
详见 http://lua-users.org/wiki/BuildingLuaInWindowsForNewbies
- 下载源码
进入 [Lua 官网](https://www.lua.org/) 下载 Lua(version 3.3.6)，即文件 `lua-5.3.6.tar.gz` ，并解压
当前目录结构如下
```
|-- lua-5.4.3
    |-- lua-5.4.3
        |-- src
            |-- lua.c
            |-- lua.h
            |-- ...
        |-- doc
            |-- readme.html
            |-- ...
        |-- Makefile
        |-- README
```
- 创建 Windows Shell 脚本
新建文件 `build.cmd` ，放在与 Makefile 文件的父级目录同级的地方，内容如下
```
@echo off
:: ========================
:: file build.cmd
:: ========================
setlocal
:: you may change the following variable's value
:: to suit the downloaded version
set lua_version=5.3.6

set work_dir=%~dp0
:: Removes trailing backslash
:: to enhance readability in the following steps
set work_dir=%work_dir:~0,-1%
set lua_install_dir=%work_dir%\lua
set compiler_bin_dir=%work_dir%\tdm-gcc\bin
set lua_build_dir=%work_dir%\lua-%lua_version%
set path=%compiler_bin_dir%;%path%

cd /D %lua_build_dir%
mingw32-make PLAT=mingw

echo.
echo **** COMPILATION TERMINATED ****
echo.
echo **** BUILDING BINARY DISTRIBUTION ****
echo.

:: create a clean "binary" installation
mkdir %lua_install_dir%
mkdir %lua_install_dir%\doc
mkdir %lua_install_dir%\bin
mkdir %lua_install_dir%\include

copy %lua_build_dir%\doc\*.* %lua_install_dir%\doc\*.*
copy %lua_build_dir%\src\*.exe %lua_install_dir%\bin\*.*
copy %lua_build_dir%\src\*.dll %lua_install_dir%\bin\*.*
copy %lua_build_dir%\src\luaconf.h %lua_install_dir%\include\*.*
copy %lua_build_dir%\src\lua.h %lua_install_dir%\include\*.*
copy %lua_build_dir%\src\lualib.h %lua_install_dir%\include\*.*
copy %lua_build_dir%\src\lauxlib.h %lua_install_dir%\include\*.*
copy %lua_build_dir%\src\lua.hpp %lua_install_dir%\include\*.*

echo.
echo **** BINARY DISTRIBUTION BUILT ****
echo.

%lua_install_dir%\bin\lua.exe -e"print [[Hello!]];print[[Simple Lua test successful!!!]]"

echo.

pause
```
当前目录结构如下
```
|-- lua-5.4.3
    |-- lua-5.4.3
        |-- src
            |-- lua.c
            |-- lua.h
            |-- ...
        |-- doc
            |-- readme.html
            |-- ...
        |-- Makefile
        |-- README
    |-- build.cmd
```
- 编译 Lua
运行 `build.cmd` ，待编译完成后，自动生成文件夹 `Lua` ，此文件夹便是整个 Lua 环境
当前目录结构如下
```
|-- lua-5.4.3
    |-- lua-5.4.3
        |-- src
            |-- lua.c
            |-- lua.h
            |-- ...
        |-- doc
            |-- readme.html
            |-- ...
        |-- Makefile
        |-- README
    |-- lua
        |-- bin
            |-- lua.exe
            |-- ...
        |-- doc
            |-- readme.html
            |-- ...
        |-- include
            |-- lua.h
            |-- ...
    |-- build.cmd
```
- 配置 Lua 环境变量
将上述生成的 `lua` 文件夹放到合适的位置，并添加环境变量
向 `PATH` 增加一项 绝对路径 + `lua-5.4.3\lua\bin`
- 验证 Lua 环境
重新打开 `CMD` ，输入 `lua -v`，如果正确显示版本则安装成功。


#### 安装 LuaRocks(version 3.7.0)
- 下载 LuaRocks 二进制文件
[下载 LuaRocks(version 3.7.0)](http://luarocks.github.io/luarocks/releases/luarocks-3.7.0-windows-64.zip)，并解压
其他版本：http://luarocks.github.io/luarocks/releases/
- 将 `luarocks.exe` 文件复制到 `lua-5.4.3\lua\bin` 文件夹中
当前目录结构如下
```
|-- lua-5.4.3
    |-- lua-5.4.3
        |-- src
            |-- lua.c
            |-- lua.h
            |-- ...
        |-- doc
            |-- readme.html
            |-- ...
        |-- Makefile
        |-- README
    |-- lua
        |-- bin
            |-- lua.exe
            |-- luarocks.exe
            |-- ...
        |-- doc
            |-- readme.html
            |-- ...
        |-- include
            |-- lua.h
            |-- ...
    |-- build.cmd
```
- 验证 LuaRocks 环境
重新打开 `CMD` ，输入 `luarocks --version`，如果正确显示版本则安装成功。
- 新建 LuaRocks 用户配置文件夹
在 `C:\Users\kanri\AppData\Roaming` 目录下新建 `luarocks` 文件夹
其中 `kanri` 为你的用户名
- 配置 LuaRocks 使用的 Lua 环境
`CMD` 内输入 `luarocks --lua-dir "lua-5.4.3\lua"`
其中， `lua-5.4.3\lua` 为 Lua 环境的绝对路径
- 配置 LuaRocks 的 INCDIR 目录
`CMD` 内输入 `luarocks config variables.LUA_INCDIR "lua-5.4.3\lua\include"`
其中， `lua-5.4.3\lua` 为 Lua 环境的绝对路径
- 验证 LuaRocks 的配置
`CMD` 内输入 `luarocks`
下面为博主的配置：
```
Configuration:
   Lua:
      Version    : 5.3
      Interpreter: E:\lua-5.4.3\lua\bin/lua.exe (ok)
      LUA_DIR    : E:\lua-5.4.3\lua (ok)
      LUA_BINDIR : E:\lua-5.4.3\lua\bin (ok)
      LUA_INCDIR : E:\lua-5.4.3\lua/include (ok)
      LUA_LIBDIR : E:\lua-5.4.3\lua/bin (ok)

   Configuration files:
      System  : C:/Program Files/luarocks/config-5.3.lua (not found)
      User    : C:/Users/kanri/AppData/Roaming/luarocks/config-5.3.lua (ok)

   Rocks trees in use:
      C:\Users\kanri\AppData\Roaming/luarocks ("user")
```
- 配置 LuaRocks 依赖的环境变量
`CMD` 内输入 `luarocks path --bin` ，将路径添加到环境变量中
注意 `CMD` 内使用 `SET` 命令配置的环境变量只适用于当前 `CMD` 窗口