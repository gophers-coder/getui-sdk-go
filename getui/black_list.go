package getui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Account) AddUserBlackList(cid string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	body := make(map[string][]string)
	body["cid"] = []string{cid}
	request, err := request(http.MethodPost, fmt.Sprintf(USERBLACKLIST, c.AppID), headers, body)
	var result = make(map[string]interface{})
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

func (c Account) RemoveUserBlackList(cid string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	body := make(map[string][]string)
	body["cid"] = []string{cid}
	request, err := request(http.MethodDelete, fmt.Sprintf(USERBLACKLIST, c.AppID), headers, body)
	var result = make(map[string]interface{})
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
