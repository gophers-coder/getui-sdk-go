package getui_test

import (
	"fmt"
	"testing"
)

// 设置标签
func TestSetTags(tests *testing.T) {
	fmt.Println(account.SetTags("7f253f9854b1acf37264436f8deb3e37", []string{"hello", "world"}))
}

// 获取标签
func TestGetTags(tests *testing.T) {
	fmt.Println(account.GetTags("7f253f9854b1acf37264436f8deb3e37"))
}
