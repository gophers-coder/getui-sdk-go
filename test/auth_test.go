package getui_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gophers-coder/getui-sdk-go/getui"
)

var account *getui.Account

func init() {
	account = getui.NewGetUiAccount(
		os.Getenv("APPID"),
		os.Getenv("APPSECRET"),
		os.Getenv("APPKEY"),
		os.Getenv("MASTERSECRET"))
	account.AuthSign()
}

func TestAccount(t *testing.T) {

	fmt.Println(account.Auth)
}

func TestAuthClose(t *testing.T) {

	account.AuthSign()
	fmt.Println(account.AuthClose())

}
