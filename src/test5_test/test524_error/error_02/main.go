package main

import (
	"errors"
	"fmt"
)

// ✅ 三、错误链（Wrapping）

func readConfig() error {
	return errors.New("file not found")
}

func loadApp() error {
	err := readConfig()
	return fmt.Errorf("load app failed: %w", err)
}

func main() {
	// 这段代码中，readConfig() 返回的错误被 loadApp() 包装了，但可以通过 errors.Unwrap() 解开。
	err := loadApp()
	if err != nil {
		print(err.Error())

		err2 := errors.Unwrap(err)
		if err2 != nil {
			fmt.Println(err2.Error())
			err3 := errors.Unwrap(err2)
			if err3 != nil {
				fmt.Println(err3.Error())
				err5 := errors.Unwrap(err3)
				if err5 != nil {
					fmt.Println(err5.Error())
					return
				}
			}
			return
		}


		return
	}

}
