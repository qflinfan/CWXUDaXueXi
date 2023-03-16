package Request

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"regexp"
	"strconv"
)

// Lesson 定义一个Lesson结构体，表示data数组中的每个元素
type Lesson struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}

// Response 定义一个Response结构体，表示整个JSON数据
type Response struct {
	Message  string   `json:"message"`
	Status   int      `json:"status"`
	Redirect string   `json:"redirect"`
	Data     []Lesson `json:"data"`
}

func GetLesson(cookie string) string {
	var response Response
	client := req.C()
	lessondata, _ := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x6308011a) XWEB/6500",
			"Content-Type":    "application/json",
			"Origin":          "https://service.jiangsugqt.org",
			"Referer":         "https://service.jiangsugqt.org/youth-h5/",
			"Accept-Language": "zh-CN,zh",
			"Connection":      "close",
			"Cookie":          cookie,
		}).
		SetBodyJsonMarshal(map[string]string{
			"page":  "1",
			"limit": "5",
		}).
		SetSuccessResult(&response).
		Post("https://service.jiangsugqt.org/api/lessons")
	defer lessondata.Body.Close()
	return strconv.Itoa(response.Data[0].ID)
}

// LessonData Data 定义一个Data结构体，表示data对象
type LessonData struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Response2 Response 定义一个Response结构体，表示整个JSON数据
type Response2 struct {
	Message  string     `json:"message"`
	Status   int        `json:"status"`
	Redirect string     `json:"redirect"`
	Data     LessonData `json:"data"`
}

func DoLesson(lessonId string, cookie string) string {
	var response2 Response2
	client := req.C()
	lessondata, _ := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x6308011a) XWEB/6500",
			"Content-Type":    "application/json",
			"Origin":          "https://service.jiangsugqt.org",
			"Referer":         "https://service.jiangsugqt.org/youth-h5/",
			"Accept-Language": "zh-CN,zh",
			"Connection":      "close",
			"Cookie":          cookie,
		}).
		SetBodyJsonMarshal(map[string]string{
			"lesson_id": lessonId,
		}).
		Post("https://service.jiangsugqt.org//api/doLesson")
	lessondataString := lessondata.String()
	err := json.Unmarshal([]byte(lessondataString), &response2)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	pattern := "daxuexi/(.*?)/index.html"
	re, _ := regexp.Compile(pattern)
	match := re.FindStringSubmatch(response2.Data.URL)[1]
	return match
}
