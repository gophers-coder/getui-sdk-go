package getui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Account) UserStatus(cid string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	request, err := request(http.MethodGet, fmt.Sprintf(USETSTATUS, c.AppID, cid), headers, nil)
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

func (c Account) UserOnlineStatics() map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	request, err := request(http.MethodGet, fmt.Sprintf(ONLINEUSER, c.AppID), headers, nil)
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
