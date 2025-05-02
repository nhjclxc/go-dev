package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"log"
	"strings"
)

/*
用户也可以通过声明式的写法来表达某个字段需要遵守的校验规则，声明式写法是通过结构体的标签来实现的：

验证函数写在 "valid" 的标签里
各个函数之间用分号 ";" 分隔，分号后面可以有空格
参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
正则函数(Match)的匹配模式用两斜杠 "/" 括起来
各个函数的结果的 key 值为字段名.验证函数名
*/

type User02 struct {
	Id     int
	Name   string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
	Age    int    `valid:"Range(1, 140)"`            // 1 <= Age <= 140，超出此范围即为不合法
	Email  string `valid:"Email; MaxSize(100)"`      // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
	Mobile string `valid:"Mobile"`                   // Mobile 必须为正确的手机号
	IP     string `valid:"IP"`                       // IP 必须为一个正确的 IPv4 地址
}
/*
在实现该接口的时候，只需要将错误信息写入validation.Validation.

StructTag 可用的验证函数：

Required 不为空，即各个类型要求不为其零值
Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证
Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
Alpha alpha 字符，有效类型：string，其他类型都将不能通过验证
Numeric 数字，有效类型：string，其他类型都将不能通过验证
AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf("%v", obj).Match)
AlphaDash alpha 字符或数字或横杠 -_，有效类型：string，其他类型都将不能通过验证
Email 邮箱格式，有效类型：string，其他类型都将不能通过验证
IP IP 格式，目前只支持 IPv4 格式验证，有效类型：string，其他类型都将不能通过验证
Base64 base64 编码，有效类型：string，其他类型都将不能通过验证
Mobile 手机号，有效类型：string，其他类型都将不能通过验证
Tel 固定电话号，有效类型：string，其他类型都将不能通过验证
Phone 手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
ZipCode 邮政编码，有效类型：string，其他类型都将不能通过验证
 */

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *User02) Valid(v *validation.Validation) {

	logs.Info("是否执行自定义验证！！！")

	// 以下可以实现一些自定义验证
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}

func main() {
	valid := validation.Validation{}
	user := User02{Name: "Beego", Age: 2, Email: "dev@web.me"}
	user.Mobile = "19325634185"
	user.IP = "127.0.0.1"
	b, err := valid.Valid(&user)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}


	// 注册自定义验证
	_ = validation.AddCustomFunc("ChinaAddress", func(v *validation.Validation, obj interface{}, key string) {
		addr, ok := obj.(string)
		if !ok {
			return
		}
		if !strings.HasPrefix(addr, "China") {
			v.AddError(key, "China address only")
		}
	})
}
