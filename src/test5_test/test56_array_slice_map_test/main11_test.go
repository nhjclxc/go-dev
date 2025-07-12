package main

import (
	"fmt"
	"testing"
)

func TestDedupByID(t *testing.T) {

	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 1, Name: "Charlie"},
		{ID: 3, Name: "David"},
		{ID: 2, Name: "Eve"},
	}

	result := DedupByID(users)

	for _, u := range result {
		fmt.Println(u)
	}
}

func TestPaginate(t *testing.T) {
	users := []User{
		{ID: 1, Name: "Alice1"},
		{ID: 2, Name: "Bob2"},
		{ID: 3, Name: "Charlie3"},
		{ID: 4, Name: "David4"},
		{ID: 5, Name: "Eve5"},
		{ID: 6, Name: "Eve6"},
		{ID: 7, Name: "Eve7"},
		{ID: 8, Name: "Eve8"},
		{ID: 9, Name: "Eve9"},
		{ID: 10, Name: "Eve10"},
	}

	fmt.Println(Paginate[User](users, 1, 3))
	fmt.Println(Paginate[User](users, 2, 3))
	fmt.Println(Paginate[User](users, 3, 3))
	fmt.Println(Paginate[User](users, 4, 3))
	fmt.Println(Paginate[User](users, 5, 3))

}
