package main

import (
	"fmt"
	"testing"
)

func TestGroupOrdersByUser(t *testing.T) {

	users := []Order{
		{ID: 1, Name: "Alice1"},
		{ID: 2, Name: "Bob2"},
		{ID: 1, Name: "Charlie3"},
		{ID: 1, Name: "David4"},
		{ID: 5, Name: "Eve5"},
		{ID: 2, Name: "Eve6"},
		{ID: 3, Name: "Eve7"},
		{ID: 8, Name: "Eve8"},
		{ID: 9, Name: "Eve9"},
		{ID: 5, Name: "Eve10"},
	}

	fmt.Println(GroupOrdersByUser(users))


}
