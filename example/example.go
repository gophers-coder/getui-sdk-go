package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gophers-coder/getui-sdk-go/getui"
)

func main() {
	account := getui.NewGetUiAccount(
		os.Getenv("APPID"),
		os.Getenv("APPSECRET"),
		os.Getenv("APPKEY"),
		os.Getenv("MASTERSECRET"))
	account.AuthSign()
	var params getui.PushSingleParams
	params.Cid = []string{"7f253f9854b1acf37264436f8deb3e37"}
	params.Title = "今天是个好日子"
	params.Text = fmt.Sprintf("时间是: %s", time.Now().Format(time.ANSIC))
	params.TransmissionContent = "消息"
	fmt.Println(account.SaveListBodyAndPushList(params))
}
