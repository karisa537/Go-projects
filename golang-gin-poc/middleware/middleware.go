package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s %s %s %d [%s] %s \n",
				param.ClientIP,
				param.Latency,
				param.Method,
				param.Path,
				param.StatusCode,
				param.TimeStamp.Format(time.RFC822),
				param.Keys,
	)
	})

}