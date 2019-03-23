package getui

type PushSingleParams struct {
	AppKey              string   `json:"appkey"`
	Text                string   `json:"text"`
	Title               string   `json:"title"`
	TransmissionContent string   `json:"transmission_content"`
	Cid                 []string `json:"cid"`
}

type PushSingleContent struct {
	Message      Message      `json:"message"`
	Notification Notification `json:"notification"`
	Cid          string       `json:"cid"`
	RequestID    string       `json:"requestid"`
}

type Message struct {
	AppKey            string `json:"appkey"`
	IsOffline         bool   `json:"is_offline"`
	OfflineExpireTime int    `json:"offline_expire_time"`
	MsgType           string `json:"msgtype"`
}

type Notification struct {
	Style               Style  `json:"style"`
	TransmissionType    bool   `json:"transmission_type"`
	TransmissionContent string `json:"transmission_content"`
}

func NewPushSingleContent(params PushSingleParams) PushSingleContent {
	return PushSingleContent{
		Message: Message{params.AppKey, true, 10000000, "notification"},
		Notification: Notification{
			Style:               Style{0, params.Text, params.Title},
			TransmissionType:    true,
			TransmissionContent: params.TransmissionContent,
		},
		Cid:       params.Cid[0],
		RequestID: timeStamp(),
	}
}

type Style struct {
	Type  int    `json:"type"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

var PushSingleTemplate = `
{
	"message": {
		"appkey": {{.AppKey}},
		"is_offline": true,
		"offline_expire_time":10000000,
		"msgtype": "notification"
	}
	"notification": {
		"style": {
        	"type": 0,
            "text": {{.Text}},
            "title": {{.Title}}
		},
        "transmission_type": true,
        "transmission_content": {{.TransmissionContent}}
	},
    "cid": {{.Cid}}
}
`

type SaveListBodyContent struct {
	Message      Message      `json:"message"`
	Notification Notification `json:"notification"`
}

func NewSaveListBodyContent(params PushSingleParams) SaveListBodyContent {
	return SaveListBodyContent{
		Message: Message{params.AppKey, true, 10000000, "notification"},
		Notification: Notification{
			Style:               Style{0, params.Text, params.Title},
			TransmissionType:    true,
			TransmissionContent: params.TransmissionContent,
		},
	}
}

func SaveListBodyContentByMap(params map[string]string) SaveListBodyContent {
	return SaveListBodyContent{
		Message: Message{params["appkey"], true, 10000000, "notification"},
		Notification: Notification{
			Style:               Style{0, params["text"], params["title"]},
			TransmissionType:    true,
			TransmissionContent: params["transmission_content"],
		},
	}
}

var SaveListBody = `
{
                   "message": {
                   "appkey": {{.AppKey}},
                   "is_offline": true,
                   "offline_expire_time":10000000,
                   "msgtype": "notification"
                },
                "notification": {
                    "style": {
                        "type": 0,
                        "text": {{.Text}},
                        "title": {{.Title}}
                    },
                    "transmission_type": true,
                    "transmission_content": {{.TransmissionContent}}
                }
           }
`
