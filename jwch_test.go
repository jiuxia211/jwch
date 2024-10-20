package jwch

import (
	"fmt"
	"os"
	"testing"

	"github.com/west2-online/jwch/utils"
)

var (
	username  = os.Getenv("JWCH_USERNAME") // 学号
	password  = os.Getenv("JWCH_PASSWORD") // 密码
	localfile = "logindata.txt"
)

var (
	islogin bool     = false
	stu     *Student = NewStudent().WithUser(username, password)
)

func login() error {

	err := stu.Login()

	if err != nil {
		return err
	}

	err = stu.CheckSession()

	if err != nil {
		return err
	}

	islogin = true
	return nil
}

func Test_GetValidateCode(t *testing.T) {
	// 获取验证码图片
	s := NewStudent()
	resp, err := s.NewRequest().Get("https://jwcjwxt1.fzu.edu.cn/plus/verifycode.asp")
	if err != nil {
		t.Error(err)
	}
	code, err := GetValidateCode(utils.Base64EncodeHTTPImage(resp.Body()))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(code)
}

func Test_GetIdentifierAndCookies(t *testing.T) {
	Identifier, cookies := stu.GetIdentifierAndCookies()
	fmt.Println(Identifier)
	fmt.Println(cookies)
}

func Test_Login(t *testing.T) {
	err := login()
	if err != nil {
		t.Error(err)
	}
}

func Test_GetCourse(t *testing.T) {
	if !islogin {
		err := login()

		if err != nil {
			t.Error(err)
		}
	}

	terms, err := stu.GetTerms()

	if err != nil {
		t.Error(err)
	}

	list, err := stu.GetSemesterCourses(terms.Terms[0], terms.ViewState, terms.EventValidation)

	if err != nil {
		t.Error(err)
	}

	fmt.Println("course num:", len(list))

	for _, v := range list {
		fmt.Println(utils.PrintStruct(v))
	}
}

func Test_GetInfo(t *testing.T) {
	if !islogin {
		err := login()

		if err != nil {
			t.Error(err)
		}
	}

	detail, err := stu.GetInfo()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(detail))
}

func Test_GetMarks(t *testing.T) {
	if !islogin {
		err := login()

		if err != nil {
			t.Error(err)
		}
	}

	marks, err := stu.GetMarks()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(marks))
}

// 使用并发后似乎快了1s(
func Test_GetQiShanEmptyRoom(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}

	rooms, err := stu.GetQiShanEmptyRoom(EmptyRoomReq{
		Campus: "旗山校区",
		Time:   "2024-09-26",
		Start:  "1",
		End:    "8",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(rooms))
}

func Test_GetJinJiangEmptyRoom(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}

	rooms, err := stu.GetEmptyRoom(EmptyRoomReq{
		Campus: "晋江校区",
		Time:   "2024-09-19",
		Start:  "1",
		End:    "2",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(rooms))
}

func Test_GetTongPanEmptyRoom(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}

	rooms, err := stu.GetEmptyRoom(EmptyRoomReq{
		Campus: "铜盘校区",
		Time:   "2024-09-19",
		Start:  "1",
		End:    "2",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(rooms))
}

func Test_GetQuanGangEmptyRoom(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}

	rooms, err := stu.GetEmptyRoom(EmptyRoomReq{
		Campus: "泉港校区",
		Time:   "2024-09-19",
		Start:  "1",
		End:    "2",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(rooms))
}

func Test_GetYiShanEmptyRoom(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}

	rooms, err := stu.GetEmptyRoom(EmptyRoomReq{
		Campus: "怡山校区",
		Time:   "2024-09-19",
		Start:  "1",
		End:    "2",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(rooms))
}

func Test_GetXiaMenEmptyRoom(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}

	rooms, err := stu.GetEmptyRoom(EmptyRoomReq{
		Campus: "厦门工艺美院",
		Time:   "2024-09-19",
		Start:  "1",
		End:    "2",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(rooms))
}

func Test_GetSchoolCalendar(t *testing.T) {
	calendar, err := stu.GetSchoolCalendar()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(calendar))
}

func Test_GetTermEvents(t *testing.T) {
	calendar, err := stu.GetSchoolCalendar()
	if err != nil {
		t.Error(err)
	}

	events, err := stu.GetTermEvents(calendar.Terms[0].TermId)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(events))
}

func Test_GetCredit(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}

	credit, err := stu.GetCredit()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(credit))
}

func Test_GetGPA(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}
	gpa, err := stu.GetGPA()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(gpa))
}

func TestGetUnifiedExam(t *testing.T) {
	if !islogin {
		err := login()
		if err != nil {
			t.Error(err)
		}
	}
	cet, err := stu.GetCET()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(cet))

	js, err := stu.GetJS()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(js))
}
