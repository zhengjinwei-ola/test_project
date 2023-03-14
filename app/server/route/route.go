package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.ForwardedByClientIP = true
	r.GET("/hello", func(ctx *gin.Context) {
		//kafka.Client.Send("test_log", "777888")
		ctx.JSON(http.StatusOK, "ok")
	})
}
