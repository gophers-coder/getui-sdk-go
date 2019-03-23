package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin/json"

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

func ExampleAccount() {
	accountExample := getui.Account{
		AppID:        "1",
		AppKey:       "2",
		AppSecret:    "3",
		MasterSecret: "4",
	}
	fmt.Println(accountExample)
}

func ExampleAuthSign() {
	account.AuthSign()
}

func ExampleAuthClose() {
	account.AuthClose()
}

func ExampleBindAlias(cid string, alias string) {
	fmt.Println(account.BindAlias(cid, alias))
}
func ExampleQueryCid(cid string) {
	fmt.Println(account.QueryAlias(cid))
}
func ExampleUnbindAlias(cid, alias string) {
	fmt.Println(account.UnbindAlias(cid, alias))
}

func ExampleSetTags(cid string, tags []string) {
	fmt.Println(account.SetTags(cid, tags))
}

func ExampleGetTags(cid string) {
	fmt.Println(account.GetTags(cid))
}

func ExampleUserStatus(cid string) {
	fmt.Println(account.UserStatus(cid))
}

func ExamplePushSingle(cid string) {
	var params getui.PushSingleParams
	params = getui.PushSingleParams{
		Title:               "你好",
		Text:                "今天是2019.03.23",
		TransmissionContent: "好日子",
		Cid:                 []string{cid},
		AppKey:              account.AppKey,
	}
	fmt.Println(account.PushSingle(params))
}

func ExampleSaveListAndPush(cid string) {
	//"7f253f9854b1acf37264436f8deb3e37"
	var params getui.PushSingleParams
	params.Cid = []string{cid}
	params.Title = "今天是个好日子"
	params.Text = fmt.Sprintf("时间是: %s", time.Now().Format(time.ANSIC))

	content := map[string]string{}
	content["Age"] = "20"
	content["School"] = "ShangHai"
	bytes, _ := json.Marshal(content)
	params.TransmissionContent = string(bytes)
	fmt.Println(account.SaveListBodyAndPushList(params))
}

func main() {
	ExampleUserStatus("7f253f9854b1acf37264436f8deb3e37")
	ExamplePushSingle("7f253f9854b1acf37264436f8deb3e37")
	ExampleSaveListAndPush("7f253f9854b1acf37264436f8deb3e37")

}
