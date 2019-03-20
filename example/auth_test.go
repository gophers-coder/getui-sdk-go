package getui_test

import (
	"fmt"
	"testing"

	"github.com/gophers-coder/getui-sdk-go/getui"
)

var account *getui.ConfigAccount

func init() {
	account = getui.NewGetUiAccount(
		"hwWyn7WUU27rtz16xtfBK8",
		"srRR0JL6R5AnCkx4NG8FS1",
		"WOT9XN5i4j72vUa45nuqk8",
		"AJQwsryY5w9cDb50DJDBk4")
}

func TestAccount(t *testing.T) {
	result := account.AuthSign()
	fmt.Println(result)
}

func TestAuthClose(t *testing.T) {

	result := account.AuthSign()
	fmt.Println(account.AuthClose(result["auth_token"]))

}
