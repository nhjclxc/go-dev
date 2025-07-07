package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go-dev/src/test_utils/utils_02_document/document_02_excel/excel_02/excelUtils"
	"net/http"
	"net/url"
	"strconv"
)

// 使用gin框架实现文件下载
func main() {
	r := gin.Default()

	r.GET("/download", func(c *gin.Context) {

		headerKeys := []string{"LocalDateTime", "LocalDate", "LocalTime", "Date",
			"String", "Integer", "Float", "Double", "Long", "BigDecimal", "Boolean",
			"ImageURL", "Base64", "PhoneNumber",
		}

		headerValues := []string{"localDateTime数据", "localDate数据", "localTime数据", "date数据",
			"string数据", "integer数据", "aFloat数据", "aDouble数据", "aLong数据", "bigDecimal数据", "aBoolean数据",
			"图片链接", "Base64数据", "手机号",
		}

		data := getTestObject()
		file := excelUtils.ExportExcel[TestObject](data, 0, 1, headerKeys, headerValues)
		err := file.SaveAs("output2.xlsx")
		if err != nil {
			fmt.Println("保存失败：", err)
		}
		// 设置响应头 - 解决中文文件名问题
		filename := url.QueryEscape("测试.xlsx")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename*=UTF-8''"+filename)

		// 写入响应
		_ = file.Write(c.Writer)

	})


	r.POST("/upload", func(c *gin.Context) {

		c.Request.ParseMultipartForm(10 << 20) // 限制最大10MB
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "无法获取文件",
			})
			return
		}
		defer file.Close()

		headerKeys := []string{"LocalDateTime", "LocalDate", "LocalTime", "Date",
			"String", "Integer", "Float", "Double", "Long", "BigDecimal", "Boolean",
			"ImageURL", "Base64", "PhoneNumber",
		}

		testObjectList := excelUtils.ImportExcelByte[TestObject](file, 0, 1, headerKeys)

		for _, val := range testObjectList {
			fmt.Printf("testObjectList = %#v \n", val)
		}

	})


	r.Run(":8080")
}



func downloadCSV(c *gin.Context) {
	// 设置响应头 - 解决乱码和中文文件名问题
	filename := url.QueryEscape("测试.csv")
	c.Header("Content-Disposition", "attachment; filename*=UTF-8''"+filename)
	c.Header("Content-Type", "text/csv; charset=utf-8")

	// 写入UTF-8 BOM头，防止中文乱码（可选，某些旧版Excel需要）
	_, _ = c.Writer.Write([]byte{0xEF, 0xBB, 0xBF})

	// 创建CSV写入器
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// 写入表头
	_ = writer.Write([]string{"姓名", "年龄"})

	// 写入数据行
	data := [][]string{
		{"xxx", "18"},
		{"yyy", "19"},
		{"zzz", "20"},
	}

	for _, row := range data {
		_ = writer.Write(row)
	}
}

func downloadExcel(c *gin.Context) {
	// 创建Excel文件
	f := excelize.NewFile()
	defer f.Close()

	// 创建工作表
	index, _ := f.NewSheet("Sheet1")

	// 设置表头
	_ = f.SetCellValue("Sheet1", "A1", "姓名")
	_ = f.SetCellValue("Sheet1", "B1", "年龄")

	// 设置数据
	data := [][]interface{}{
		{"xxx", "18"},
		{"yyy", "19"},
		{"zzz", "20"},
	}

	for i, row := range data {
		_ = f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), row[0])
		_ = f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), row[1])
	}

	// 设置活动工作表
	f.SetActiveSheet(index)

	// 设置响应头 - 解决中文文件名问题
	filename := url.QueryEscape("测试.xlsx")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename*=UTF-8''"+filename)

	// 写入响应
	_ = f.Write(c.Writer)
}
