package middlewarelog
import(
"github.com/gin-gonic/gin"
"github.com/sirupsen/logrus"
"fmt"
)

func Logger(log *logrus.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        // use logrus
	   fmt.Println("<--------------------printing in log file---------------------->")
	   fmt.Println(log.WithError)

    }
}