package getui

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"spark-url/utils/json"
	"strconv"
	"time"
)

func request(method string, url string, header http.Header, body interface{}) (*http.Request, error) {
	content, err := json.Marshal(body)
	if err != nil {
		return &http.Request{}, fmt.Errorf("json.Marshal fail")
	}

	var request *http.Request

	if body != nil {
		request, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(content)))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return &http.Request{}, err
	}
	request.Header = header
	return request, nil
}

// HttpClientDo
func HttpClientDo(r *http.Request) ([]byte, error) {

	response, err := http.DefaultClient.Do(r)
	if err != nil {
		return []byte{}, fmt.Errorf("http.DefaultClietn.Do fail")
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

// timeStamp
func timeStamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10) + "000"
}

// sign
func sign(key string, timestamp string, masterSecret string) string {
	h := sha256.New()
	h.Write([]byte(key + timestamp + masterSecret))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func valueOrEnv(value string, index int) string {
	if index > 4 {
		return "-1"
	}
	var envMap = make(map[int]string)
	envMap = map[int]string{
		0: "APPID",
		1: "APPSECRET",
		2: "APPKEY",
		3: "MASTERSECRET",
	}
	if value == "" {
		return os.Getenv(envMap[index])
	}
	return value
}