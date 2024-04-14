package dao

import (
	"context"
	"fmt"
	"time"
)

func SetHotComments() { //放置最新20条评论
	ctx := context.Background()
	RDB.Set(ctx, "key", "123", 500*time.Millisecond)
	result := RDB.Get(ctx, "key")

	fmt.Println(result.Val())
}
func GetHotComments() { //获取最新20条评论

}
