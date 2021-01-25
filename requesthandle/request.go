package RequestHandle

import("github.com/gin-gonic/gin"
	dm "demo/sendingoption"
	"demo/middlewarelog"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
func HandleEmailRequest(c *gin.Context){
	var ma dm.MailFormat

	err := c.BindJSON(&ma)
	if err != nil {
		logrus.Error(err.Error())
		middlewarelog.Logger(log)
	}else{
		//code 
		//fmt.Println(ma)
		dm.Mailoption(&ma)
	}

}