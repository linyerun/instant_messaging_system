package utils

import (
	"context"
	"github.com/jordan-wright/email"
	"instant_messaging_system/db"
	"instant_messaging_system/define"
	"math/rand"
	"net/smtp"
	"regexp"
	"strconv"
	"time"
)

// VerifyEmail 验证邮箱
func VerifyEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

// GetCode 生成验证码
func GetCode() string {
	rand.Seed(time.Now().UnixNano()) // 种随机种子
	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

// SendCode 发送验证码
func SendCode(toUserEmail, code string) error {
	// 避免验证码1分钟发送多次
	ctx := context.Background()
	const prefix = "code_to_email:"
	if err := db.RedisClient.Get(ctx, prefix+toUserEmail).Err(); err.Error() != db.RedisNil {
		return err
	}
	defer db.RedisClient.Set(ctx, prefix+toUserEmail, "", time.Duration(define.EachCodeForEmailWaitTime)*time.Second)
	// 发送邮件
	e := email.NewEmail()
	e.From = "ImSystem <2338244917@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2338244917@qq.com", define.MailPassword, "smtp.qq.com"))
}
