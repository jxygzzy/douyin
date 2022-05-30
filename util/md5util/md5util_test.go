package md5util_test

import (
	"douyin/util/md5util"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	md5 := md5util.MD5WithSalt("jxygzzy", "123456")
	fmt.Printf("md5:%v", md5)
}
