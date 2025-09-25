package colly_03

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strings"
	"testing"
)

// https://m.w3cschool.cn/colly/colly-examples-cryptocoinmarketcap.html
// 用 Colly 抓取【https://hellogithub.com/en/report/tiobe】里面的编程语言排行榜到csv里面，一键导出 cryptocoinmarketcap.csv，Excel 直接打开。

func TestMain01(t *testing.T) {

	c := colly.NewCollector()

	//c.OnHTML("table", func(e *colly.HTMLElement) {
	c.OnHTML("table.w-min.min-w-full.table-fixed.divide-y-2.divide-gray-200.text-sm.dark\\:divide-gray-700", func(e *colly.HTMLElement) {

		if strings.Contains(e.Text, "RankLanguagePopularity📊") {
			//fmt.Println("table", e.Text)
			headers := make([]string, 0)
			e.ForEach("thead tr th", func(i int, thead *colly.HTMLElement) {
				fmt.Println("thead", i, thead.Text)
				headers = append(headers, thead.Text)
			})
			headers[3] = "趋势"

			data := make([][]string, 0)
			e.ForEach("tbody tr", func(i int, tr *colly.HTMLElement) {
				//fmt.Println("\ttr", i, tr.Text)
				trData := make([]string, 0)
				tr.ForEach("td", func(i int, td *colly.HTMLElement) {
					//fmt.Println("\t\ttd", i, td.Text)
					// 直接获取 td 下 span > svg 的 class 属性
					svgClass := td.ChildAttr("span svg", "class")
					//fmt.Println("第", i, "列 svg class:", svgClass)
					if "text-green-500" == svgClass {
						trData = append(trData, "上升")
					} else if "text-red-500" == svgClass {
						trData = append(trData, "下降")
					} else {
						trData = append(trData, td.Text)
					}
				})
				data = append(data, trData)
			})

			writeCsvFile("RankLanguagePopularity.csv", headers, data)

		}

	})

	c.Visit("https://hellogithub.com/en/report/tiobe")

}

/*
GOROOT=/Users/lxc20250729/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.3.darwin-arm64 #gosetup
GOPATH=/Users/lxc20250729/go #gosetup
/Users/lxc20250729/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.3.darwin-arm64/bin/go test -c -o /Users/lxc20250729/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand/___spider_02_colly_colly_03__TestMain01.test spider_02_colly/colly_03 #gosetup
/Users/lxc20250729/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.3.darwin-arm64/bin/go tool test2json -t /Users/lxc20250729/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand/___spider_02_colly_colly_03__TestMain01.test -test.v -test.paniconexit0 -test.run ^\QTestMain01\E$
=== RUN   TestMain01
aa &{table RankLanguagePopularity📊1Python25.98%2C++8.80%3C8.65%4Java8.35%5C#6.38%6JavaScript3.22%7Visual Basic2.84%8Go2.32%9Delphi/Object Pascal2.26%10Perl2.03%11SQL1.86%12Fortran1.49%13R1.43%14Ada1.27%15PHP1.25%16Scratch1.18%17Assembly language1.04%18Rust1.01%19MATLAB0.98%20Kotlin0.95% [{ class w-min min-w-full table-fixed divide-y-2 divide-gray-200 text-sm dark:divide-gray-700}] 0x14000237480 0x1400003e140 0x14000400b70 0}
thead 0 &{tr RankLanguagePopularity📊 [] 0x14000237480 0x1400003e140 0x14000400c30 0} RankLanguagePopularity📊
aa &{table RankLanguagePopularityChange🏅️Year1Python25.98%-0.16%2024, 20212C++8.80%-0.38%2022, 20033C8.65%-0.38%2019, 20174Java8.35%-0.24%2015, 20055C#6.38%0.86%20236JavaScript3.22%0.07%20147Visual Basic2.84%0.51%-8Go2.32%0.21%2016, 20099Delphi/Object Pascal2.26%0.44%-10Perl2.03%-0.05%-11SQL1.86%0.14%-12Fortran1.49%-0.26%-13R1.43%0.06%-14Ada1.27%-0.25%-15PHP1.25%-0.02%200416Scratch1.18%0.03%-17Assembly language1.04%0.01%-18Rust1.01%-0.12%-19MATLAB0.98%-0.21%-20Kotlin0.95%-0.15%- [{ class w-min min-w-full table-fixed divide-y-2 divide-gray-200 text-sm dark:divide-gray-700}] 0x14000237480 0x1400003e140 0x14000400c90 1}
thead 0 &{tr RankLanguagePopularityChange🏅️Year [] 0x14000237480 0x1400003e140 0x14000400d50 0} RankLanguagePopularityChange🏅️Year
--- PASS: TestMain01 (0.31s)
PASS

进程 已完成，退出代码为 0
*/
func TestMain0100(t *testing.T) {
	// 使用"encoding/csv"写数据的示例代码

	// 1. 创建文件
	csvFile, err := os.Create("csvtest.csv")
	if err != nil {
		fmt.Println("文件创建失败：", err)
		return
	}

	// 2. 创建 csv.Writer
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush() // 写完数据之后刷新一下文件

	// 3. 写入表头
	headers := []string{"id", "name", "age"}
	err = csvWriter.Write(headers)
	if err != nil {
		fmt.Println("表头写入失败：", err)
		return
	}

	// 4. 写入多行数据
	data := [][]string{
		{"1", "zhangsan1", "18"},
		{"2", "zhangsan2", "28"},
		{"3", "zhangsan3", "38"},
	}
	err = csvWriter.WriteAll(data)
	if err != nil {
		fmt.Println("表头数据失败：", err)
		return
	}
	fmt.Println("CSV 文件写入完成！")
}

func TestMain01000(t *testing.T) {

	headers := []string{"id", "name", "age"}
	data := [][]string{
		{"1", "zhangsan1", "18"},
		{"2", "zhangsan2", "28"},
		{"3", "zhangsan3", "38"},
	}
	writeCsvFile("csvtest2.csv", headers, data)

}
func writeCsvFile(filename string, headers []string, data [][]string) error {

	// 1. 创建文件
	csvFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("文件创建失败：", err)
		return err
	}

	// 2. 创建 csv.Writer
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush() // 写完数据之后刷新一下文件

	// 3. 写入表头
	//headers := []string{"id", "name", "age"}
	err = csvWriter.Write(headers)
	if err != nil {
		fmt.Println("表头写入失败：", err)
		return err
	}

	// 4. 写入多行数据
	//data := [][]string{
	//	{"1", "zhangsan1", "18"},
	//	{"2", "zhangsan2", "28"},
	//	{"3", "zhangsan3", "38"},
	//}
	err = csvWriter.WriteAll(data)
	if err != nil {
		fmt.Println("表头数据失败：", err)
		return err
	}
	fmt.Println("CSV 文件写入完成！")

	return nil
}
