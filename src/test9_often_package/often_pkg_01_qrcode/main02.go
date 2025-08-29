package main

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"net/http"
	"os"
)

// go get -u github.com/boombuler/barcode
func main02() {

	qrCode, _ := qr.Encode("https://github.com/nhjclxc", qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create("./src/test5_test/test_often_package/often_pkg_01_qrcode/qrcode05.png")
	defer file.Close()

	png.Encode(file, qrCode)

	http.HandleFunc("/qrcode", generateQRCode)
	http.ListenAndServe(":8080", nil)
}

func generateQRCode(w http.ResponseWriter, r *http.Request) {
	qrCode, _ := qr.Encode("https://github.com/nhjclxc", qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	// 创建内存缓冲区，不落盘
	var buf bytes.Buffer
	_ = png.Encode(&buf, qrCode)

	// 返回图片流给前端
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}
