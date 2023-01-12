package helper

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

const (
	QQMailSMTPCode = "knntcfcyxzmmbafj"
	QQMailSender   = "1598273095@qq.com"
	QQMailTitle    = "验证"
	SMTPAdr        = "smtp.qq.com"
	SMTPPort       = 587
	MailListSize   = 2048

	SecretId    = "AKIDSjlgKhpgFXYIkvpRmqc8MK1ScU5PSGo4"
	SecreKey    = "1lDOJxKQFn44nFPItDH4ZnkGWUCdmeFZ"
	SmsSdkAppId = "1400696970"
	TemplateId  = "1452825"
)

type MailboxConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 收件人列表
	RecipientList []string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	SPassword string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
}

// SendMail QQ邮箱验证码
func SendMail(RedisDB *redis.Client, mail string) (error, bool) {
	var mailConf MailboxConf
	var mails []string
	mails = append(mails, mail)
	mailConf.Title = QQMailTitle
	//这里就是我们发送的邮箱内容，但是也可以通过下面的html代码作为邮件内容
	// mailConf.Body = "坚持才是胜利，奥里给"

	//这里支持群发，只需填写多个人的邮箱即可，我这里发送人使用的是QQ邮箱，所以接收人也必须都要是
	//QQ邮箱
	mailConf.RecipientList = mails
	mailConf.Sender = QQMailSender

	//这里QQ邮箱要填写授权码，网易邮箱则直接填写自己的邮箱密码，授权码获得方法在下面
	mailConf.SPassword = QQMailSMTPCode

	//下面是官方邮箱提供的SMTP服务地址和端口
	// QQ邮箱：SMTP服务器地址：smtp.qq.com（端口：587）
	// 雅虎邮箱: SMTP服务器地址：smtp.yahoo.com（端口：587）
	// 163邮箱：SMTP服务器地址：smtp.163.com（端口：25）
	// 126邮箱: SMTP服务器地址：smtp.126.com（端口：25）
	// 新浪邮箱: SMTP服务器地址：smtp.sina.com（端口：25）

	mailConf.SMTPAddr = SMTPAdr
	mailConf.SMTPPort = SMTPPort

	//产生六位数验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//发送的内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vcode)

	m := gomail.NewMessage()

	// 第三个参数是我们发送者的名称，但是如果对方有发送者的好友，优先显示对方好友备注名
	m.SetHeader(`From`, mailConf.Sender, "NewGym")
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	// m.Attach("./Dockerfile") //添加附件
	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	if err != nil {
		logx.Error("Send Email Fail, %s", err.Error())
		return err, false
	}
	for _, j := range mails {
		//	1.定义验证码和手机号的key的形式
		codeKey := j + vcode
		mailKey := "pkey" + j
		//	2.查看用户发送验证码的数量
		ctn, _ := RedisDB.Get(mailKey).Int()
		//  ctn为零值的时候，说明该用户还没有发送过验证码
		if ctn == 0 {
			RedisDB.Set(mailKey, 1, time.Minute*60*24) //设置手请求验证码的时间为一天。
		}
		if ctn <= 99 {
			RedisDB.Incr(mailKey)
		} else {
			return nil, false
		}
		//	将验证码存入redis,设置过期时间为2分钟
		RedisDB.Set(codeKey, vcode, time.Minute*5)
	}
	logx.Info("Send Email Success")
	return nil, true
}
