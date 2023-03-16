package Request

import (
	"github.com/imroc/req/v3"
)

// InfoData 定义一个结构体来存储json数据中的data字段
type InfoData struct {
	ID       string `json:"_id"`
	CreateAt string `json:"create_at"`
	Invest   string `json:"invest"`
}

// Response3 定义一个结构体来存储整个json数据
type Response3 struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data []InfoData `json:"data"`
}

// AuthenticationData 定义一个 AuthenticationData 结构体，包含四个字段
type AuthenticationData struct {
	ExpiredTime       int    `json:"expiredTime"`
	RequestId         string `json:"requestId"`
	XCosSecurityToken string `json:"XCosSecurityToken"`
	Authorization     string `json:"Authorization"`
}

// AuthenticationResponse 定义一个 AuthenticationResponse 结构体，包含三个字段
type AuthenticationResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data AuthenticationData `json:"data"` // 这里使用了 Data 结构体作为字段类型
}

func GetAuthentication(authorization string) (string, string) {
	var response AuthenticationResponse
	client := req.C()
	body := map[string]bool{"stsGray": true}
	_, _ = client.R().
		SetHeaders(map[string]string{
			"Authorization":   authorization,
			"Referer":         "https://servicewechat.com/wx23d8d7ea22039466/1908/page-frame.html",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
		SetBody(body).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/api/appGetCosSTS")
	return response.Data.XCosSecurityToken, response.Data.Authorization
}

func GetInfo(authorization string, member_id string) (string, string, string) {
	var response Response3
	client := req.C()
	_, _ = client.R().
		SetHeaders(map[string]string{
			"Authorization":   authorization,
			"Referer":         "https://servicewechat.com/wx23d8d7ea22039466/1908/page-frame.html",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
		SetPathParams(map[string]string{
			"members": member_id,
		}).
		SetSuccessResult(&response).
		Get("https://a.welife001.com/info/getParent?type=-1&members={members}&page=0&size=10&date=-1&hasMore=true")
	return response.Data[0].ID, response.Data[0].CreateAt, response.Data[0].Invest
}
