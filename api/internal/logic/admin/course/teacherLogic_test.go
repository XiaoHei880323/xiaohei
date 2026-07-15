package course

import (
	"api/model"
	"encoding/json"
	"strings"
	"testing"
)

func TestValidateTeacher(t *testing.T) {
	status := 1
	tests := []struct {
		name         string
		teacherNo    string
		teacherName  string
		gender       int
		status       *int
		password     string
		passwordNeed bool
		want         string
	}{
		{name: "valid add", teacherNo: "T001", teacherName: "李老师", gender: 1, status: &status, password: "123456", passwordNeed: true},
		{name: "teacher number required", teacherName: "李老师", status: &status, password: "123456", passwordNeed: true, want: "请输入教师工号"},
		{name: "name required", teacherNo: "T001", status: &status, password: "123456", passwordNeed: true, want: "请输入教师姓名"},
		{name: "gender invalid", teacherNo: "T001", teacherName: "李老师", gender: 3, status: &status, password: "123456", passwordNeed: true, want: "性别参数错误"},
		{name: "password too short", teacherNo: "T001", teacherName: "李老师", status: &status, password: "123", passwordNeed: true, want: "密码不能少于6位"},
		{name: "empty update password", teacherNo: "T001", teacherName: "李老师", status: &status},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := validateTeacher(test.teacherNo, test.teacherName, test.gender, test.status, test.password, test.passwordNeed)
			if got != test.want {
				t.Fatalf("validateTeacher() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestTeacherResponseDoesNotExposePassword(t *testing.T) {
	info := teacherToInfo(&model.Teacher{Id: 1, TeacherNo: "T001", Name: "李老师", Password: "secret-hash"})
	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("marshal teacher response: %v", err)
	}
	response := string(data)
	if strings.Contains(response, "password") || strings.Contains(response, "secret-hash") {
		t.Fatalf("teacher response exposes password: %s", response)
	}
}
