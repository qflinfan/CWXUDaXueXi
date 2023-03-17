package Request

import "Daxuexi/Config"

func Task() {
	cookie, authorization, member_id, identity := Config.ReadConfig()
	lessonid := GetLesson(cookie)
	pngid := DoLesson(lessonid, cookie)
	Download(pngid)
	XCosSecurityToken, Authorization := GetAuthentication(authorization)
	key := generateRandomString(4)
	filepath := SubmitPng(key, XCosSecurityToken, Authorization)
	id, create_at := GetInfo(authorization, member_id)
	if identity == "student" {
		investid, subjectid, itemDetailstudentid1, itemDetailstudentid2 := GetInfoStudent(authorization, id)
		SubmitDataStudent(id, create_at, filepath, member_id, authorization, investid, subjectid, itemDetailstudentid1, itemDetailstudentid2)
	} else {
		//在accepts数组有数据是启用
		//answerid, subjecteacherid, itemDetailsteacherid1, itemDetailsteacherid2 := GetInfoTeacher(authorization, id)
		SubmitDataTeacher(id, create_at, filepath, member_id, authorization)
	}
}
