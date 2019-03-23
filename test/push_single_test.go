package getui_test

import (
	"fmt"
	"testing"

	"github.com/gophers-coder/getui-sdk-go/getui"
)

func TestPushSingle(tests *testing.T) {
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
		fmt.Println(account)
		fmt.Println(account.PushSingle(t))
	}

}
