package getui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ConfigAccount struct {
	AppID        string
	AppSecret    string
	AppKey       string
	MasterSecret string
}

func NewGetUiAccount(AppID string, AppSecret string, AppKey string, MasterSecret string) *ConfigAccount {
	var (
		id           string
		secret       string
		key          string
		masterSecret string
	)
	id = valueOrEnv(AppID, 0)
	secret = valueOrEnv(AppSecret, 1)
	key = valueOrEnv(AppKey, 2)
	masterSecret = valueOrEnv(MasterSecret, 3)
	return &ConfigAccount{
		AppID:        id,
		AppSecret:    secret,
		AppKey:       key,
		MasterSecret: masterSecret,
	}
}

func (c ConfigAccount) AuthSign() map[string]string {
	timestamp := timeStamp()
	signString := sign(c.AppKey, timestamp, c.MasterSecret)

	body := map[string]string{
		"sign":      signString,
		"timestamp": timestamp,
		"appkey":    c.AppKey,
	}

	headers := http.Header{
		"Content-Type": []string{"application/json"},
	}
	request, _ := request("POST", fmt.Sprintf(AUTHSIGN, c.AppID), headers, body)
	content, _ := HttpClientDo(request)

	var result map[string]string
	json.Unmarshal(content, &result)
	return result

}

func (c ConfigAccount) AuthClose(auth string) map[string]string {
	headers := http.Header{
		"authtoken": []string{auth},
	}
	request, err := request(http.MethodPost, fmt.Sprintf(AUTHCLOSE, c.AppID), headers, nil)

	var result = make(map[string]string)
	if err != nil {
		log.Panic(err)
		return result
	}

	content, err := HttpClientDo(request)

	if err != nil {
		log.Panic(err)
		return result
	}
	json.Unmarshal(content, &result)
	return result
}
