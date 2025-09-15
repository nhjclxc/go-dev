`go-playground/validator` 是 Go 最流行的字段验证库，它提供了 **丰富的验证标签**，几乎覆盖了大部分常见格式校验。下面我给你做一个详细汇总：

---

## 1️⃣ 基础验证（基本规则）

| 标签            | 描述                                     | 示例                                        |
| ------------- | -------------------------------------- | ----------------------------------------- |
| `required`    | 必填                                     | `Username string validate:"required"`     |
| `omitempty`   | 空值时跳过验证                                | `Age int validate:"omitempty,min=0"`      |
| `eq=xxx`      | 等于某个值                                  | `Status string validate:"eq=active"`      |
| `ne=xxx`      | 不等于某个值                                 | `Status string validate:"ne=inactive"`    |
| `gt` / `gte`  | 大于 / 大于等于                              | `Age int validate:"gte=18"`               |
| `lt` / `lte`  | 小于 / 小于等于                              | `Age int validate:"lte=65"`               |
| `min` / `max` | 最小长度 / 最大长度（string, array, slice, map） | `Password string validate:"min=6,max=20"` |
| `len`         | 精确长度（string, array, slice, map）        | `Code string validate:"len=6"`            |

---

## 2️⃣ 数字相关验证

| 标签                       | 描述      |                                        |
| ------------------------ | ------- | -------------------------------------- |
| `number`                 | 是否是数字   |                                        |
| `numeric`                | 数字字符串   |                                        |
| `hexadecimal`            | 十六进制字符串 |                                        |
| `gt`, `gte`, `lt`, `lte` | 数值比较    |                                        |
| `oneof`                  | 枚举值     | `Status string validate:"oneof=0 1 2"` |

---

## 3️⃣ 字符串格式验证

| 标签                        | 描述                           |
| ------------------------- | ---------------------------- |
| `email`                   | 邮箱格式                         |
| `url`                     | URL 格式                       |
| `uri`                     | URI                          |
| `base64`                  | Base64 编码                    |
| `uuid`                    | UUID（支持 uuid3, uuid4, uuid5） |
| `hostname`                | 主机名                          |
| `ipv4`                    | IPv4 地址                      |
| `ipv6`                    | IPv6 地址                      |
| `ip`                      | IPv4 或 IPv6                  |
| `mac`                     | MAC 地址                       |
| `contains=<substring>`    | 必须包含某字符串                     |
| `startswith=<prefix>`     | 以某前缀开头                       |
| `endswith=<suffix>`       | 以某后缀结尾                       |
| `alpha`                   | 仅字母                          |
| `alphanum`                | 字母+数字                        |
| `numeric`                 | 数字字符                         |
| `printascii`              | 可打印 ASCII 字符                 |
| `lowercase` / `uppercase` | 全部小写/大写                      |

---

## 4️⃣ 数组 / slice / map 验证

| 标签     | 描述                       |
| ------ | ------------------------ |
| `dive` | 遍历 slice/map/array 元素做验证 |
| 例子：    |                          |

```go
type User struct {
    Tags []string `validate:"dive,required"`
}
```

* `dive` 会对 `Tags` 的每个元素都应用 `required` 验证。

---

## 5️⃣ 时间相关

| 标签                   | 描述        |                                                  |
| -------------------- | --------- | ------------------------------------------------ |
| `datetime=<layout>`  | 验证时间格式    | `Birthday string validate:"datetime=2006-01-02"` |
| `gtfield=FieldName`  | 大于另一个字段   |                                                  |
| `gtefield=FieldName` | 大于等于另一个字段 |                                                  |

---

## 6️⃣ 组合使用示例

```go
type UserRequest struct {
    Username string   `json:"username" validate:"required,alphanum,min=3,max=20"`
    Email    string   `json:"email" validate:"required,email"`
    Age      int      `json:"age" validate:"gte=18,lte=65"`
    Tags     []string `json:"tags" validate:"dive,required"`
}
```

* `Username` 必填，字母数字组合，长度 3\~20
* `Email` 必填且格式合法
* `Age` 必须在 18\~65
* `Tags` 每个元素都必须非空

---

如果你需要，我可以帮你整理一张 **完整的 validator 标签参考表**，包括 **所有常用格式 + 使用场景 + 示例**，方便你直接查用。

你希望我整理吗？
