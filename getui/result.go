package getui

// ResultForAuthSign
type ResultForAuthSign struct {
	Result    string `json:"result"`
	AuthToken string `json:"auth_token"`
	Expire    string `json:"expire_time"`
}
