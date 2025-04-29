package controllers

import (
	"encoding/json"
	"fmt"
	beecontext "github.com/beego/beego/v2/server/web/context"
	"strconv"
	"time"
	"web_06_input/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
	UserMap map[string]*models.User
}

// Init 初始化控制器
func (c *UserController) Init(ctx *beecontext.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	// 自定义初始化逻辑（如初始化 UserMap）
	c.UserMap = make(map[string]*models.User)
}

func (this *UserController) printUserList() {

	if len(this.UserMap) <= 0 {
		fmt.Println("UserList暂无数据")
		return
	}

	count := 1
	for _, user := range this.UserMap {
		fmt.Printf("第 %d 个数据，详细数据为：%v \n", count, user)
		count++
	}

}

// @Title InsertUser
// @Description 新增用户操作
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /insertUser [post]
func (this *UserController) InsertUser(ctx *beecontext.Context) {
	// // http://localhost:8080/v1/user/insertUser
	/*
		{
		    "id": "zxc123",
		    "username": "root",
		    "password": "root123"
		}
	*/

	this.printUserList()

	var user models.User
	json.Unmarshal(ctx.Input.RequestBody, &user)

	user.UserId = "user_id_" + strconv.FormatInt(time.Now().UnixNano(), 10)

	this.UserMap[user.UserId] = &user

	this.printUserList()

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": user.UserId,
	})

}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /deleteUser/:uid [delete]
func (this *UserController) DeleteUser(ctx *beecontext.Context) {

	//ctx.JSONResp(map[string]any{
	//	"code": 200,
	//	"data": user.Id,
	//})

}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (this *UserController) UpdateUser(ctx *beecontext.Context) {

	// 获取 userId
	userId := ctx.Input.Param(":userId")

	// 获取请求体
	var user models.User
	json.Unmarshal(ctx.Input.RequestBody, &user)

	var user2 models.User
	ctx.BindJSON(&user2)

	fmt.Printf("userId = %v \n", userId)
	fmt.Printf("user = %v \n", user)
	fmt.Printf("user2 = %v \n", user2)

	// 执行修改
	if userId == "" {
		ctx.JSONResp(map[string]any{
			"code": 200,
			"msg":  "userId 不能为空！！！",
		})
		return
	}

	userDB := this.UserMap[userId]

	fmt.Printf("修改前 userDB = %v \n", userDB)
	userDB.Username = user.Username
	userDB.Password = user.Password
	fmt.Printf("修改后 userDB = %v \n", userDB)
	fmt.Printf("修改后 this.UserMap[userId] = %v \n", this.UserMap[userId])

	ctx.JSONResp(map[string]any{
		"code": 200,
		"msg":  this.UserMap[userId],
	})

}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /getAll [get]
func (this *UserController) GetAll(ctx *beecontext.Context) {
	userList := make([]models.User, 0, len(this.UserMap))
	for _, userPtr := range this.UserMap {
		if userPtr != nil {
			userList = append(userList, *userPtr) // 取指针的值
		}
	}

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": userList,
	})
}

// @Title GetById
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /getById/:uid [get]
func (this *UserController) GetById(ctx *beecontext.Context) {

	// http://localhost:8080/v1/user/getById/123?username=root&password=root@123

	// 获取路径参数
	// 在参数前面加一个冒号 :
	// 注意这个路径参数一定要和 rotuer 里面定义的路径参数一致，否则无法获取
	userId := ctx.Input.Param(":userId")
	userId2 := ctx.Input.Query(":userId")

	// 获取查询参数
	username := ctx.Input.Query("username")
	password := ctx.Input.Query("password")

	fmt.Println("userId = ", userId)
	fmt.Println("userId2 = ", userId2)
	fmt.Println("username = ", username)
	fmt.Println("password = ", password)

	user := this.UserMap[userId]
	if user == nil {
		//panic("不存在对应 userId 的用户！！！")
		ctx.JSONResp(map[string]any{
			"code": 500,
			"msg":  "不存在对应 userId 的用户！！！",
		})
		return
	}

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": user,
	})

	fmt.Println("this.Data = ", this.Data)
	fmt.Println("this.Ctx = ", this.Ctx)
}

func (this *UserController) GetById2() {
	fmt.Println("this.Data = ", this.Data)
	fmt.Println("this.Ctx = ", this.Ctx)

	userId := this.Ctx.Input.Param(":userId") // 获取路径参数
	queryParam := this.Ctx.Input.Query("key") // 获取查询参数
	bodyData := this.Ctx.Input.RequestBody    // 获取请求体
	fmt.Println("userId = ", userId)
	fmt.Println("queryParam = ", queryParam)
	fmt.Println("bodyData = ", bodyData)
	this.Data["json"] = "result" // 设置响应数据
	this.ServeJSON()             // 返回JSON
}

// @Title GetList
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /getList/:uid [get]
func (this *UserController) GetList(ctx *beecontext.Context) {

}

// @Title GetListPage
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /getListPage/:uid [get]
func (this *UserController) GetListPage(ctx *beecontext.Context) {

}

//
// 文件上传与下载学习
// https://beegodoc.com/zh/developing/web/file/

// 文件上传方式1
// GetFile(key string) (multipart.File, *multipart.FileHeader, error)：
// 该方法主要用于用户读取表单中的文件名 the_file，然后返回相应的信息，用户根据这些变量来处理文件上传、过滤、保存文件等。
func (this *UserController) Upload1(ctx *beecontext.Context) {
	form := ctx.Request.MultipartForm
	file := form.File["file1"][0]
	fmt.Println("file.Size = ", file.Size)
	open, err := file.Open()
	if err != nil {
		return
	}
	buf := make([]byte, 512)
	open.Read(buf)

	fmt.Println("data  = ", string(buf))

	getFile, m, err := this.GetFile("file1")
	if err != nil {
		return
	}

	buf2 := make([]byte, 512)
	getFile.Read(buf2)

	fmt.Println("m.Size  = ", m.Size)
	fmt.Println("m.Filename  = ", m.Filename)
	fmt.Println("data2  = ", string(buf2))
}

// 文件上传方式2
// SaveToFile(fromfile, tofile string) error：
// 该方法是在 GetFile 的基础上实现了快速保存的功能。fromfile是提交时候表单中的name
func (this *UserController) Upload2(ctx *beecontext.Context) {
	//this.SaveToFile()

	fmt.Println("111")

}

// 文件下载
// func (output *BeegoOutput) Download(file string, filename ...string) {}
func (this *UserController) DownloadFile(ctx *beecontext.Context) {

	fmt.Println("12345678")

	//this.Ctx.Output.Download()
	// The file LICENSE is under root path.
	// and the downloaded file name is license.txt
	//this.Ctx.Output.Download("LICENSE", "license.txt")
	// 尤其要注意的是，Download方法的第一个参数，是文件路径，也就是要下载的文件；第二个参数是不定参数，代表的是用户保存到本地时候的文件名。
	//如果第一个参数使用的是相对路径，那么它代表的是从当前工作目录开始计算的相对路径。

	// 第一个参数 file：是这个文件的路径，先对于 web_06_input.exe（当前工作目录） 的路径
	// 第二个参数 filename：是要输出的文件名
	ctx.Output.Download("go.mod", "aaaaa.mod")

	fmt.Println("qwertyui")

}
