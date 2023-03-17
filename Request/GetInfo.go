package Request

import (
	"fmt"
	"github.com/imroc/req/v3"
)

type AuthenticationData struct {
	XCosSecurityToken string `json:"XCosSecurityToken"`
	Authorization     string `json:"Authorization"`
}

type AuthenticationResponse struct {
	Data AuthenticationData `json:"data"`
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

type InfoResponse struct {
	Data []InfoData `json:"data"`
}

type InfoData struct {
	ID       string `json:"_id"`
	CreateAt string `json:"create_at"`
}

func GetInfo(authorization string, member_id string) (string, string) {
	var response InfoResponse
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
	return response.Data[0].ID, response.Data[0].CreateAt
}

type InfoResponseStudent struct {
	Data InfoDataStudent `json:"data"`
}

type InfoDataStudent struct {
	Notify Notify `json:"notify"`
}

type Notify struct {
	Invest Invest `json:"invest"`
}

type Invest struct {
	ID      string    `json:"_id"`
	Subject []Subject `json:"subject"`
}

type Subject struct {
	ID          string        `json:"_id"`
	ItemDetails []ItemDetails `json:"item_details"`
	CreateAt    string        `json:"create_at"`
}

type ItemDetails struct {
	ID string `json:"_id"`
}

func GetInfoStudent(authorization string, id string) (string, string, string, string) {
	var response InfoResponseStudent
	client := req.C()
	_, _ = client.R().
		SetHeaders(map[string]string{
			"Authorization":   authorization,
			"Referer":         "https://servicewechat.com/wx23d8d7ea22039466/1908/page-frame.html",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
		SetBodyJsonMarshal(map[string]string{
			"extra": "1",
			"_id":   id,
			"page":  "0",
			"size":  "10",
		}).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/notify/checkNew2Parent")
	investid := response.Data.Notify.Invest.ID
	subjectid := response.Data.Notify.Invest.Subject[0].ID
	itemDetailstudentid1 := response.Data.Notify.Invest.Subject[0].ItemDetails[0].ID
	itemDetailstudentid2 := response.Data.Notify.Invest.Subject[0].ItemDetails[1].ID
	fmt.Println(investid, subjectid, itemDetailstudentid1, itemDetailstudentid2)
	return investid, subjectid, itemDetailstudentid1, itemDetailstudentid2
}

//以下代码在accept数组有数据时可用
//type InfoResponseTeacher struct {
//	Data InfoDataTeacher `json:"data"`
//}
//
//type InfoDataTeacher struct {
//	Accept Accept `json:"accept"`
//}
//
//type Accept struct {
//	Answer Answer `json:"answer"`
//}
//
//type Answer struct {
//	ID             string           `json:"_id"`
//	SubjectTeacher []SubjectTeacher `json:"subject"`
//}
//
//type SubjectTeacher struct {
//	ID          string        `json:"_id"`
//	ItemDetails []ItemDetails `json:"item_details"`
//	CreateAt    string        `json:"create_at"`
//}
//
//type ItemDetailsTeacher struct {
//	ID string `json:"_id"`
//}

//func GetInfoTeacher(authorization string, id string) (string, string, string, string) {
//	var response InfoResponseTeacher
//	client := req.C()
//	_, _ = client.R().
//		SetHeaders(map[string]string{
//			"Authorization":   authorization,
//			"Referer":         "https://servicewechat.com/wx23d8d7ea22039466/1908/page-frame.html",
//			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
//			"Content-Type":    "application/json",
//			"Accept-Language": "zh-CN,zh",
//		}).
//		SetBodyJsonMarshal(map[string]string{
//			"extra": "1",
//			"_id":   id,
//			"page":  "0",
//			"size":  "10",
//		}).
//		SetSuccessResult(&response).
//		Post("https://a.welife001.com/applet/notify/checkNew2Parent")
//	answerid := response.Data.Accept.Answer.ID
//	subjecteacherid := response.Data.Accept.Answer.SubjectTeacher[0].ID
//	itemDetailsteacherid1 := response.Data.Accept.Answer.SubjectTeacher[0].ItemDetails[0].ID
//	itemDetailsteacherid2 := response.Data.Accept.Answer.SubjectTeacher[0].ItemDetails[1].ID
//	return answerid, subjecteacherid, itemDetailsteacherid1, itemDetailsteacherid2
//}
