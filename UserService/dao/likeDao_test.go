package dao

import (
	"fmt"
	"testing"
)

func TestInsertLike(t *testing.T) {
	InitDB()
	err := InsertLike(6, 4)
	fmt.Printf("%v", err)
}

func TestGetLike(t *testing.T) {
	InitDB()
	like, err := GetLike(4, 8)
	fmt.Printf("%v", like)
	fmt.Printf("%v", err)
}

func TestUpdateLike(t *testing.T) {
	InitDB()
	err := UpdateLike(4, 8, 0)
	fmt.Printf("%v", err)
}

func TestLikeList(t *testing.T) {
	InitDB()
	vidList, err := LikeList(4)
	fmt.Printf("%v", vidList)
	fmt.Printf("%v", err)
}

func TestLikeUserList(t *testing.T) {
	InitDB()
	uidList, err := LikeUserList(11)
	fmt.Printf("%v", uidList)
	fmt.Printf("%v", err)
}
