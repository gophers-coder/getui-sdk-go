package getui_test

import (
	"fmt"
	"testing"
)

// 绑定别名
func TestBindAlias(tests *testing.T) {
	fmt.Println(account.BindAlias("7f253f9854b1acf37264436f8deb3e37", "xiewei_2019"))
}

// 通过别名找 cid
func TestQueryAliasCid(tests *testing.T) {
	fmt.Println(account.QueryCid("xiewei_2019"))
}

// 通过 cid 找别名
func TestQueryCidAlias(tests *testing.T) {
	fmt.Println(account.QueryAlias("7f253f9854b1acf37264436f8deb3e37"))
}

// 解除cid 绑定的别名
func TestUnbindAlias(tests *testing.T) {
	fmt.Println(account.UnbindAlias("7f253f9854b1acf37264436f8deb3e37", "xiewei_2019"))
}

// 解除所有的别名的绑定
func TestUnbindAliasAll(tests *testing.T) {
	fmt.Println(account.UnbindAliasAll("xiewei_2019"))
}
