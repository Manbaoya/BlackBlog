package tool

import (
	"BLACKBLOG/controller"
	"reflect"
	"testing"
)

func TestCheckPhoneCheckPhone(t *testing.T) {

	mp := []string{
		"123",
		"1234567890111",
		"1334g",
	}
	var got, want controller.Respond
	for _, v := range mp {
		got = CheckPhone(v)
		want = controller.BadPhone
		if !reflect.DeepEqual(want, got) {
			t.Errorf(" phone:%v excepted:%v result:%v", v, want, got)
		}
	}

}
