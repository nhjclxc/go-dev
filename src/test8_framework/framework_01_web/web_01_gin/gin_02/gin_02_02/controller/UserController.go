package controller

import (
	"gin_02_02/model/user"
	md5 "gin_02_02/utils/encrypt/md5"
	"gin_02_02/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// 定义 UserController 数据结果
type UserController struct {
	BaseController
}

const expire = 1 * 60 * 60 * 24

var userDataList = []user.User{
	user.User{Id: 1, Username: "zhangsan1", Password: "abc1231", Name: "张三1", Age: 21, Addr: "北京1"},
	user.User{Id: 2, Username: "zhangsan2", Password: "abc1232", Name: "张三2", Age: 22, Addr: "北京2"},
	user.User{Id: 3, Username: "zhangsan3", Password: "abc1233", Name: "张三3", Age: 23, Addr: "北京3"},
	user.User{Id: 4, Username: "zhangsan5", Password: "abc1234", Name: "张三4", Age: 24, Addr: "北京4"},
	user.User{Id: 5, Username: "zhangsan4", Password: "abc1235", Name: "张三5", Age: 25, Addr: "北京5"},
}

func findByUsername(username string) *user.User {
	if username == "" {
		return nil
	}
	for _, u := range userDataList {
		if u.Username == username {
			return &u
		}
	}
	return nil
}

func findById(id int) *user.User {
	if id == 0 {
		return nil
	}
	for _, u := range userDataList {
		if u.Id == id {
			return &u
		}
	}
	return nil
}

// 登录
func (this *UserController) Login(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			this.Error(context, err)
		}
	}()

	// 获取参数
	userLogin := user.UserLogin{}
	context.ShouldBind(&userLogin)

	log.Println("userLogin: ", userLogin)

	// 判空
	if userLogin.Username == "" {
		panic("用户名不能为空！！！")
	}
	if userLogin.Password == "" {
		panic("密码不能为空！！！")
	}

	// 校验用户名密码
	userDB := findByUsername(userLogin.Username)
	if userDB == nil {
		panic("用户不存在，请先注册")
	}

	// md5(abc1231) ===>>> c39fa77ca90e664db1e4d670e2b353b6
	// 鉴于密码不能进行铭文传输，因此这里使用简单的md5加密串来传输匹配
	if !md5.Match(userDB.Password, userLogin.Password) {
		//if userLogin.Password != userDB.Password {
		panic("密码不正确！！！")
	}

	// 生成 Token
	token, err := jwt.GenerateToken(strconv.Itoa(userDB.Id), expire)
	if err != nil {
		return
	}

	// token 返回前端
	this.Success(context, "Bearer "+token)
}

// 根据id查询详细
func (this *UserController) GetById(context *gin.Context) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		this.Error(context, err)
	//	}
	//}()

	id := context.Query("id")

	if id == "" {
		panic("id不能为空！！！")
	}
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	u := findById(atoi)

	value, exists := context.Get("uid")
	log.Printf("Authentication 接口鉴权 中带到接口里面的参数：%v, %v \n", value, exists)

	// 在子协程里面是否可以获取到上下文数据呢？？？
	go func() {
		// 延迟处理，让主协程先返回，看看子协程是否还能获取到 上下文数据
		time.Sleep(5 * time.Second)

		value2, exists2 := context.Get("uid")
		log.Printf("go - Authentication 接口鉴权 中带到接口里面的参数：%v, %v \n", value2, exists2)
	}()

	// 在子协程里面显然是不能获取到主协程上下文数据的
	// 因此，gin 给我们提供了一个方法 context.Copy() 来将主协程数据拷贝一份出来，已在子协程中使用
	copyContext := context.Copy()
	go func() {
		// 延迟处理，让主协程先返回，看看子协程是否还能获取到 上下文数据
		time.Sleep(5 * time.Second)

		value3, exists3 := copyContext.Get("uid")
		log.Printf("copyContext - Authentication 接口鉴权 中带到接口里面的参数：%v, %v \n", value3, exists3)
	}()

	this.Success(context, u)
}

// 分页查询
func (this *UserController) PageList(context *gin.Context) {
	this.Success(context, nil)
}
