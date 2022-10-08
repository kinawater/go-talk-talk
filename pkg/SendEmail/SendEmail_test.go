package SendEmail

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	err := SendEmail([]string{"1774989203@qq.com"}, "今天天气真不赖，可以出来玩了", []byte("左三圈，右三圈，脖子扭扭，屁股扭扭，早睡早起我身体好"))
	if err != nil {
		t.Error("send email err,err info is ", err)
	}
}
