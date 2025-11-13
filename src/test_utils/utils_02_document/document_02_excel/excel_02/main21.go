package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go-dev/src/test_utils/utils_02_document/document_02_excel/excel_02/excelUtils"
	"math/big"
	"net/http"
	"time"
)

type TestObject struct {
	LocalDateTime time.Time  // 对应 Java 的 LocalDateTime
	LocalDate     time.Time  // 对应 Java 的 LocalDate（可只用日期部分）
	LocalTime     time.Time  // 对应 Java 的 LocalTime（可只用时间部分）
	Date          time.Time  // 对应 java.util.Date
	String        string     // 对应 String
	Integer       int        // 对应 Integer（可用 int32 或 int64 视情况而定）
	Float         float32    // 对应 Float
	Double        float64    // 对应 Double
	Long          int64      // 对应 Long
	BigDecimal    *big.Float // 对应 BigDecimal（精确计算建议使用 math/big）
	Boolean       bool       // 对应 Boolean
	ImageURL      string     // 对应 URL 字符串
	Base64        string     // 对应 Base64 编码字符串
	PhoneNumber   string     // 电话号码推荐用 string（避免丢失前导0）
}

// go get -u github.com/xuri/excelize/v2
func main01() {

	doExportTest111()

	//doImportTest111()

}

func doExportTest111() {

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

}

// 如何将excelize.File数据返回给gin响应
func ExportExcel(c *gin.Context) {
	f := excelize.NewFile()
	sheetName := "Sheet1"

	// 写一些示例数据
	f.SetCellValue(sheetName, "A1", "Name")
	f.SetCellValue(sheetName, "B1", "Score")
	f.SetCellValue(sheetName, "A2", "Alice")
	f.SetCellValue(sheetName, "B2", 95)

	// 将 Excel 写入内存缓冲区
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		c.String(http.StatusInternalServerError, "生成Excel失败: %v", err)
		return
	}

	// 设置下载响应头
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", `attachment; filename="report.xlsx"`)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	// 写入响应体
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())

	// f.Write(&buf) 是把 Excel 文件内容写进内存；
	//c.Data(...) 是把内存中的 Excel 数据发送给客户端。

	// 以下是简写

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", `attachment; filename="report.xlsx"`)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	if err := f.Write(c.Writer); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

}
func doImportTest111() {

	headerKeys := []string{"LocalDateTime", "LocalDate", "LocalTime", "Date",
		"String", "Integer", "Float", "Double", "Long", "BigDecimal", "Boolean",
		"ImageURL", "Base64", "PhoneNumber",
	}

	testObjectList := excelUtils.ImportExcelFile[TestObject]("output.xlsx", 0, 1, headerKeys)

	for _, val := range testObjectList {
		fmt.Printf("testObjectList = %#v \n", val)
	}

}

//func doImportTest() {
//
//	// 1. 打开 Excel 文件
//	f, err := excelize.OpenFile("output.xlsx")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	// 2. 获取默认工作表名称（如 Sheet1）
//	sheetName := f.GetSheetName(f.GetActiveSheetIndex())
//
//	// 3. 获取所有行的数据（按行读取）
//	rows, err := f.GetRows(sheetName)
//	if err != nil {
//		panic(err)
//	}
//
//	testObjectList := make([]TestObject, len(rows))
//
//	// headerKeys 和 headerValues必须一一对应
//	headerKeys := []string{"LocalDateTime", "LocalDate", "LocalTime", "Date",
//		"String", "Integer", "Float", "Double", "Long", "BigDecimal", "Boolean",
//		"ImageURL", "Base64", "PhoneNumber",
//	}
//
//	var firstColAxis int32 = 0
//	var firstRowAxis int = 1
//
//
//	// 4. 遍历行与列
//	for rowIndex, row := range rows {
//		if rowIndex < firstRowAxis {
//			continue
//		}
//		//fmt.Printf("Row %d:\n", rowIndex)
//		testObject := TestObject{}
//		for colIndex, cell := range row {
//			if int32(colIndex) < firstColAxis {
//				break
//			}
//			//fmt.Printf("  Col %d: %s\n", colIndex, cell)
//			err := setFieldByName(&testObject, headerKeys[colIndex], cell)
//			if err != nil {
//				fmt.Println("设置错误", err)
//				break
//			}
//		}
//		testObjectList[ rowIndex - firstRowAxis] = testObject
//	}
//
//	for val := range testObjectList {
//		fmt.Printf("testObjectList = %#v \n", val)
//	}
//
//}
//
//func doExportTest() {
//	// 1. 导入依赖包  "github.com/xuri/excelize/v2"
//
//	//2. 创建实例
//	sheetName := "testSheet"
//
//	// func WriteToExcel(headers []string, data [][]interface{}, options ExcelOptions)
//	file := excelize.NewFile()
//	sheetIndex, _ := file.NewSheet(sheetName)
//	file.SetActiveSheet(sheetIndex) // 默认sheet
//
//	var firstColAxis int32 = 0
//	var firstRowAxis int = 1
//
//	// headerKeys 和 headerValues必须一一对应
//	headerKeys := []string{"LocalDateTime", "LocalDate", "LocalTime", "Date",
//		"String", "Integer", "Float", "Double", "Long", "BigDecimal", "Boolean",
//		"ImageURL", "Base64", "PhoneNumber",
//	}
//
//	headerValues := []string{"localDateTime数据", "localDate数据", "localTime数据", "date数据",
//		"string数据", "integer数据", "aFloat数据", "aDouble数据", "aLong数据", "bigDecimal数据", "aBoolean数据",
//		"图片链接", "Base64数据", "手机号",
//	}
//
//	// 3. 设置表头
//	for i, headerName := range headerValues {
//		tempColAxis := firstColAxis + int32(i)
//		writeCell(file, sheetName, tempColAxis, firstRowAxis, headerName)
//	}
//
//	// 4. 填充数据
//	data := getTestObject()
//	for i, item := range data {
//		tempRowAxis := firstRowAxis + i + 1
//		for j, header := range headerKeys {
//			value := GetFieldByName(&item, header)
//			tempColAxis := firstColAxis + int32(j)
//			writeCell(file, sheetName, tempColAxis, tempRowAxis, value)
//		}
//	}
//
//	//5. 设置样式
//
//	//6. 写到输出流
//	// 保存为文件
//	err := file.SaveAs("output.xlsx")
//	if err != nil {
//		fmt.Println("保存失败：", err)
//	}
//}
//
//// 设置单元格内容
//func writeCell(file *excelize.File, sheetName string, collAxis int32, rowAxis int, value interface{}) {
//	// SetCellValue：填写的单元格下标，先列下标，再行下标。例如，第一行第一列：A1，第二行第三列：C2。
//	err := file.SetCellValue(sheetName, getCellIndex(collAxis, rowAxis), value)
//	if err != nil {
//		fmt.Println("excel set cell value error:%v", err)
//	}
//}
//
//// 获取单元格下标
//func getCellIndex(colAxis int32, rowAxis int) string {
//	return fmt.Sprintf("%c%d", colAxis + 97, rowAxis)
//}
//
//type ExcelOptions struct {
//	ColWidth float64
//	RowHeight float64
//	FirstColAxis int32
//	FirstRowAxis int
//}
//
//// 设置样式
//func setStyle(file *excelize.File, sheetName string, maxCol int32, maxRow int, options ExcelOptions) {
//	err := file.SetColWidth(sheetName, string(options.FirstColAxis), string(maxCol), options.ColWidth) // 设置列宽
//	if err != nil {
//		fmt.Println("excel SetColWidth error:%v", err)
//	}
//	err = file.SetRowHeight(sheetName, options.FirstRowAxis, options.RowHeight) // 设置行高
//	if err != nil {
//		fmt.Println("excel SetRowHeight error:%v", err)
//	}
//	style := &excelize.Style{
//		Font: &excelize.Font{
//			Bold: true,
//		},
//	}
//	styleId, err := file.NewStyle(style)
//	if err != nil {
//		fmt.Println("excel NewStyle error:%v", err)
//	}
//	err = file.SetCellStyle(sheetName, getCellIndex(options.FirstColAxis, options.FirstRowAxis), getCellIndex(maxCol, maxRow),
//		styleId)
//	if err != nil {
//		fmt.Println("excel SetCellStyle error:%v", err)
//	}
//}
//
//type TestObject struct {
//	LocalDateTime time.Time  // 对应 Java 的 LocalDateTime
//	LocalDate     time.Time  // 对应 Java 的 LocalDate（可只用日期部分）
//	LocalTime     time.Time  // 对应 Java 的 LocalTime（可只用时间部分）
//	Date          time.Time  // 对应 java.util.Date
//	String        string     // 对应 String
//	Integer       int        // 对应 Integer（可用 int32 或 int64 视情况而定）
//	Float         float32    // 对应 Float
//	Double        float64    // 对应 Double
//	Long          int64      // 对应 Long
//	BigDecimal    *big.Float // 对应 BigDecimal（精确计算建议使用 math/big）
//	Boolean       bool       // 对应 Boolean
//	ImageURL      string     // 对应 URL 字符串
//	Base64        string     // 对应 Base64 编码字符串
//	PhoneNumber   string     // 电话号码推荐用 string（避免丢失前导0）
//}
//
//
//// 泛型通用方法：通过字段名获取任意结构体的字段值
//func GetFieldByName[T any](obj *T, fieldName string) interface{} {
//	v := reflect.ValueOf(obj).Elem()
//	f := v.FieldByName(fieldName)
//	if !f.IsValid() {
//		return nil
//	}
//	if f.Type() == reflect.TypeOf(time.Time{}) {
//		fmt.Println("字段是 time.Time 类型")
//		t, ok := f.Interface().(time.Time)
//		if ok {
//			return FormatTime(t)
//		}
//	}
//	return f.Interface()
//}
//func FormatTime(t time.Time) string {
//	return t.Format("2006-01-02 15:04:05")
//}
//
//func ParseTime(s string) (time.Time, error) {
//	return time.Parse("2006-01-02 15:04:05", s)
//}
//
//
//// 泛型方法：通过字段名设置任意结构体字段的值
//func setFieldByName[T any](obj *T, fieldName string, value interface{}) error {
//	v := reflect.ValueOf(obj)
//	if v.Kind() != reflect.Ptr || v.IsNil() {
//		return errors.New("obj must be a non-nil pointer to struct")
//	}
//
//	elem := v.Elem()
//	if elem.Kind() != reflect.Struct {
//		return errors.New("obj must be a pointer to struct")
//	}
//
//	field := elem.FieldByName(fieldName)
//	if !field.IsValid() {
//		return fmt.Errorf("field '%s' does not exist", fieldName)
//	}
//	if !field.CanSet() {
//		return fmt.Errorf("field '%s' cannot be set", fieldName)
//	}
//
//	// 自动转换 value 到目标字段类型
//	targetType := field.Type()
//	converted, err := convertValue(value, targetType)
//	if err != nil {
//		return fmt.Errorf("failed to convert value: %v", err)
//	}
//
//	field.Set(converted)
//	return nil
//}
//
//// 类型转换逻辑
//func convertValue(value interface{}, targetType reflect.Type) (reflect.Value, error) {
//	val := reflect.ValueOf(value)
//
//	// 如果类型可直接转换
//	if val.Type().ConvertibleTo(targetType) {
//		return val.Convert(targetType), nil
//	}
//
//	// 字符串类型处理
//	if str, ok := value.(string); ok {
//		switch targetType.Kind() {
//		case reflect.String:
//			return reflect.ValueOf(str), nil
//		case reflect.Int, reflect.Int32, reflect.Int64:
//			i, err := strconv.ParseInt(str, 10, 64)
//			if err != nil {
//				return reflect.Value{}, err
//			}
//			return reflect.ValueOf(i).Convert(targetType), nil
//		case reflect.Float32, reflect.Float64:
//			f, err := strconv.ParseFloat(str, 64)
//			if err != nil {
//				return reflect.Value{}, err
//			}
//			return reflect.ValueOf(f).Convert(targetType), nil
//		case reflect.Bool:
//			b, err := strconv.ParseBool(str)
//			if err != nil {
//				if str == "1" || strings.ToUpper(str) == "TRUE" {
//					b = true
//				} else if str == "0"  || strings.ToUpper(str) == "FALSE" {
//					b = false
//				} else {
//					return reflect.Value{}, err
//				}
//			}
//			return reflect.ValueOf(b), nil
//
//		default:
//			// time.Time 特殊判断
//			if targetType == reflect.TypeOf(time.Time{}) {
//				t, err := ParseTime(str)
//				if err != nil {
//					return reflect.Value{}, err
//				}
//				return reflect.ValueOf(t), nil
//			}
//			// 放在 convertValue 的 string 判断分支中
//			if targetType == reflect.TypeOf(&big.Float{}) {
//				bf := new(big.Float)
//				bf, ok := bf.SetString(str)
//				if !ok {
//					return reflect.Value{}, fmt.Errorf("cannot parse '%s' to *big.Float", str)
//				}
//				return reflect.ValueOf(bf), nil
//			}
//			fmt.Printf("default: %v \n", value)
//		}
//	}
//
//	// 其他 fallback 情况
//	return reflect.Value{}, fmt.Errorf("unsupported conversion from %T to %s", value, targetType.String())
//}
//

func getTestObject() []TestObject {
	objects := make([]TestObject, 10)

	now := time.Now()

	for i := 0; i < 10; i++ {
		objects[i] = TestObject{
			LocalDateTime: now.Add(time.Duration(i) * time.Hour),
			LocalDate:     now.AddDate(0, 0, i),
			LocalTime:     time.Date(0, 1, 1, i, i, 0, 0, time.Local), // 模拟不同时间
			Date:          now.Add(time.Duration(i) * time.Minute),
			String:        fmt.Sprintf("字符串-%d", i),
			Integer:       i + 1,
			Float:         float32(i) + 0.1,
			Double:        float64(i) + 0.01,
			Long:          int64(i) * 1000,
			BigDecimal:    big.NewFloat(float64(i) + 0.001),
			Boolean:       i%2 == 0,
			ImageURL:      fmt.Sprintf("https://example.com/image%d.png", i),
			Base64:        fmt.Sprintf("YmFzZTY0LWRhdGEt%d", i),
			PhoneNumber:   fmt.Sprintf("138000000%02d", i),
		}
	}

	//// 打印一下其中一个看看效果
	//for i, obj := range objects {
	//	fmt.Printf("对象 %d: %+v\n", i, obj)
	//}

	return objects

}
