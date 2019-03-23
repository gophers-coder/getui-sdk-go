package getui_test

import (
	"fmt"
	"testing"

	"github.com/gophers-coder/getui-sdk-go/getui"
)

func TestPushResultQueryAppUser(tests *testing.T) {
	fmt.Println(account.QueryAppUser("2019-03-22"))
}

func TestPushResultQueryAppPush(tests *testing.T) {
	fmt.Println(account.QueryAppPush("2019-03-23"))
}

func TestPushResultQueryByTasks(tests *testing.T) {
	fmt.Println(account.PushResultQueryByTask([]string{"RASS_0323_bb79a456d0941696b8434b46bff829c7"}))
}

func TestSaveListBodyAndPushList(tests *testing.T) {
	var tt = []getui.PushSingleParams{
		getui.PushSingleParams{
			Title:               "你好",
			Text:                "今天是2019.03.23",
			TransmissionContent: "好日子",
			Cid:                 []string{"7f253f9854b1acf37264436f8deb3e37"},
			AppKey:              account.AppKey,
		},
	}
	for _, t := range tt {
		fmt.Println(account.SaveListBodyAndPushList(t))

	}
}
