package md5util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func MD5WithSalt(s string, salt string) string {
	sum := md5.Sum([]byte(s + salt))
	return hex.EncodeToString(sum[:])
}

func CalcFileMD5(filename string) (string, error) {
    f, err := os.Open(filename) //打开文件
    if nil != err {
        fmt.Println(err)
        return "", err
    }
    defer f.Close()
    
    md5Handle := md5.New()      //创建 md5 句柄
    _, err = io.Copy(md5Handle, f)  //将文件内容拷贝到 md5 句柄中
    if nil != err {
        fmt.Println(err)
        return "", err
    }
    md := md5Handle.Sum(nil)    //计算 MD5 值，返回 []byte    
    md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
    return md5str, nil
}