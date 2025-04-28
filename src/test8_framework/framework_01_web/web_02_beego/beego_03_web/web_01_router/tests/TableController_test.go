package test

import (
	beego "github.com/beego/beego/v2/server/web"
	"testing"
	"web_01_router/controllers"
)

// TestGet is a sample to run an endpoint test
func TestTableController01(t *testing.T) {

	// get http://localhost:8080/api/user/helloworld
	// you will see return "Hello, world"
	ctrl := &controllers.UserController{}
	beego.AutoPrefix("api", ctrl)
	beego.Run()

}
