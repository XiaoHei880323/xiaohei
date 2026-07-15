package course

import (
	"api/model"
	"encoding/json"
	"strings"
	"testing"
)

func TestValidateStudent(t *testing.T) {
	status := 1
	tests := []struct {
		name         string
		studentNo    string
		studentName  string
		gender       int
		status       *int
		password     string
		passwordNeed bool
		want         string
	}{
		{name: "valid add", studentNo: "S001", studentName: "张三", gender: 1, status: &status, password: "123456", passwordNeed: true},
		{name: "student number required", studentName: "张三", status: &status, password: "123456", passwordNeed: true, want: "请输入学号"},
		{name: "name required", studentNo: "S001", status: &status, password: "123456", passwordNeed: true, want: "请输入学生姓名"},
		{name: "gender invalid", studentNo: "S001", studentName: "张三", gender: 3, status: &status, password: "123456", passwordNeed: true, want: "性别参数错误"},
		{name: "password too short", studentNo: "S001", studentName: "张三", status: &status, password: "123", passwordNeed: true, want: "密码不能少于6位"},
		{name: "empty update password", studentNo: "S001", studentName: "张三", status: &status},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := validateStudent(test.studentNo, test.studentName, test.gender, test.status, test.password, test.passwordNeed)
			if got != test.want {
				t.Fatalf("validateStudent() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestStudentResponseDoesNotExposePassword(t *testing.T) {
	info := studentToInfo(&model.Student{Id: 1, StudentNo: "S001", Name: "张三", Password: "secret-hash"})
	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("marshal student response: %v", err)
	}
	response := string(data)
	if strings.Contains(response, "password") || strings.Contains(response, "secret-hash") {
		t.Fatalf("student response exposes password: %s", response)
	}
}
