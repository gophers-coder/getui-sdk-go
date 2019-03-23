package main

import (
	"os"

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
func main() {

}
