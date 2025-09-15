package pkg11_github

import (
	"fmt"
	"github.com/samber/lo"
	"testing"
)

// https://github.com/samber/lo包的练习

type User struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
	Sex        string `json:"sex"`
}

var u *User
var userList []*User

func init() {
	u = &User{
		UserId:     11,
		UserPwd:    "123",
		UserName:   "zhangsan",
		UserStatus: 1,
		Sex:        "2",
	}

	userList = append(userList, &User{
		UserId:     1,
		UserPwd:    "11",
		UserName:   "zhangsan1",
		UserStatus: 2,
		Sex:        "2",
	})
	userList = append(userList, &User{
		UserId:     2,
		UserPwd:    "2",
		UserName:   "zhangsan22",
		UserStatus: 1,
		Sex:        "1",
	})
	userList = append(userList, &User{
		UserId:     3,
		UserPwd:    "3",
		UserName:   "zhangsan33",
		UserStatus: 3,
		Sex:        "2",
	})
	userList = append(userList, &User{
		UserId:     5,
		UserPwd:    "5",
		UserName:   "zhangsan-55",
		UserStatus: 5,
		Sex:        "1",
	})
	userList = append(userList, &User{
		UserId:     6,
		UserPwd:    "6",
		UserName:   "zhangsan-55",
		UserStatus: 3,
		Sex:        "2",
	})
}

func Test01_01(t *testing.T) {
	// 将某个结构体的某个属性提取成[]string

	idList := lo.Map(userList, func(item *User, _ int) int {
		return item.UserId
	})
	usernameList := lo.Map(userList, func(item *User, _ int) string {
		return item.UserName
	})

	fmt.Println("idList", idList)
	fmt.Println("usernameList", usernameList)
}

func Test01_02(t *testing.T) {
	// filter

	filter := lo.Filter(userList, func(item *User, index int) bool {
		fmt.Printf("%#v - %d \n", item, index)
		return true
	})
	fmt.Printf("filter %#v  \n\n", filter)

	users := lo.Filter(userList, func(item *User, index int) bool {
		fmt.Printf("%#v - %d \n", item, index)
		return false
	})
	fmt.Printf("users %#v  \n\n", users)

	users2 := lo.Filter(userList, func(item *User, index int) bool {
		return item.UserId > 3
	})
	fmt.Printf("users2 %#v  \n\n", users2)
	fmt.Printf("users2[0] %#v  \n\n", users2[0])
	fmt.Printf("users2[1] %#v  \n\n", users2[1])

}

func Test01_03(t *testing.T) {
	// forEach

	lo.ForEach(userList, func(item *User, index int) {
		fmt.Printf("i = %d , val = %#v \n", index, item)
	})
}

func Test01_05(t *testing.T) {
	// sum

	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("sum = ", lo.Sum(ints))

	by := lo.SumBy(userList, func(item *User) int {
		return item.UserId
	})

	fmt.Printf("sum = %d \n", by)

}

func Test01_06(t *testing.T) {
	// max min

	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("max = ", lo.Max(ints))
	fmt.Println("Min = ", lo.Min(ints))

	max1 := lo.MaxBy(userList, func(a *User, b *User) bool {
		return a.UserId > b.UserId
	})
	fmt.Printf("max1 = %v \n", max1)
	max2 := lo.MaxBy(userList, func(a *User, b *User) bool {
		return a.UserId < b.UserId
	})
	fmt.Printf("max2 = %v \n", max2)

	min1 := lo.MinBy(userList, func(a *User, b *User) bool {
		return a.UserId > b.UserId
	})
	fmt.Printf("min1 = %v \n", min1)
	min2 := lo.MinBy(userList, func(a *User, b *User) bool {
		return a.UserId < b.UserId
	})
	fmt.Printf("min2 = %v \n", min2)

}

func Test01_08(t *testing.T) {
	// count average

	ints := []int{1, 2, 3, 2, 5, 6, 2, 8, 9}
	fmt.Println("max = ", lo.Count(ints, 2))  // 统计ints里面等于2的元素个数
	fmt.Println("max = ", lo.Count(ints, 22)) //
}

func Test01_09(t *testing.T) {
	// 查找
	/*
		| Java Stream / List | 功能             | `samber/lo` 对应                                   |
		| ------------------ | ---------       | ------------------------------------------------ |
		| `findFirst()`      | 找到第一个满足条件 | `lo.Find(slice, func(item, idx) bool) (T, bool)` |
		| `anyMatch()`       | 是否存在满足条件   | `lo.Some(slice, func(item, idx) bool) bool`      |
		| `allMatch()`       | 全部满足条件      | `lo.Every(slice, func(item, idx) bool) bool`     |
		| `noneMatch()`      | 全部不满足        | `lo.None(slice, func(item, idx) bool) bool`      |
		| `contains()`       | 包含元素         | `lo.Contains(slice, item)`                       |
	*/
	ints := []int{1, 2, 3, 2, 5, 6, 2, 8, 9}
	find, b := lo.Find(ints, func(item int) bool {
		//return item == 5
		return item == 51
	})
	fmt.Printf("find = %v, b = %v \n", find, b)

	fmt.Println(" 8 ", lo.Contains(ints, 8))
	fmt.Println(" 18 ", lo.Contains(ints, 18))

	by := lo.ContainsBy(userList, func(item *User) bool {
		return item.UserName == "zhangsan1"
	})
	fmt.Printf("by = %v \n", by)

}

func Test01_10(t *testing.T) {
	// 转换 / 收集
	/*
		| Java Stream / Guava              | 功能       | `samber/lo` 对应                                 |
		| -------------------------------- | -------- | ---------------------------------------------- |
		| `collect(Collectors.toList())`   | 收集为 List | `lo.Map(...)` / `lo.Filter(...)` 返回 slice      |
		| `collect(Collectors.toMap(...))` | 收集为 Map  | `lo.KeyBy(slice, func(item) key) map[key]item` |
		| `Collectors.toSet()`             | 去重集合     | `lo.Uniq(slice)`                               |
		| `Lists.partition(list, n)`       | 切片成块     | `lo.Chunk(slice, n)`                           |
	*/
	strings := lo.Map(userList, func(item *User, index int) string {
		return item.UserName
	})
	fmt.Printf("strings = %#v \n", strings)

	by := lo.KeyBy(userList, func(item *User) string {
		return item.UserName // 当key重复时，使用新数据替换旧数据
	})
	fmt.Printf("by = %#v \n", by)

	ints := []int{1, 2, 3, 2, 5, 6, 2, 8, 9}
	fmt.Printf("Uniq = %#v \n\n", lo.Uniq(ints))

	uniqBy := lo.UniqBy(userList, func(item *User) string {
		return item.UserName
	})
	fmt.Printf("uniqBy = %#v \n\n", uniqBy)

}

func Test01_11(t *testing.T) {
	// 分组
	ints := []int{1, 2, 3, 2, 5, 6, 2, 8, 9}
	by := lo.GroupBy(ints, func(item int) int {
		return item
	})
	fmt.Printf("by = %#v \n\n", by)

	// map[item.UserName]*User
	by2 := lo.GroupBy(userList, func(item *User) string {
		return item.UserName
	})
	fmt.Printf("by2 = %#v \n\n", by2)

}
