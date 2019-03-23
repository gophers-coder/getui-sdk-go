package getui_test

import (
	"fmt"
	"testing"
)

func TestUserBlackList(tests *testing.T) {
	fmt.Println(account.AddUserBlackList("7f253f9854b1acf37264436f8deb3e37"))
}
func TestRemoveUserBlackList(tests *testing.T) {
	fmt.Println(account.RemoveUserBlackList("7f253f9854b1acf37264436f8deb3e37"))
}
