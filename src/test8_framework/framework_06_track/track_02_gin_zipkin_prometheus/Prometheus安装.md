你可以通过多种方式安装 Prometheus，以下是最常见的 **三种安装方式**（推荐方式已标记 ✅）：

---

## ✅ **方式一：Docker 安装（推荐）**

这是最方便快捷的安装方式。

### 🧰 步骤

1. **创建配置文件 `prometheus.yml`**（放在当前目录）

```yaml
global:
  scrape_interval: 5s

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
```

2. **启动 Prometheus 容器**

拉镜像
```bash
docker pull swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/prom/prometheus:v3.4.1
```

启动容器
```bash
docker run -d \
  -p 9090:9090 \
  -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml \
  --name prometheus \
  swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/prom/prometheus:v3.4.1
```

3. **访问 Prometheus**
   浏览器打开：[http://localhost:9090](http://localhost:9090)


4. 重新加载配置

方法一：重启容器
```
docker restart prometheus
```


方法二：通过 HTTP API 触发热加载，Prometheus 提供了热加载 API，可以通过 curl 或 Postman 触发重新加载配置。
```
curl -X POST http://localhost:9090/-/reload
```


方法三：发送 SIGHUP 信号（推荐方式）
```
docker kill --signal=SIGHUP prometheus
或
docker kill --signal=SIGHUP 1234abcd5678
```

5. 




---

## ✅ **方式二：二进制安装（适用于本地开发调试）**

### 🧰 步骤

1. [下载 Prometheus 二进制包](https://prometheus.io/download/)

   示例（以 Linux 为例）：

   ```bash
   wget https://github.com/prometheus/prometheus/releases/download/v2.52.0/prometheus-2.52.0.linux-amd64.tar.gz
   tar -xvf prometheus-2.52.0.linux-amd64.tar.gz
   cd prometheus-2.52.0.linux-amd64
   ```

2. **运行 Prometheus**

   ```bash
   ./prometheus --config.file=prometheus.yml
   ```

3. **访问界面**
   [http://localhost:9090](http://localhost:9090)

---

## ✅ **方式三：Docker Compose（用于生产部署推荐）**

### 示例 `docker-compose.yml`

```yaml
version: '3'
services:
  prometheus:
    image: prom/prometheus
    container_name: prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
```

启动命令：

```bash
docker-compose up -d
```

---

## 🚀 验证 Prometheus 是否正常

浏览器打开：[http://localhost:9090](http://localhost:9090)

输入 `up` 查询服务状态，若返回结果为 `1` 表示运行正常。

---

## 📦 配合 Grafana 可视化（可选）

```bash
docker run -d \
  -p 3000:3000 \
  --name=grafana \
  grafana/grafana
```

Grafana UI：[http://localhost:3000](http://localhost:3000)
默认账号密码：`admin / admin`

在 Grafana 添加数据源：

* 类型：Prometheus
* 地址：[http://host.docker.internal:9090（或](http://host.docker.internal:9090（或) docker 内的 `prometheus:9090`）

---

需要我生成完整的 `docker-compose.yml` 来同时启动 `prometheus + grafana + zipkin` 吗？可以直接打包部署。
