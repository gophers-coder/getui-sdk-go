package getui_test

import (
	"fmt"
	"testing"
)

func TestUserStatus(tests *testing.T) {
	fmt.Println(account.UserStatus("7f253f9854b1acf37264436f8deb3e37"))
}

func TestOnlineUser(tests *testing.T) {
	fmt.Println(account.UserOnlineStatics())
}
