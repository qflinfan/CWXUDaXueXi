package Request

import "Daxuexi/Config"

func Task() {
	cookie, authorization, member_id := Config.ReadConfig()
	lessonid := GetLesson(cookie)
	pngid := DoLesson(lessonid, cookie)
	Download(pngid)
	XCosSecurityToken, Authorization := GetAuthentication(authorization)
	filepath := SubmitPng(XCosSecurityToken, Authorization)
	id, create_at, invest := GetInfo(authorization, member_id)
	SubmitData(id, create_at, filepath, invest, member_id, authorization)
}
