package getui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Account) PushResult(taskList []string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	request, err := request(http.MethodPost, fmt.Sprintf(PUSHRESULT, c.AppID), headers, nil)
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

func (c Account) QueryAppUser(date string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	request, err := request(http.MethodGet, fmt.Sprintf(QUERYAPPUSER, c.AppID, date), headers, nil)
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

func (c Account) QueryAppPush(date string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	request, err := request(http.MethodGet, fmt.Sprintf(QUERYAPPPUSH, c.AppID, date), headers, nil)
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

func (c Account) PushResultQueryByTask(tasks []string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	body := map[string][]string{
		"taskIdList": tasks,
	}
	request, err := request(http.MethodPost, fmt.Sprintf(PUSHRESULTBYTASK, c.AppID), headers, body)
	var result map[string]interface{}
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
