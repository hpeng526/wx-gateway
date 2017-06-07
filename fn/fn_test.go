package fn

import (
	"fmt"
	"github.com/hpeng526/wx-gateway/po"
	"strconv"
	"strings"
	"testing"
)

func TestTmpData(t *testing.T) {
	var nums = []int{1, 2, 3}
	strs := Map(nums, strconv.Itoa).([]string)
	fmt.Println(strs[2])
	users := []po.User{{UserWXId: "xxxx", UserId: 123}, {UserWXId: "ssss", UserId: 222}}
	ustr := Map(users, func(u po.User) string {
		return u.UserWXId
	}).([]string)
	fmt.Println(strings.Join(ustr, ","))
}
