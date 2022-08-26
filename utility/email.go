package utility

import (
	"fmt"
	"log"

	. "zjutjh/Join-Us/config"
	"zjutjh/Join-Us/db/model"

	"gopkg.in/gomail.v2"
)

// MailboxConf 邮箱配置
type MailboxConf struct {
	Title         string
	Body          string
	RecipientList []string
	Sender        string
	SPassword     string
	SMTPAddr      string
	SMTPPort      int
}

func SendEmail(form model.NormalForm) error {
	var mailConf MailboxConf
	mailConf.Title = "【精弘网络】收到一条新的招新表。"
	mailConf.RecipientList = []string{Config.Email.Receiver}
	mailConf.Sender = Config.Email.Sender

	mailConf.SPassword = Config.Email.Pwd

	mailConf.SMTPAddr = Config.Email.SmtpAddr
	mailConf.SMTPPort = Config.Email.SmtpPort

	html := fmt.Sprintf(
		`<p收到一条新的招新表单</p>
		<p>姓名:%v</p>
		<p>性别:%v</p>
		<p>学号:%v</p>
		<p>电话号:%v</p>
		<p>QQ:%v</p>
		<p>校区:%v</p>
		<p>专业:%v</p>>
		<p>学院:%v</p>
		<p>第一志愿:%v</p>
		<p>第二志愿:%v</p>
		<p>简介:%v</p>
		<p>反馈:%v</p>`,
		form.Name,
		Gender[form.Gender],
		form.StuID,
		form.Phone,
		form.QQ,
		Regions[form.Region],
		form.College,
		form.Campus,
		Departments[form.Want1],
		Departments[form.Want2],
		form.Profile,
		form.Feedback,
	)

	m := gomail.NewMessage()

	m.SetHeader(`From`, mailConf.Sender, "技术部-精弘首页")
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	if !Config.Dev {
		err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
		if err != nil {
			log.Printf("Send Email Fail, %s", err.Error())
			return err
		}
	}
	return nil
}
