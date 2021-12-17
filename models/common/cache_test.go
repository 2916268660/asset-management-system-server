package common

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	err := SetKey("ws", "123", time.Second*60)
	if err != nil {
		t.Error("设置key成功", err)
	}
	ok := IsExistKey("ws")
	t.Log(ok)
	key, err := GetKey("ws")
	if err != nil {
		t.Error("获取key失败", err)
	}
	t.Log(key)
}
