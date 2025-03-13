package main

//
//import (
//	"bytes"
//	"encoding/hex"
//	"fmt"
//	"log"
//	"strconv"
//	"strings"
//
//	// https://blog.csdn.net/chenxin0215/article/details/129178592
//	// go get github.com/tarm/serial
//	// go install github.com/tarm/serial
//	// go mod download
//	"github.com/jacobsa/go-serial/serial"
//)
//
//// 解析称重仪表的十六进制数据
//func parseWeightData(hexData string) (float64, error) {
//	// 将十六进制字符串转换为字节数组
//	data, err := hex.DecodeString(hexData)
//	if err != nil {
//		return 0, fmt.Errorf("十六进制解码错误: %v", err)
//	}
//
//	// 解析符号位
//	sign := 1
//	if data[3] == 0x2D { // `2D` 代表负数
//		sign = -1
//	} else if data[3] != 0x2B { // `2B` 代表正数
//		return 0, fmt.Errorf("未知的符号位: %X", data[3])
//	}
//
//	// 解析重量数据（ASCII 数字）
//	weightStr := string(data[4:10]) // 6 个字符
//	weightInt, err := strconv.Atoi(weightStr)
//	if err != nil {
//		return 0, fmt.Errorf("重量数据解析错误: %v", err)
//	}
//
//	// 解析小数点位置
//	decimalPlaces := int(data[10] - '0') // `33` = '3'，对应 3 位小数
//
//	// 计算最终重量值
//	weight := float64(weightInt) / float64(pow10(decimalPlaces)) * float64(sign)
//
//	return weight, nil
//}
//
//// 计算 10 的幂次方
//func pow10(n int) int {
//	result := 1
//	for i := 0; i < n; i++ {
//		result *= 10
//	}
//	return result
//}
//
//func main() {
//
//	var buffer []byte = readData3()
//
//	hexData := bytes2String(buffer)
//
//	// 样例数据
//	hexData1 := "0241422B30303130393033313403" // 6.090
//	hexData2 := "0241422D30313231383033313703" // -12.180
//	hexData3 := "4603022B3031313531303031"     // 11510.000 kg
//	parseData(hexData1)
//	parseData(hexData2)
//	parseData(hexData3)
//	parseData(hexData)
//}
//
//func bytes2String(bytes []byte) string {
//	var sb strings.Builder
//	for _, b := range bytes {
//		sb.WriteString(fmt.Sprintf("%d", b))
//	}
//	return sb.String()
//}
//
//func readData3() []byte {
//	// Set up options.
//	options := serial.OpenOptions{
//		PortName:        "COM3", // 根据你的串口修改
//		BaudRate:        9600,
//		DataBits:        8,
//		StopBits:        1,
//		MinimumReadSize: 4,
//	}
//
//	// Open the port.
//	port, err := serial.Open(options)
//	if err != nil {
//		log.Fatalf("serial.Open: %v", err)
//	}
//
//	// Make sure to close it later.
//	defer port.Close()
//
//	// Write 4 bytes to the port.
//	//b := []byte{0x02, 0x41, 0x41, 0x30, 0x30, 0x03}
//	var b []byte
//	b = []byte{0x02, 0x41, 0x42, 0x30, 0x30, 0x03}
//	n, err := port.Write(b)
//	if err != nil {
//		log.Fatalf("port.Write: %v", err)
//	}
//
//	fmt.Println("Wrote", n, "bytes.")
//
//	buffer := make([]byte, 128)
//	n2, err2 := port.Read(buffer)
//	if err2 != nil {
//		log.Fatal(err2)
//	}
//	data := buffer[:n2]
//	fmt.Println(string(data))
//	fmt.Println(strconv.Quote(string(data))) // 以可视化方式显示特殊字符
//	// 去除可能的无效字节（如 0 值）
//	data = bytes.Trim(data, "\x00")
//	// 转换为字符串并打印
//	fmt.Println("aaaaa --- ", string(data))
//
//	//tenToHex(string(data))
//	return buffer
//}
//
//func parseData(hexData1 string) {
//	// 解析数据
//	weight1, err := parseWeightData(hexData1)
//	if err != nil {
//		fmt.Println("解析错误:", err)
//	} else {
//		fmt.Printf("解析出的重量: %.3f kg\n", weight1)
//	}
//}
