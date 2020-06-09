package common

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

// 百度Token
type BaiduToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// Token字符串
var BaiduAccessToken string

func GetBaiduToken() {
	BaiduAccessToken := GetString("Application_Baidu_Api_Token")
	if BaiduAccessToken == "" {
		res, err := http.Post(beego.AppConfig.String("baidu_api_access_token"), "application/json", strings.NewReader(""))
		if err != nil {
			beego.Debug(err)
		} else {
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				beego.Error(err.Error())
			} else {
				var baiduToken BaiduToken
				err := json.Unmarshal(body, &baiduToken)
				if err != nil {
					beego.Error(err.Error())
				} else {
					beego.Debug(baiduToken)
					BaiduAccessToken = baiduToken.AccessToken
					Put("Application_Baidu_Api_Token", BaiduAccessToken)
					Expire("Application_Baidu_Api_Token", baiduToken.ExpiresIn)
				}
			}
		}
	}
}
