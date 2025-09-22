
# Golang爬虫学习

## [colly](https://github.com/gocolly/colly)

[docs](https://go-colly.org/docs/)


获取 `go get -u github.com/gocolly/colly/v2`

[快速开始](https://juejin.cn/post/7231130096337207353)

[教程案例](https://darjun.github.io/2021/06/30/godailylib/colly/)




cdn
http -> 302



## [goquery](https://github.com/PuerkitoBio/goquery)
https://github.com/PuerkitoBio/goquery


## [go-crawler](https://github.com/lizongying/go-crawler)
https://github.com/lizongying/go-crawler



# Golang爬虫学习

Go 里写爬虫，常用的库可以分成几类：**HTTP 请求库、HTML 解析库、完整爬虫框架、动态页面处理工具、辅助库（限速/并发/存储）**。我给你按用途列一个清单：

---

## 🔗 1. HTTP 请求相关

* **标准库 `net/http`**
  Go 自带，功能完整，支持连接池、超时、代理配置。性能好，适合自定义。

* **[resty](https://github.com/go-resty/resty)**
  一个高级 HTTP 客户端，支持链式调用、自动重试、代理、请求日志。写起来比 `net/http` 简洁。

---

## 📝 2. HTML/DOM 解析

* **[goquery](https://github.com/PuerkitoBio/goquery)**
  最常用的解析库，提供类似 jQuery 的选择器 API，比如 `doc.Find("div.title").Text()`。

* **[htmlquery](https://github.com/antchfx/htmlquery)**
  支持 XPath 语法解析 HTML，适合熟悉 XPath 的人。

* **[antchfx/xpath](https://github.com/antchfx/xpath)**
  纯 XPath 解析引擎，可以和 `htmlquery` / `xmlquery` 搭配。

---

## 🕷️ 3. 爬虫框架

* **[Colly](https://github.com/gocolly/colly)**
  Go 生态里最流行的爬虫框架，支持：

    * 自动去重
    * 并发抓取
    * URL 限制（域名/深度）
    * 请求回调（OnRequest, OnResponse, OnHTML）
    * 插件扩展（代理池、存储到 Redis）

* **[Ferret](https://github.com/MontFerret/ferret)**
  类似浏览器的爬取引擎，支持 JavaScript 渲染页面，可以用类 SQL 的查询语言写抓取逻辑。

* **[grab](https://github.com/cavaliercoder/grab)**
  更像是“下载器”库，支持断点续传、并发下载，适合抓文件/大资源。

---

## 🖥️ 4. 动态页面（JS 渲染）

* **[chromedp](https://github.com/chromedp/chromedp)**
  直接操作 Chrome DevTools 协议，轻量、无头浏览器。

* **[rod](https://github.com/go-rod/rod)**
  比 chromedp API 更易用，社区活跃，支持自动化、反检测处理。

* **[playwright-go](https://github.com/playwright-community/playwright-go)**
  Go 绑定的 Playwright，可以跨浏览器（Chromium、Firefox、WebKit）。

---

## ⚡ 5. 并发、限速、任务管理

* **[golang.org/x/time/rate](https://pkg.go.dev/golang.org/x/time/rate)**
  官方限速器，实现令牌桶，常用于控制抓取频率。

* **[workqueue](https://github.com/kubernetes/client-go/tree/master/util/workqueue)**
  Kubernetes 用的任务队列库，可以当爬虫调度器。

* **[ants](https://github.com/panjf2000/ants)**
  高性能 goroutine 池，适合大规模并发爬取时节省资源。

---

## 💾 6. 存储与去重

* **[go-redis](https://github.com/redis/go-redis)**
  常用来做 URL 队列和去重（set/bloom filter）。

* **[tidwall/buntdb](https://github.com/tidwall/buntdb)**
  嵌入式内存+磁盘 KV 库，可以当本地去重/存储引擎。

---

✅ **总结：**

* 写轻量级爬虫 → `net/http` + `goquery`。
* 快速搭建爬虫框架 → `Colly`。
* 动态页面 → `chromedp` / `rod`。
* 控制并发/限速 → `rate` + `ants`。
* 分布式/去重 → `go-redis`。

---

要不要我给你整理一个 **“Go 爬虫常用库组合模板”**，比如：`Colly + goquery + rate + Redis`，直接能跑的示例项目？
