package sizhi

import (
	"bytes"
	"encoding/json"
	"github.com/anjude/terminalx/utils/http_util"
)

type Req struct {
	Spoken string `json:"spoken"`
	Appid  string `json:"appid"`
	Userid string `json:"userid"`
}
type Resp struct {
	Message string `json:"message"`
	Data    struct {
		Type int `json:"type"`
		Info struct {
			Text string `json:"text"`
		} `json:"info"`
	} `json:"data"`
}

// GetSizhiMsg https://api.ownthink.com/bot
func GetSizhiMsg(content, userId string) (string, error) {
	botMsg := Req{
		Spoken: content,
		Appid:  "9ca3ec3b6b541c2f3285852986ce3d46",
		Userid: userId,
	}
	bytesData, _ := json.Marshal(botMsg)
	rspBody, err := http_util.DefaultClient.DoReq("POST", "https://api.ownthink.com/bot", bytes.NewReader(bytesData), nil)
	if err != nil {
		return "", err
	}
	ret := Resp{}
	err1 := json.Unmarshal(rspBody, &ret)
	if err1 != nil {
		return "", err
	}
	return ret.Data.Info.Text, nil
}
