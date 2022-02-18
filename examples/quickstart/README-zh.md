# Polaris Go

[English Document](./README.md)

## 目录说明

> provider

北极星被调端样例包含了最简单的被调端基本操作

> consumer

北极星主调端样例包含了最简单的客户端基本操作


## 如何构建

> provider

直接依赖go mod进行构建

- linux/mac构建命令
```
cd ./provider
go build -o provider
```
- windows构建命令
```
cd ./provider
go build -o provider.exe
```

> consumer

- linux/mac构建命令
```
cd ./consumer
go build -o consumer
```
- windows构建命令
```
cd ./consumer
go build -o consumer.exe
```

## 如何使用

### 创建服务

预先通过北极星控制台创建对应的服务，如果是通过本地一键安装包的方式安装，直接在浏览器通过127.0.0.1:8091打开控制台

### 修改配置

指定北极星服务端地址，需编辑polaris.yaml文件，填入服务端地址

```
global:
  serverConnector:
    addresses:
    - 127.0.0.1:8091
```

### 执行程序

直接执行生成的可执行程序

> provider

- linux/mac运行命令
```
./provider
```

- windows运行命令
```
./provider.exe
```

> consumer


- linux/mac运行命令
```
./consumer
```

- windows运行命令
```
./consumer.exe
```

### 验证

```
curl http://127.0.0.1:18080/echo

Hello, I'm EchoServerGolang Provider
```
