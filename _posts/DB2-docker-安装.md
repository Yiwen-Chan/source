<!--
 * @Author: Kanri
 * @Date: 2022-01-11 16:28:57
 * @LastEditors: Kanri
 * @LastEditTime: 2022-07-16 19:47:47
 * @Description: 
-->
---
title: DB2 docker 安装
date: 2021-09-22 10:05:34
top_img: https://pic3.zhimg.com/80/v2-a2de900602d31d2fa5f3c4792b7ceaab_720w
cover: https://pic3.zhimg.com/80/v2-a2de900602d31d2fa5f3c4792b7ceaab_720w
tags: 
    - DB2
    - docker
    - database
---

### 官方文档
[Installing the Db2 Community Edition Docker image on Linux systems](https://www.ibm.com/support/producthub/db2/docs/content/SSEPGG_11.5.0/com.ibm.db2.luw.db2u_openshift.doc/doc/t_install_db2CE_linux_img.html)

### 前期准备
- 需要安装并掌握 docker 的使用方法

### 安装流程
#### 拉取 docker 镜像
1. 拉取官方最新版本的 DB2 docker 镜像
`docker pull ibmcom/db2`
2. 查看已经拉取的 docker 镜像
`docker image ls`
可以看到返回结果
`ibmcom/db2            latest    a6a5ee354fb1   2 months ago   2.95GB`
#### 进入 docker 容器
1. 进入Docker容器：
`docker run -d -p 50000:50000 --name db2 --privileged=true -e DB2INST1_PASSWORD=123456 -e DBNAME=testdb -e LICENSE=accept ibmcom/db2`
* -d: 表示在后台启动容器
* -p 50000:50000: 容器内部的 50000 端口映射到主机的 50000 端口
* --name db2: 将容器命名 db2
* --privileged=true: 使得容器内的 root 拥有真正的 root 权限
* -e DB2INST1_PASSWORD=123456：设置内置实例用户 db2inst1 的密码为 123456
* -e DBNAME=testdb：容器启动时自动创建一个名为 testdb 的数据库，如果不指定该参数则不创建数据库
* -e LICENSE=accept：接受协议
2. 查看已有的 docker 容器
`docker ps`
可以看到返回结果
`d16a04516597   ibmcom/db2            "/var/db2_setup/lib/…"   6 days ago   Up 2 hours   22/tcp, 55000/tcp, 60006-60007/tcp, 0.0.0.0:50000->50000/tcp   db2`
3. 启动 db2 容器
`docker start db2`
4. 进入 db2 容器
`docker attach db2`
#### 信息查看
1. 切换到 db2iadm1 组的用户 db2inst1 
`su db2inst1`
2. 查看 DB2 版本
`db2licm -l`
可以看到返回结果
`Product name:                     "DB2 Community Edition"`
`License type:                     "Community"`
`Expiry date:                      "Permanent"`
`Product identifier:               "db2dec"`
`Version information:              "11.5"`
`Max amount of memory (GB):        "16"`
`Max number of cores:              "4"`
`Max amount of table space (GB):   "100"`
3. 查看实例信息
`db2level`
可以看到返回结果
`DB21085I  This instance or install (instance name, where applicable:
"testdb") uses "64" bits and DB2 code release "SQL11056" with level
identifier "0607010F".
Informational tokens are "DB2 v11.5.6.0", "s2106111000", "DYN2106111000AMD64",
and Fix Pack "0".
Product is installed at "/opt/ibm/db2/V11.5".`
4. 查看该系统下的实例列表
`db2ilist`
5. 查看该用户下的实例列表
`db2 list db directory`
#### 启动数据库
1. 用实例对应的用户，启动数据库
`db2start`
可以看到返回结果
`SQL1063N  DB2START processing was successful.`
2. 使用本地连接数据库
`db2 connect to testdb user db2inst1 using 123456`
可以看到返回结果
`Database Connection Information`
`Database server        = DB2/LINUXX8664 11.5.6.0`
`SQL authorization ID   = DB2INST1`
`Local database alias   = TESTDB`
