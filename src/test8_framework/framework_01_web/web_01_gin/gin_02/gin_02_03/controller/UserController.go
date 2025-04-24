package controller

import (
	"fmt"
	captcha "gin_02_03/utils/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 定义 UserController 结构体来保存这个实例的相关数据
type UserController struct {
	// userService UsersService
}

// 定义 UserController 对应的接口方法
func (this *UserController) insertUser(context *gin.Context) {

	context.JSON(200, gin.H{
		"data": fmt.Sprintf("insertUser。"),
	})
}
func (this *UserController) GetById(context *gin.Context) {

	// 测试日志

	log.Println("Println")
	log.Printf("Printf")
	//log.Fatalln("Fatalln")

	context.JSON(200, gin.H{
		"data": fmt.Sprintf("getById。"),
	})
}

// GetCaptcha gin 中验证码的生成

// 使用一个 map[string]string 来模拟缓存reids
var cachePool map[string]string = make(map[string]string)

func (this *UserController) GetCaptcha(context *gin.Context) {

	captchaId, answer, b64s := captcha.GenerateCharCaptcha()
	fmt.Println(captchaId)
	fmt.Println(answer)
	//fmt.Println(b64s)
	cachePool[captchaId] = answer

	context.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"captchaId": captchaId,
			"b64s":      b64s,
		},
	})
}

// 校验验证码是否正确
func (this *UserController) ValidateCaptcha(context *gin.Context) {

	captchaId := context.Query("captchaId")
	answer := context.Query("answer")

	if cachePool[captchaId] != answer {
		context.JSON(200, gin.H{
			"code": 500,
			"data": fmt.Sprintf("验证码校验错误！！！"),
		})
		return
	}

	// 校验成功删除已存在的验证码
	delete(cachePool, captchaId)

	context.JSON(200, gin.H{
		"code": 200,
		"data": fmt.Sprintf("验证码校验成功。"),
	})

}

// 结构体校验示例
type UserInfo struct {
	//不能为空并且大于10
	Age       int       `form:"age" binding:"required,gt=10"`
	Name      string    `form:"name" binding:"required"`
	Birthday  time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	UserEmail string    `form:"userEmail" binding:"required,email"`
}

/*
	binding 语法

Go语言中常用的结构体标签（tag）语法，尤其是在使用像 Gin 框架时，binding 标签是用来进行参数校验的，背后通常依赖的是 go-playground/validator 库。

以下是常见的 binding 语法（校验器标签）分类和说明：
标签			含义
required	必填字段
omitempty	忽略空值字段的校验
len=10		长度必须等于10（字符串、切片等）
min=3		最小值为3（数字、切片、字符串等）
max=20		最大值为20
gt=10		必须大于10
lt=100		必须小于100
gte=1		大于等于1
lte=10		小于等于10
eq=10		等于10
ne=5		不等于5



✅ 字符串特有的标签

标签				含义
email			必须是有效邮箱
url				必须是有效URL
uuid			必须是有效UUID
alphanum		只能包含字母和数字
alpha			只能包含字母
numeric	只		能包含数字
contains=foo	必须包含子串 "foo"
startswith=abc	必须以 abc 开头
endswith=xyz	必须以 xyz 结尾



✅ 集合类校验

标签			含义
unique		slice 中元素必须唯一
dive		用于 slice/map 的每个元素校验



✅ 条件和组合标签

标签格式						含义
required_without=Field1		当前字段在Field1为空时必须有值
required_with=Field1		Field1有值时当前字段也必须有值
required_if=Field1 value	Field1等于value时必须有值
eqfield=OtherField			当前字段值等于另一个字段值
nefield=OtherField			当前字段值不等于另一个字段值



✅ 时间相关标签

标签					含义
time_format			指定时间格式（与time.Parse一致），即2006-01-02 15:04:05
time_utc			时间是否为UTC
time_location		时区，例如 Asia/Shanghai



✅ 自定义验证器
你也可以自定义校验函数，并在初始化时注册，例如：
```
func myValidator(fl validator.FieldLevel) bool {
    return fl.Field().String() == "go"
}

// 注册：
if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("isgo", myValidator)
}
```

Name string `binding:"required,isgo"`


*/
