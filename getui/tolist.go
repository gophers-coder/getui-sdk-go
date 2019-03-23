package getui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Account) SaveListBodyAndPushList(value PushSingleParams) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	value.AppKey = c.AppKey
	body := NewSaveListBodyContent(value)

	var result = make(map[string]interface{})
	outputError := func(err error) map[string]interface{} {
		log.Panic(err)
		return result
	}
	requestSaveListBody, err := request(http.MethodPost, fmt.Sprintf(SAVELISTBODY, c.AppID), headers, body)
	if err != nil {
		outputError(err)
	}
	content, err := HttpClientDo(requestSaveListBody)
	if err != nil {
		outputError(err)
	}
	json.Unmarshal(content, &result)
	if result["result"] != "ok" {
		return result
	}

	pushListBody := map[string]interface{}{
		"cid":         value.Cid,
		"taskid":      result["taskid"].(string),
		"need_detail": true,
	}

	requestPushList, err := request(http.MethodPost, fmt.Sprintf(PUSHLIST, c.AppID), headers, pushListBody)
	if err != nil {
		outputError(err)
	}
	contentForPush, err := HttpClientDo(requestPushList)
	if err != nil {
		outputError(err)
	}
	var resultForPush = make(map[string]interface{})
	json.Unmarshal(contentForPush, &resultForPush)
	return resultForPush
}
