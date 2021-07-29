package auth

import (
	"HUST-Matcher/database"
	"HUST-Matcher/model"
	"fmt"
	"math/rand"
	"net/smtp"
	"regexp"
	"time"
)

const (
	SMTP_MAIL_HOST     = "smtp.qq.com"
	SMTP_MAIL_PORT     = "587"
	SMTP_MAIL_USER     = model.QQ_EMAIL
	SMTP_MAIL_PWD      = model.QQ_PWD
	SMTP_MAIL_NICKNAME = model.QQ_NICKNAME
)

func EmailAuth(StudentID string) string {
	var err error
	var subject, body string              // 主题，邮件内容
	address := StudentID + "@hust.edu.cn" // 收件人
	auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	contentType := "Content-Type: text/html; charset=UTF-8"
	c := database.GetRClient()

	//	生成六位随机验证码，存入redis
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	err = c.Set(StudentID, vcode, time.Minute*5).Err()
	if err != nil {
		panic(err)
	}

	//	设置邮件主题和正文，写入验证码
	subject = "HUST-Matcher邮箱验证"
	body = vcode
	fmt.Println(vcode)

	//	要发送的消息，先用s格式化
	s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",
		address, SMTP_MAIL_NICKNAME, SMTP_MAIL_USER, subject, contentType, body)
	msg := []byte(s)

	//	邮件服务地址格式 host:port
	addrFormat := fmt.Sprintf("%s:%s", SMTP_MAIL_HOST, SMTP_MAIL_PORT)

	//	发邮件
	err = smtp.SendMail(addrFormat, auth, SMTP_MAIL_USER, []string{address}, msg)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("send email successfully")
	}

	return vcode

}

func CheckString(string string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", string); !ok {
		return false
	}
	return true
}
