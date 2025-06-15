

Beanstalks


go get github.com/zeromicro/go-queue@latest

一个 Redis 实例来支持延迟队列，因为 delayqueue 底层依赖 Redis 的 zset。


# 安装 beanstalkd
https://blog.csdn.net/perfectzq/article/details/74905808

https://github.com/kr/beanstalkd
https://github.com/beanstalkd/beanstalkd

1. 下载源码：
    `wget https://codeload.github.com/beanstalkd/beanstalkd/tar.gz/refs/tags/v1.4.6`

2. 解压：
    `tar xzf beanstalkd-1.4.6.tar.gz`

3. 安装

```shell
cd beanstalkd-1.4.6
./configure
make
make instal
```

4. 