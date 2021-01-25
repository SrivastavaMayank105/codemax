# codemax

# For deoplying project
#1 . Log-in the MailJet and get your API key and Secret key
#2. In windows add the evniornment variable by below commands
#: set MJ_APIKEY_PUBLIC ="your API key" hit enter
#: set MJ_APIKEY_PRIVATE="your secret key" hit enter
#3: run the main.go
#4: go to the Postman select Post request and enter the url:http://localhost:8080/mail
#5: click on the Body tab and enter the below:
{
    "to":"to_mailid ",
    "from":"from_mailid",
    "bcc":"bcc_mailid"

}
eg:{
    "to":"mayanksri105@gmail.com",
    "from":"mayanksri105@gmail.com",
    "bcc":"mayanksri105@gmail.com"

}

#6. if you want to use the AmazonSes and Sendgrind, go in the demo/sendingoption 
and change the below code in Mailoption func in  sending.go file
sm.AmazonSes()
or sm.SendGrid()

#6: I have use Gin framework for API

