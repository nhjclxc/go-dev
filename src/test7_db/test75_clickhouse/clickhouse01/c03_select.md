好，这里我直接给你一套 **围绕 `my_first_table2` 的“实战级查询示例”**，从**最基础**到**ClickHouse 风格的进阶查询**，每条都是真实工程里会用到的。

表结构回顾一下：

```sql
CREATE TABLE my_first_table2
(
    user_id UInt32,
    message String,
    timestamp DateTime,
    metric Float32
)
ENGINE = MergeTree
PRIMARY KEY (user_id, timestamp);
```

---

## 一、基础查询（你 100% 会用到）

### 1️⃣ 查询最近 10 条数据

```sql
SELECT *
FROM my_first_table2
ORDER BY timestamp DESC
LIMIT 10;
```

📌 常用于：

* 校验数据是否写入成功
* 手工排查问题

---

### 2️⃣ 查询指定用户的数据

```sql
SELECT user_id, message, timestamp, metric
FROM my_first_table2
WHERE user_id = 1001
ORDER BY timestamp;
```

📌 **PRIMARY KEY 生效场景**（`user_id + timestamp`）

---

### 3️⃣ 时间范围查询（ClickHouse 最常见）

```sql
SELECT *
FROM my_first_table2
WHERE timestamp >= now() - INTERVAL 1 HOUR
ORDER BY timestamp;
```

📌 日志 / 指标系统必备

---

## 二、聚合分析（ClickHouse 的强项）

### 4️⃣ 按用户统计指标

```sql
SELECT
    user_id,
    count()           AS cnt,
    avg(metric)       AS avg_metric,
    max(metric)       AS max_metric
FROM my_first_table2
GROUP BY user_id
ORDER BY cnt DESC;
```

📌 秒杀 MySQL 的地方就在这里

---

### 5️⃣ 按时间窗口聚合（准实时）

```sql
SELECT
    toStartOfMinute(timestamp) AS minute,
    count()                    AS cnt,
    avg(metric)                AS avg_metric
FROM my_first_table2
WHERE timestamp >= now() - INTERVAL 30 MINUTE
GROUP BY minute
ORDER BY minute;
```

📌 常用于：

* QPS
* 指标趋势
* 时间序列分析

---

## 三、ClickHouse 风格查询（很重要）

### 6️⃣ Top N 用户（高频用户）

```sql
SELECT
    user_id,
    count() AS cnt
FROM my_first_table2
GROUP BY user_id
ORDER BY cnt DESC
LIMIT 5;
```

---

### 7️⃣ 条件过滤 + 聚合

```sql
SELECT
    user_id,
    avg(metric) AS avg_metric
FROM my_first_table2
WHERE
    timestamp >= now() - INTERVAL 1 DAY
    AND metric > 50
GROUP BY user_id
ORDER BY avg_metric DESC;
```

---

## 四、利用 PRIMARY KEY 的查询方式（性能关键）

> ClickHouse 的 PRIMARY KEY = 排序键
> 查询条件必须 **从左到右命中**

### 8️⃣ 命中完整主键（最快）

```sql
SELECT *
FROM my_first_table2
WHERE user_id = 42
  AND timestamp BETWEEN
      toDateTime('2026-02-01 00:00:00')
      AND toDateTime('2026-02-01 01:00:00');
```

🚀 **最快路径**

---

### 9️⃣ 只用 user_id（仍然高效）

```sql
SELECT *
FROM my_first_table2
WHERE user_id = 42
ORDER BY timestamp DESC
LIMIT 100;
```

---

### ⚠️ 10️⃣ 只用 timestamp（会扫更多数据）

```sql
SELECT *
FROM my_first_table2
WHERE timestamp >= now() - INTERVAL 1 HOUR;
```

📌 能跑，但不会利用 `user_id` 排序优势

---

## 五、常见运维 / 排查查询

### 🔍 11️⃣ 表行数估算（超快）

```sql
SELECT count()
FROM my_first_table2;
```

---

### 🔍 12️⃣ 查看表大小

```sql
SELECT
    table,
    sum(rows) AS rows,
    formatReadableSize(sum(bytes_on_disk)) AS size
FROM system.parts
WHERE table = 'my_first_table2'
GROUP BY table;
```

---

### 🔍 13️⃣ 查看分区（如果后续加了 PARTITION）

```sql
SELECT
    partition,
    count() AS parts
FROM system.parts
WHERE table = 'my_first_table2'
GROUP BY partition;
```

---

## 六、Go 里查询示例（配合你现有 DB）

```go
rows, err := ch.DB.QueryContext(ctx,
	`SELECT user_id, message, timestamp, metric
	 FROM my_first_table2
	 WHERE user_id = ?
	 ORDER BY timestamp DESC
	 LIMIT 10`,
	42,
)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
	var (
		userID    uint32
		message   string
		ts        time.Time
		metric    float32
	)
	if err := rows.Scan(&userID, &message, &ts, &metric); err != nil {
		log.Fatal(err)
	}
	fmt.Println(userID, message, ts, metric)
}
```

---

## 七、一句话总结（你现在该怎么用）

* **写入**：`PrepareBatch`（你已经做对了）
* **查询**：`database/sql` + 聚合思维
* **时间查询**：尽量配合 `timestamp`
* **高性能查询**：命中 `PRIMARY KEY` 前缀

---

如果你愿意，下一步我可以帮你：

* 🔥 **基于这个表设计一个“指标表”版本（含分区 / TTL）**
* 🔥 **把常用查询封装成 DAO**
* 🔥 **分析哪些查询会慢 & 怎么改 ORDER BY**

你更想往 **查询优化** 还是 **表设计进阶** 继续？
