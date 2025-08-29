package main

import (
	"fmt"
	"github.com/tuotoo/qrcode"
	"os"
)

// go get -u github.com/tuotoo/qrcode
// 识别二维码
func main() {
	//tuotoo/qrcode识别二维码
	//	动态二值化
	//	提升图片扫描的速度:OK
	//	修复标线取值: OK
	//	容错码纠正数据:OK
	//	数据编码方式 Numbert alphanumeric 8-bit byte: OK Kanji
	//	识别各角度倾斜的二维码

	fi, err := os.Open("./src/test5_test/test_often_package/often_pkg_01_qrcode/qrcode05.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
	//fmt.Println(qrmatrix.Size)
	//fmt.Println(qrmatrix.Data)
	//fmt.Println(qrmatrix.OrgImage)
	//fmt.Println(qrmatrix.OrgSize)
	//fmt.Println(qrmatrix.Points)
	//fmt.Println(qrmatrix.Version())

}
