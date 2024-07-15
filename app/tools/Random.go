package tools

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

// 生成指定长度的随机字符串 length代表长度
func Get_Random_String(length int) string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charSetLength := len(charSet)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(charSetLength)
		result[i] = charSet[randomIndex]
	}

	return string(result)
}

// 生成指定长度的随机数字 length代表长度
func Get_Random_Int(length int) string {
	charSet := "0123456789"
	charSetLength := len(charSet)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(charSetLength)
		result[i] = charSet[randomIndex]
	}

	return string(result)
}

// 生成年月日 示例：060102
func Get_ymd() string {
	return time.Now().Format("060102")
}

// 生成年月日 示例：20060102
func Get_Ymd() string {
	return time.Now().Format("20060102")
}

// 生成年月日时分秒毫秒 示例：202407131112311000
func Get_Ymdhism() string {
	currentTime := time.Now()
	millisecond := currentTime.Nanosecond() / 1000000 //毫秒
	return Get_Ymdhis() + fmt.Sprintf("%d", millisecond)
}

// 生成年月日时分秒 示例：20240713111231
func Get_Ymdhis() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102150405") //年月日时分秒
	return formattedTime
}

// 获取本机IP地址
func MyIP() string {
	var result = ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if strings.TrimSpace(ipnet.IP.String()) != "" {
					result = result + ipnet.IP.String() + "\n"
				}
			}
		}
	}
	return strings.TrimSpace(result)
}
