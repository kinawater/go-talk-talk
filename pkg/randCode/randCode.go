package randCode

import (
	"math/rand"
	"strconv"
	"time"
)

// 去掉了0，O，o，i,I，1等容易混淆的元素
var AlphanumericSet = []rune{
	'2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

//NotRepeatingRandCode 生成一个uid一个固定的邀请码
//uid string 用户uid，可以是数字id也可以是用户名
//len int code的长度
func NotRepeatingRandCode(uid string, length int) []rune {
	var code []rune
	uidNum := getUidConvInt(uid)
	for i := 0; i < length; i++ {
		AlphanumericSetIndex := uidNum % int64(len(AlphanumericSet))
		code = append(code, AlphanumericSet[AlphanumericSetIndex])
		uidNum = uidNum / int64(len(AlphanumericSet))
	}
	return code
}

//GetRandCode 生成一个随机验证码，不保证不碰撞
//uid string 用户uid，可以是数字id也可以是用户名
//len int code的长度
func GetRandCode(uid string, length int) []rune {
	var code []rune
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		idx := r.Intn(len(AlphanumericSet))
		code = append(code, AlphanumericSet[idx])
	}
	return code
}

func getUidConvInt(uid string) int64 {
	var uidString string
	for _, r := range []rune(uid) {
		tempInt := int64(r)
		uidString = uidString + strconv.FormatInt(tempInt, 10)
	}
	num, _ := strconv.ParseInt(uidString, 10, 64)
	return num
}
