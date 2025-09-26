在 macOS 上通过 **Homebrew 安装 Neo4j** 非常方便，下面给你完整步骤：

---

## **1. 更新 Homebrew**

```bash
brew update
```

---

## **2. 安装 Neo4j**

```bash
brew install neo4j
```

> 如果你想安装最新的社区版（Community Edition），直接 `brew install neo4j` 就可以。
> 如果要安装企业版，需要从官网下载安装包。

---

## **3. 启动 Neo4j 服务**

### **启动一次性服务（当前终端有效）**

```bash
neo4j console
```

* 在终端可以看到日志和控制台信息
* 停止服务按 `Ctrl+C`

### **启动后台服务（长期运行）**

```bash
brew services start neo4j
```

* 会自动作为后台服务启动
* 查看状态：

```bash
brew services list
```

* 停止服务：

```bash
brew services stop neo4j
```

---

## **4. 设置密码（首次使用）**

Neo4j 默认账号：

```
用户名: neo4j
默认密码: neo4j
```

首次登录需要修改密码，可以在 **浏览器访问 Neo4j Desktop** 或用命令行：

```bash
neo4j start
# 然后访问浏览器 http://localhost:7474/ 修改密码
```

> 默认 Web UI 地址：[http://localhost:7474](http://localhost:7474)

---

## **5. 常用命令**

```bash
# 启动 Neo4j
neo4j start

# 停止 Neo4j
neo4j stop

# 查看状态
neo4j status

# 打开 Neo4j 控制台
neo4j console
```

---

如果你需要，我可以帮你写一条 **Homebrew 安装 Neo4j 并设置为后台服务 + 自动打开浏览器** 的一键命令，让安装和启动更省心。

你希望我直接给这个一键命令吗？




你想在 **Go** 中操作 **Neo4j**，我可以给你一个完整的指导，包括连接、执行查询、插入数据和遍历图结构的示例。我们一步步来看。

---

## 1️⃣ 安装 Neo4j Go 驱动

Neo4j 官方提供了 Go 驱动：[https://github.com/neo4j/neo4j-go-driver](https://github.com/neo4j/neo4j-go-driver)

```bash
go get github.com/neo4j/neo4j-go-driver/v5/neo4j
```

---

## 2️⃣ 建立连接

```go
package main

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	// 连接配置
	uri := "neo4j://localhost:7687"
	username := "neo4j"
	password := "password"

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close(context.Background())

	// 测试连接
	session := driver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())

	result, err := session.Run(context.Background(), "RETURN 'Hello, Neo4j!' AS message", nil)
	if err != nil {
		log.Fatal(err)
	}

	for result.Next(context.Background()) {
		fmt.Println(result.Record().Values[0]) // 输出: Hello, Neo4j!
	}
}
```

⚠️ 注意：

* Neo4j 5.x 使用 `neo4j-go-driver/v5`
* 连接 URI 可以是 `neo4j://` 或 `bolt://`
* Neo4j 默认用户名 `neo4j`，密码在初始化时设置

---

## 3️⃣ 创建节点

```go
_, err := session.Run(context.Background(),
	"CREATE (p:Person {name: $name, age: $age})",
	map[string]any{
		"name": "Alice",
		"age":  30,
	})
if err != nil {
	log.Fatal(err)
}
```

这里用 `$name`、`$age` 做参数，防止注入。

---

## 4️⃣ 查询节点

```go
result, err := session.Run(context.Background(),
	"MATCH (p:Person) WHERE p.name = $name RETURN p.name, p.age",
	map[string]any{"name": "Alice"},
)
if err != nil {
	log.Fatal(err)
}

for result.Next(context.Background()) {
	record := result.Record()
	name, _ := record.Get("p.name")
	age, _ := record.Get("p.age")
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

---

## 5️⃣ 创建关系

```go
_, err := session.Run(context.Background(),
	`MATCH (a:Person {name: $name1}), (b:Person {name: $name2})
	CREATE (a)-[:FRIEND]->(b)`,
	map[string]any{"name1": "Alice", "name2": "Bob"},
)
```

---

## 6️⃣ 遍历图

```go
result, err := session.Run(context.Background(),
	"MATCH (a:Person)-[:FRIEND]->(b:Person) RETURN a.name, b.name",
	nil,
)

for result.Next(context.Background()) {
	record := result.Record()
	fmt.Printf("%s is friend with %s\n", record.Values[0], record.Values[1])
}
```

---

## 🔹 Tips

1. **事务**：可以使用 `session.ExecuteWrite` / `session.ExecuteRead` 来封装事务。
2. **性能**：批量插入可以用 `UNWIND`。
3. **调试**：Neo4j Browser 或 Neo4j Desktop 是最好的可视化工具。
4. **安全**：尽量不要在 URI 中写明密码，使用环境变量或配置文件。

---

如果你需要，我可以帮你写一个 **完整的 Go 小程序**，实现 **添加节点、查询节点、创建关系、遍历图** 的一整套操作流程。

你希望我帮你写吗？
