package getui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Account) SetTags(cid string, tags []string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	body := Tags{
		Cid:     cid,
		TagList: tags,
	}
	request, err := request(http.MethodPost, fmt.Sprintf(SETTAGS, c.AppID), headers, body)
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

func (c Account) GetTags(cid string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	request, err := request(http.MethodGet, fmt.Sprintf(GETTAGS, c.AppID, cid), headers, nil)
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
