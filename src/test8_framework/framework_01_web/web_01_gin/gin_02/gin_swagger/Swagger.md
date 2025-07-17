

来源：https://www.bookstack.cn/read/golang_development_notes/zh-1.11.md

#   


在使用 **Swagger（swaggo）** 编写接口注释时，不同的 HTTP 方法（POST、GET、DELETE、PUT）在写法上只有少量差异，主要体现在：

| HTTP 方法 | 常用于  | 请求参数写法        | @Router 结尾的方式 | 是否支持 @Param body |
| ------- | ---- | ------------- | ------------- | ---------------- |
| GET     | 查询数据 | query/path    | `[get]`       | ❌ 不支持 body 参数    |
| POST    | 新增资源 | body/formData | `[post]`      | ✅ 支持 body 参数     |
| PUT     | 更新资源 | body/path     | `[put]`       | ✅ 支持 body 参数     |
| DELETE  | 删除资源 | path/query    | `[delete]`    | ❌ 通常不支持 body     |

---

## ✅ 各方法示例

### 🔵 1. GET 示例（不支持 body）

```go
// @Summary 获取用户详情
// @Description 通过用户ID获取用户详情
// @Tags 用户模块
// @Param id query int true "用户ID"
// @Success 200 {object} UserResponse
// @Router /user [get]
```

### 🔵 2. POST 示例（提交 JSON）

```go
// @Summary 创建用户
// @Description 提交用户信息创建用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "用户信息"
// @Success 200 {object} Response
// @Router /user [post]
```

### 🔵 3. PUT 示例（更新数据）

```go
// @Summary 更新用户信息
// @Description 通过用户ID更新信息
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body UpdateUserRequest true "用户信息"
// @Success 200 {object} Response
// @Router /user/{id} [put]
```

### 🔵 4. DELETE 示例（通过路径参数）

```go
// @Summary 删除用户
// @Description 通过用户ID删除用户
// @Tags 用户模块
// @Param id path int true "用户ID"
// @Success 200 {object} Response
// @Router /user/{id} [delete]
```

---

## ✅ 参数类型说明

| 注释                    | 含义                |
| --------------------- | ----------------- |
| `@Param xxx query`    | URL 查询参数 (?xxx=1) |
| `@Param xxx path`     | 路径参数（/user/{id}）  |
| `@Param xxx body`     | 请求体 JSON          |
| `@Param xxx formData` | 表单上传参数（如文件）       |

---

## 🧠 小建议

* 🔒 建议所有需要鉴权的接口都加 `@Security BearerAuth`
* 🧪 添加 `@Failure 400/401/500` 可以完善失败响应说明
* 🔄 PUT 一般用于更新，PATCH 用于部分更新，Swagger 写法类似

---

如果你贴出某个接口的代码，我也可以帮你生成完整的注释模板（含所有方法写法）。是否需要？
