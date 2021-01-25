package sendingoption

import(
	 "demo/servicemail"
)

type MailFormat struct{
	To string `json:"to" binding:"required"`
	From string `json:"from" binding:"required"`
	Bcc string `json:"bcc" binding:"required"`
}

type Mailsend interface{
	MailInject()
	AmazonSes()
	SendGrid()
}

func (m *MailFormat)MailInject(){
	//code
	servicemail.MailinjectMailservice(m.To,m.From,m.Bcc)

}

func (m *MailFormat)AmazonSes(){
	//code
	servicemail.AmazonSesservice(m.To,m.From,m.Bcc)
}

func (m *MailFormat)SendGrid(){
	//code
	servicemail.Sendgrindservice(m.To,m.From,m.Bcc)
}

func Mailoption(sm Mailsend){
	sm.MailInject()
}