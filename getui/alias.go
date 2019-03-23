package getui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Account) BindAlias(cid string, alias string) map[string]string {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	body := AliasParams{
		[]OneCidAlias{
			{
				Cid:   cid,
				Alias: alias,
			},
		},
	}
	request, err := request(http.MethodPost, fmt.Sprintf(BINDALIAS, c.AppID), headers, body)
	var result = make(map[string]string)
	if err != nil {
		log.Panic(err)
		return result
	}
	content, err := HttpClientDo(request)
	json.Unmarshal(content, &result)
	return result
}

func (c Account) QueryCid(alias string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	request, err := request(http.MethodGet, fmt.Sprintf(QUERYCID, c.AppID, alias), headers, nil)
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

func (c Account) QueryAlias(cid string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	request, err := request(http.MethodGet, fmt.Sprintf(QUERYALIAS, c.AppID, cid), headers, nil)
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

func (c Account) UnbindAlias(cid string, alias string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	body := OneCidAlias{
		Cid:   cid,
		Alias: alias,
	}
	request, err := request(http.MethodPost, fmt.Sprintf(UNBINDALIAS, c.AppID), headers, body)
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

func (c Account) UnbindAliasAll(alias string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	body := make(map[string]string)
	body["alias"] = alias
	request, err := request(http.MethodPost, fmt.Sprintf(UNBINDALLALIAS, c.AppID), headers, body)
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
