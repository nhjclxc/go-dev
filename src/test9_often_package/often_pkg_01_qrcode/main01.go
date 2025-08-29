package main

import (
	"github.com/skip2/go-qrcode"
	"image/color"
)

// go get -u github.com/skip2/go-qrcode
func main01() {

	// 创建一个256x256的PNG图片：
	//var png []byte
	//png, err := qrcode.Encode("https://github.com/nhjclxc", qrcode.Medium, 256)
	//err = os.WriteFile("./src/test5_test/test_often_package/often_pkg_01_qrcode/qrcode01.png", png, 0644)
	//if err != nil {
	//	panic(err)
	//}

	//创建一个256x256的PNG图像并写入文件：
	//qrcode.WriteFile("https://github.com/nhjclxc", qrcode.Medium, 256,
	//	"./src/test5_test/test_often_package/often_pkg_01_qrcode/qrcode02.png")

	//创建具有自定义颜色的256x256 PNG图像并写入文件：
	qrcode.WriteColorFile("https://github.com/nhjclxc", qrcode.Medium, 256, color.Black, color.White,
		"./src/test5_test/test_often_package/often_pkg_01_qrcode/qrcode03.png")

}
