package tool

import (
	"BLACKBLOG/controller"
	"reflect"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	mp := make(map[string]string)

	//209 密码只能由数字、大小写字母、‘@’、‘#’、‘$’、‘&’、‘.’组成,6-12位
	mp = map[string]string{
		"123":           "123",
		"1234567890111": "1234567890111",
		" sdfsd":        " sdfsd",
		"^123qwe":       "^123qwwe",
	}
	var got, want controller.Respond
	for i, v := range mp {
		got = CheckPassword(i, v)
		want = controller.BadPassword
		if !reflect.DeepEqual(want, got) {
			t.Errorf(" password:%v repassword:%v excepted:%v result:%v", i, v, want, got)
		}
	}

	//210两次密码不一致
	got = CheckPassword("1vsavd", "1vsfbaavd")
	want = controller.DifferentPwd
	if !reflect.DeepEqual(want, got) {
		t.Errorf(" password:1vsavd repassword:1vsfbaavd excepted:%v result:%v", want, got)
	}

}
