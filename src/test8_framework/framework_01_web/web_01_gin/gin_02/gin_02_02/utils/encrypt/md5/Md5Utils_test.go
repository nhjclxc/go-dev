package md5

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {

	password := "abc1231"

	data := []byte(password)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	fmt.Println(md5str)

	data2 := []byte(password)
	has2 := md5.Sum(data2)
	md5str2 := fmt.Sprintf("%x", has2)
	fmt.Println(md5str2)
	fmt.Println(md5str == md5str2)

}
