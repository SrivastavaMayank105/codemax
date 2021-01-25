package servicemail

import (
mailjet "github.com/mailjet/mailjet-apiv3-go"
"os"
"github.com/sirupsen/logrus"
"demo/middlewarelog"

)

var log = logrus.New()
func MailinjectMailservice(to,from,bcc string){
	//fmt.Println("I am sending through Mail Inject Services",to,from,bcc)
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")
	mj := mailjet.NewMailjetClient(publicKey, secretKey)
	email := &mailjet.InfoSendMail {
		FromEmail: from,
		FromName: "Mayank Srivastava",
		Subject: "First Email",
		TextPart: "This is my first experience with email,hope you will enjoy",
		HTMLPart: "<h3>Hey!! How you are doing</h3>",
		To: to,
		Bcc: bcc,
	  }
	  _, err := mj.SendMail(email)
	  if err != nil {
			  //log.Println(err)
			logrus.Error(err.Error())
			  middlewarelog.Logger(log)
	  } else {
			//  log.Println("Success")
			//lg.Println(err)
			 //middlewarelog.Logger(res,"")
	  }
}