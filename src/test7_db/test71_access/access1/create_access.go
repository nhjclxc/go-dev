package main

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func main2() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	accessApp, err := oleutil.CreateObject("Access.Application")
	if err != nil {
		panic(fmt.Sprintf("无法创建 Access 对象: %v", err))
	}

	accessDispatch, err := accessApp.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		panic(fmt.Sprintf("无法获取 IDispatch 接口: %v", err))
	}
	defer accessDispatch.Release()

	// 保存路径，确保目录存在
	//filePath := `D:\data\test.accdb`
	filePath := `D:\nbhkkj\odbc\test.accdb` // D:\nbhkkj\odbc\test.accdb
	_, err = oleutil.CallMethod(accessDispatch, "NewCurrentDatabase", filePath)
	if err != nil {
		panic(fmt.Sprintf("无法创建数据库文件: %v", err))
	}

	// 创建表
	createSQL := `
        CREATE TABLE my_table (
            id AUTOINCREMENT PRIMARY KEY,
            name TEXT(100),
            value TEXT(100)
        )
    `
	_, err = oleutil.CallMethod(accessDispatch, "DoCmd.RunSQL", createSQL)
	if err != nil {
		panic(fmt.Sprintf("建表失败: %v", err))
	}

	fmt.Println("✅ 创建数据库并建表成功：", filePath)

	// 关闭 Access
	oleutil.CallMethod(accessDispatch, "CloseCurrentDatabase")
	oleutil.CallMethod(accessDispatch, "Quit")
}
