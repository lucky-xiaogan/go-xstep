package xredis

import (
	"context"
	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
	ctx := context.Background()
	rdb := NewRedis("127.0.0.1:6379", "", DialDatabase(0))
	r, err := rdb.Set(ctx, "test", 1, 10*time.Second).Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}
