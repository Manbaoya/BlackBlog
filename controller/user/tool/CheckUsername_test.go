package tool

import (
	"BLACKBLOG/controller"
	"reflect"
	"testing"
)

func TestCheckUsername(t *testing.T) {
	mp := []string{
		"涂山璟",
	}
	var got, want controller.Respond
	for _, v := range mp {
		got = CheckUsername(v)
		want = controller.OK
		if !reflect.DeepEqual(want, got) {
			t.Errorf(" username:%v excepted:%v result:%v", v, want, got)
		}
	}
}
