package redis

import (
	"strconv"

	"fmt"

	. "github.com/open-fightcoder/oj-web/common/store"
)

func SubmitCountGet(userId int64) (string, error) {
	res := RedisClient.HMGet("submit_count", strconv.FormatInt(userId, 10))
	if res.Err() != nil {
		return "", res.Err()
	}
	return fmt.Sprint(res.Val()[0]), nil
}