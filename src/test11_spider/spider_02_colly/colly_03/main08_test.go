package colly_03

//
//import (
//	"fmt"
//	"github.com/otiai10/gosseract/v2"
//	"testing"
//)
//
//// 使用colly爬虫的登陆遇到验证码时改如何操作？？？
//
//// 参考：https://blog.csdn.net/asfdsgdf/article/details/142949314
//
//// colly爬虫、gocv图像处理，gosseract是ocr识别，
////go get -u github.com/gocolly/colly/v2
//// brew install pkg-config
////go get -u gocv.io/x/gocv
////go get -u github.com/otiai10/gosseract/v2
////go get -u github.com/avast/retry-go
//
//func TestMain0801(t *testing.T) {
//	// 验证码识别
//
//	//img, err := preprocessImage("yzm.jpg")
//	//if err != nil {
//	//	return
//	//}
//	//
//	//fmt.Println(img)
//	//
//	//// 处理图像
//	//processedImage, err := preprocessImage("captcha.png")
//	//if err != nil {
//	//	fmt.Println("处理图像失败:", err)
//	//	return
//	//}
//
//	// 使用 Tesseract 进行 OCR 识别
//	client := gosseract.NewClient()
//	defer client.Close()
//
//	err := client.SetImage("yzm.jpg")
//	if err != nil {
//		fmt.Println("SetImage 识别失败:", err)
//		return
//	}
//	text, err := client.Text()
//	if err != nil {
//		fmt.Println("OCR 识别失败:", err)
//	} else {
//		fmt.Println("识别结果:", text)
//	}
//
//}
//
////func preprocessImage(imgPath string) (image.Image, error) {
////	img := gocv.IMRead(imgPath, gocv.IMReadColor)
////
////	// 转为灰度图像
////	grayImg := gocv.NewMat()
////	gocv.CvtColor(img, &grayImg, gocv.ColorBGRToGray)
////
////	// 二值化处理
////	binaryImg := gocv.NewMat()
////	gocv.Threshold(grayImg, &binaryImg, 150, 255, gocv.ThresholdBinary)
////
////	// 保存处理后的图像
////	gocv.IMWrite("captcha_processed.png", binaryImg)
////
////	return binaryImg.ToImage()
////}
