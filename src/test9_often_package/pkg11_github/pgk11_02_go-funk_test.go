package pkg11_github

import (
	"fmt"
	"iter"
	"maps"
	"strings"
	"testing"
)

type data struct {
	name string
}

func Test02_15(t *testing.T) {

	m := map[string]data{"x": {"one"}}
	fmt.Println("d.name ", m["x"].name)
	//m["x"].name = "www" // 无法分配给 m["x"].name

	d := m["x"]
	fmt.Println("d ", d)
	d.name = "zzzzz"
	fmt.Println("d ", d)

	str := "gormgogen"
	pre := "gorm"

	// strings.TrimLeft 并不是“去掉前缀”，而是 去掉字符串开头处所有在 cutset 里的任意字符。
	fmt.Println("result = ", strings.TrimLeft(str, pre))   // 结果是en，而不是gogen
	fmt.Println("result = ", strings.TrimPrefix(str, pre)) // gogen

}

func Test02_16(t *testing.T) {

	m := make(map[string]struct{})
	for i := 0; i < 5; i++ {
		m[fmt.Sprintf("x-%d", i+5)] = struct{}{}
	}

	var keyList iter.Seq[string] = maps.Keys(m)
	fmt.Println(keyList)

	ks2 := Keys(m)
	fmt.Println(ks2)

}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
