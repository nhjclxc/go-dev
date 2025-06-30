package main

import (
	"fmt"
	"sort"
)

func main02() {


	names := []string{"Bob", "Alice", "David"}
	sort.Strings(names)
	fmt.Println(names) // [Alice Bob David]

}
