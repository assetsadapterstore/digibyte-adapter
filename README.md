# digibyte-adapter

digibyte-adapter继承了bitcoin-adapter，主要修改了如下内容：

- 重写了Symbol = "DGB"。
- 重写了addressDecoder，实现了DGB地址编码。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建DGB.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 0
# node api url, if RPC Server Type = 0, use bitcoin core full node
serverAPI = "http://127.0.0.1:8333/"
# node api url, if RPC Server Type = 1, use bitbay insight-api
; serverAPI = "http://127.0.0.1:20003/insight-api/"
# RPC Authentication Username
rpcUser = "user"
# RPC Authentication Password
rpcPassword = "password"
# Is network test?
isTestNet = true
# minimum transaction fees
minFees = "0.00001"

```

## 资料介绍

### 区块浏览器

https://digiexplorer.info/

### github

https://github.com/digibyte/digibyte/
